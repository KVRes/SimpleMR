package operations

import "sync"

type WorkerPool struct {
	WorkerThreads []*WorkerThread
	MaxThreads    int
	wg            sync.WaitGroup
}

type WorkerThread struct {
	Idle bool
	Lck  *sync.RWMutex
}

func (wp *WorkerPool) AddWorkerThread() {
	wp.WorkerThreads = append(wp.WorkerThreads, &WorkerThread{
		Idle: true,
		Lck:  &sync.RWMutex{},
	})
}

func NewWorkerPool(maxThreads int) *WorkerPool {
	pl := &WorkerPool{
		MaxThreads: maxThreads,
	}

	for i := 0; i < maxThreads; i++ {
		pl.AddWorkerThread()
	}
	return pl
}

func (w *WorkerThread) IsIdle() bool {
	w.Lck.RLock()
	defer w.Lck.RUnlock()
	return w.Idle
}

func (w *WorkerThread) SetIdle(idle bool) {
	w.Lck.Lock()
	defer w.Lck.Unlock()
	w.Idle = idle
}

func (w *WorkerThread) setWorkState() bool {
	w.Lck.Lock()
	defer w.Lck.Unlock()
	if !w.Idle {
		return false
	}
	w.Idle = false
	return true
}

func (wp *WorkerPool) AssignWork(fx WorkFx) {
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
			w.work(fx, wp.wg.Done)
			return
		}
	}
}

type WorkFx func()

func (w *WorkerThread) work(fx WorkFx, doneFx WorkFx) {
	fx()
	w.Lck.Lock()
	defer w.Lck.Unlock()
	w.Idle = true
	doneFx()
}

func (wp *WorkerPool) AssignWorkQueue(fx []WorkFx) {
	for _, f := range fx {
		wp.AssignWork(f)
	}
}
