package version3

import (
	"sync"
)

type Semaphore struct {
	key *sync.Cond
}

func New() *Semaphore {
	return &Semaphore{key: sync.NewCond(&sync.Mutex{})}
}

func (sem *Semaphore) Wait() {
	sem.key.Wait()
}

func (sem *Semaphore) Signal() {
	sem.key.Signal()
}

func (sem *Semaphore) Broadcast() {
	sem.key.Broadcast()
}
