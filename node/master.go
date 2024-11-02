package node

import (
	"github.com/KVRes/SimpleMR/types"
	"github.com/KVRes/SimpleMR/worker"
)

type Master struct {
	splitter    types.Splitter
	mWorker     *worker.Pool
	rWorker     *worker.Pool
	partitioner types.Partitioner
}

func (master *Master) Start(
	data any, nMap, nReduce int,
	mapFx func(*MapContext, any),
	reduceFx func(types.Intermediate) any,
	combineFx func(*MapContext, types.Intermediate),
) []any {
	master.initState()

	// Map
	mTasks := master.splitter(data, nMap)
	applyAll(mTasks, func(mTask any) {
		master.mWorker.AssignWork(func() any {
			ctx := NewMapContext()
			mapFx(ctx, mTask)
			if combineFx != nil {
				_rst := ctx.m
				ctx = NewMapContext()
				combineFx(ctx, _rst)
			}
			return ctx
		})
	})

	master.mWorker.WaitAll()
	rstUntyped := master.mWorker.Results()
	rst := mapAll(rstUntyped, func(v any) types.Intermediate {
		return v.(*MapContext).m
	})

	// Shuffle
	rTasks := make([]types.Intermediate, nReduce)
	for i := 0; i < nReduce; i++ {
		rTasks[i] = make(map[string][]any)
	}
	for _, r := range rst {
		for k, v := range r {
			bucketId := master.partitioner(k) % nReduce
			src := rTasks[bucketId]
			rTasks[bucketId][k] = append(src[k], v...)
		}
	}

	// Reduce
	applyAll(rTasks, func(rTask types.Intermediate) {
		master.rWorker.AssignWork(func() any {
			return reduceFx(rTask)
		})
	})

	master.rWorker.WaitAll()

	return master.rWorker.Results()
}
