package old

import (
	"sync"
)

type Semaphore interface {
	Wait()
	Signal()
}

type BinarySemaphore struct{
	key sync.Mutex
}

func NewBinarySemaphore() *BinarySemaphore {
	return &BinarySemaphore{ key: sync.Mutex{} }
}

func (bs *BinarySemaphore) Wait() {
	bs.key.Lock()
}

func (bs *BinarySemaphore) Signal() {
	bs.key.Unlock();
}

type empty struct {}

type GeneralSemaphore struct{
	resources chan empty
}

func NewGeneralSemaphore(resources int) *GeneralSemaphore {
	gSem := GeneralSemaphore{ resources: make(chan empty, resources) }
	for i:= 0; i < resources; i++ {
		gSem.resources <- empty{}
	}
	
	return &gSem
}

func (sem *GeneralSemaphore) Wait() {
	<- sem.resources
}

func (sem *GeneralSemaphore) Signal() {
	sem.resources <- empty{}
}




