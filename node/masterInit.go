package node

import (
	"github.com/KVRes/SimpleMR/types"
	"github.com/KVRes/SimpleMR/worker"
)

func (master *Master) WithSplitter(s types.Splitter) *Master {
	master.splitter = s
	return master
}

func (master *Master) WithMapWorker(m int) *Master {
	master.mWorker = worker.NewWorkerPool(m)
	return master
}

func (master *Master) WithReduceWorker(r int) *Master {
	master.rWorker = worker.NewWorkerPool(r)

	return master
}

func (master *Master) validate() bool {
	return master.splitter != nil &&
		master.mWorker != nil &&
		master.rWorker != nil
}

func (master *Master) initState() {
	//m := master.mWorker.MaxThreads
	//r := master.rWorker.MaxThreads
	//master.state = make([][]MasterCell, m)
	//for i := 0; i < len(master.state); i++ {
	//	master.state[i] = make([]MasterCell, r)
	//}
	if master.partitioner == nil {
		master.partitioner = HashPartitioner
	}

	if !master.validate() {
		panic("Master not initialised")
	}
}

func NewMaster() *Master {
	return &Master{}
}
