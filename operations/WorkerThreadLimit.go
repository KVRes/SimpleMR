package operations

import "sync"

type SyncArr struct {
	l   sync.Mutex
	Raw []any
}

func (a *SyncArr) Push(data any) {
	a.l.Lock()
	defer a.l.Unlock()
	a.Raw = append(a.Raw, data)
}

func (a *SyncArr) Get(idx int) any {
	a.l.Lock()
	defer a.l.Unlock()
	return a.Raw[idx]
}

type WorkerPool struct {
	WorkerThreads []*WorkerThread
	MaxThreads    int
	wg            sync.WaitGroup
	rst           *SyncArr
}

type WorkerThread struct {
	Idle bool
	Lck  *sync.RWMutex
}

func (wp *WorkerPool) WaitAll() {
	wp.wg.Wait()
}

func (wp *WorkerPool) Results() []any {
	return wp.rst.Raw
}

func (wp *WorkerPool) ClearResults() {
	wp.rst = &SyncArr{}
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

	pl.ClearResults()

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
			w.work(fx, wp)
			return
		}
	}
}

type WorkFx func() any

func (w *WorkerThread) work(fx WorkFx, pool *WorkerPool) {
	defer pool.wg.Done()

	rst := fx()
	pool.rst.Push(rst)
	w.Lck.Lock()
	defer w.Lck.Unlock()
	w.Idle = true
}

func (wp *WorkerPool) AssignWorkQueue(fx []WorkFx) {
	for _, f := range fx {
		wp.AssignWork(f)
	}
}
