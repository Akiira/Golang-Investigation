//***************************************************
//** HW#4 Problem 2: use for parts a,b, c, and d   **
//**               order.cm                        **
//***************************************************

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sem"
)

var (
	semB = sem.NewGeneralSemaphore(0)
	semC = sem.NewGeneralSemaphore(0)
)

///////////////////////////////////
func PrintOut(pid int) {
	fmt.Println("Process ", pid)
}

///////////////////////////////////
func Process1() {
	for i := 0; i < 3; i++ {	
		PrintOut(1)
	}
	
	semB.Signal();
}

///////////////////////////////////
func Process2() {
	semB.Wait()
	
	for i := 0; i < 3; i++ {		
		PrintOut(2)
	}
	
	semC.Signal();
}

///////////////////////////////////
func Process3() {
	semC.Wait()
	
	for i := 0; i < 3; i++ {	
		PrintOut(3)	
	}
}

///////////////////////////////////
func main() {
	fmt.Println("Starting...\n")
	// This will start all 3 processes concurrently
	CoBegin(
		Process1,
		Process2,
		Process3,
	)
}

//This can be moved to a package so the students don't see it.
func CoBegin(procs ...func()) {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(len(procs) + 2)
	
	for _, proc := range procs {
		go func(varF func()) { varF(); wg.Done()}(proc)
		wg.Add(1)
	}

	wg.Wait()
}
