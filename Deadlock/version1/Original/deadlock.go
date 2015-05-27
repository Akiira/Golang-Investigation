//********************
//** HW#4 Problem 3 **
//**  deadlock.cm   **
//********************

package main

import (
	"fmt"
	"runtime"
	"sem"
	"sync"
)

/* Global Variables */
var (
	resource1 = sem.NewGeneralSemaphore(1)
	resource2 = sem.NewGeneralSemaphore(1)
)

//////////////////////////////////////////////////////////////
func printmsg(x int) {
	fmt.Println("Process ", x, " now has both resources.")
}

//////////////////////////////////////////////////////////////
func Process1() {
	resource1.Wait()
	resource2.Wait()
	printmsg(1)
	resource2.Signal()
	resource1.Signal()
}

//////////////////////////////////////////////////////////////
func Process2() {
	resource2.Wait()
	resource1.Wait()
	printmsg(2)
	resource1.Signal()
	resource2.Signal()
}

//////////////////////////////////////////////////////////////
func main() {
	fmt.Println("Starting...\n")

	CoBegin(
		Process1,
		Process2,
	)
}

//This can be moved to a package so the students don't see it.
func CoBegin(procs ...func()) {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(len(procs) + 2)

	for _, proc := range procs {
		go func(varF func()) { varF(); wg.Done() }(proc)
		wg.Add(1)
	}

	wg.Wait()
}
