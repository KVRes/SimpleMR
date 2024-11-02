package worker

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
