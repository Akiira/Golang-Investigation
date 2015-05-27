package Normal_Version

import (
	"sync"
	"sem"
)

type Buffer struct {
	buffer []product
	resource_smphr sem.GeneralSemaphore
	buff_mutex sync.Mutex
}

func (b *Buffer) Add(prdct product) {
	b.buff_mutex.Lock()
	b.buffer = append(b.buffer, prdct)
	b.buff_mutex.Unlock()
	
	b.resource_smphr.Signal()
}

func (b *Buffer) Get() (p product) {
	b.resource_smphr.Wait()
	
	b.buff_mutex.Lock()
	p = b.buffer[0]
	b.buffer = b.buffer[1:]
	b.buff_mutex.Unlock()
	
	return p
}