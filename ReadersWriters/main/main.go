// Based off of:
// http://greenteapress.com/semaphores/


package main

import (
	"sync"
	"cmpsc"
)

var (
	readers  int        = 0
	rdrs_mtx sync.Mutex = sync.Mutex{}
	empty    sync.Mutex = sync.Mutex{}
)

func Reader() {
	for {
		rdrs_mtx.Lock()
		if readers == 0 {
			empty.Lock()
		} else {
			readers++
		}
		rdrs_mtx.Unlock()
		
		//Now in critical section
		
		
		//Exit critical section
		rdrs_mtx.Lock()
		
		readers--
		
		if readers == 0 {
			empty.Unlock()
		}
		
		rdrs_mtx.Unlock()
	}
}

func Writer() {
	empty.Lock()
	
	//Critical section
	
	empty.Unlock()
}

func main() {
	cmpsc.CoBegin(
		Writer,
		Writer,
		Reader,
		Reader,
		Reader,
		Reader,
		Reader,
		Reader,
	)
}
