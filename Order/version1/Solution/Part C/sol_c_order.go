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
	semA = sem.NewGeneralSemaphore(1)
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
		semA.Wait()
		PrintOut(1)
		semB.Signal();
	}
}

///////////////////////////////////
func Process2() {
	for i := 0; i < 3; i++ {
		semB.Wait()
		PrintOut(2)
		semC.Signal();
	}
}

///////////////////////////////////
func Process3() {
	for i := 0; i < 3; i++ {
		semC.Wait()
		PrintOut(3)
		semA.Signal();
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
