package node

import (
	"github.com/KVRes/SimpleMR/operations"
	"github.com/KVRes/SimpleMR/types"
)

type MapResult map[string]any

type Master struct {
	state      [][]MasterCell // M * R
	splitter   types.Splitter
	mapTracker []int
	mWorker    *operations.WorkerPool
	rWorker    *operations.WorkerPool
}

func (master *Master) Start(data any, n int, m, r int, mapFx func(any) MapResult) {
	master.state = make([][]MasterCell, m)
	master.mWorker = operations.NewWorkerPool(m)
	master.rWorker = operations.NewWorkerPool(r)

	for i := 0; i < m; i++ {
		master.state[i] = make([]MasterCell, r)
	}
	splitted := master.splitter.SplitDataIntoMPieces(data, n)
	for _, s := range splitted {
		master.mWorker.AssignWork(func() any {
			return mapFx(s)
		})
	}

	master.mWorker.WaitAll()

	rst := master.mWorker.Results()



}

func (master *Master) IsMapFinished(mapId int) bool {
	if len(master.state) == 0 {
		return true
	}
	col := len(master.state[mapId])

	for i := 0; i < col; i++ {
		if !master.state[mapId][i].MapOk {
			return false
		}
	}
	return true
}

type MasterCell struct {
	MapOk   bool
	MapList []any
}
