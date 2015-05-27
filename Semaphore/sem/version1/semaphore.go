package version1

import (
	"sync"
)

type Semaphore struct{
	key sync.Mutex
}

func New() *Semaphore {
	return &Semaphore{ key: sync.Mutex{} }
}

func (sem *Semaphore) Wait() {
	sem.key.Lock()
}

func (sem *Semaphore) Signal() {
	sem.key.Unlock();
}
