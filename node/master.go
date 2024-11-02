package node

import (
	"github.com/KVRes/SimpleMR/types"
	"github.com/KVRes/SimpleMR/worker"
)

type Master struct {
	splitter types.Splitter
	mWorker  *worker.Pool
	rWorker  *worker.Pool
}

func (master *Master) Start(
	data any, nMap, nReduce int,
	mapFx func(*MapContext, any),
	reduceFx func(types.ReduceTask) any,
) []any {
	master.initState()

	// Map
	mTasks := master.splitter.SplitDataIntoMPieces(data, nMap)
	applyAll(mTasks, func(mTask any) {
		master.mWorker.AssignWork(func() any {
			ctx := NewMapContext()
			mapFx(ctx, mTask)
			return ctx
		})
	})

	master.mWorker.WaitAll()
	rstUntyped := master.mWorker.Results()
	rst := mapAll(rstUntyped, func(v any) MapResult {
		return v.(*MapContext).m
	})

	// TODO: Combinator

	// Shuffle
	rTasks := make([]types.ReduceTask, nReduce)
	for i := 0; i < nReduce; i++ {
		rTasks[i] = make(map[string][]any)
	}
	for _, r := range rst {
		for k, v := range r {
			bucketId := HashBucket(k, nReduce)
			rTasks[bucketId][k] = v
		}
	}

	// Reduce
	applyAll(rTasks, func(rTask types.ReduceTask) {
		master.rWorker.AssignWork(func() any {
			return reduceFx(rTask)
		})
	})

	master.rWorker.WaitAll()

	return master.rWorker.Results()

}
