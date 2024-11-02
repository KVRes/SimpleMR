package worker

import "sync"

type Pool struct {
	WorkerThreads []*Thread
	MaxThreads    int
	wg            sync.WaitGroup
	rst           *SyncArr
}

func NewWorkerPool(maxThreads int) *Pool {
	pl := &Pool{MaxThreads: maxThreads}
	pl.ClearResults()

	for i := 0; i < maxThreads; i++ {
		pl.AddWorkerThread()
	}
	return pl
}

func (wp *Pool) WaitAll() {
	wp.wg.Wait()
}

func (wp *Pool) Results() []any {
	return wp.rst.Raw
}

func (wp *Pool) ClearResults() {
	wp.rst = &SyncArr{}
}

func (wp *Pool) AddWorkerThread() {
	wp.WorkerThreads = append(wp.WorkerThreads, &Thread{
		Idle: true,
		lck:  &sync.RWMutex{},
	})
}

func (wp *Pool) AssignWork(fx WorkFx) {
	wp.wg.Add(1)
	for {
		for i := 0; i < len(wp.WorkerThreads); i++ {
			w := wp.WorkerThreads[i]
			if !w.IsIdle() {
				continue
			}
			if !w.setWorkState() {
				continue
			}
			go w.work(fx, wp)
			return
		}
	}
}

func (wp *Pool) AssignAllWorks(fx []WorkFx) {
	for _, f := range fx {
		wp.AssignWork(f)
	}
}
