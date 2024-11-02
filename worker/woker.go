package worker

import "sync"

type WorkFx func() any

type Thread struct {
	Idle bool
	lck  *sync.RWMutex
}

func (w *Thread) IsIdle() bool {
	w.lck.RLock()
	defer w.lck.RUnlock()
	return w.Idle
}

func (w *Thread) SetIdle(idle bool) {
	w.lck.Lock()
	defer w.lck.Unlock()
	w.Idle = idle
}

func (w *Thread) setWorkState() bool {
	w.lck.Lock()
	defer w.lck.Unlock()
	if !w.Idle {
		return false
	}
	w.Idle = false
	return true
}

func (w *Thread) work(fx WorkFx, pool *Pool) {
	defer pool.wg.Done()

	rst := fx()
	pool.rst.Push(rst)
	w.lck.Lock()
	defer w.lck.Unlock()
	w.Idle = true
}
