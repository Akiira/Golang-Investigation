//***************************************************
//** HW#4 Problem 2: use for parts a,b, c, and d   **
//**               order.cm                        **
//***************************************************

package main

import (
	"fmt"
	"runtime"
	"sync"
)

///////////////////////////////////
func PrintOut(pid int) {
	fmt.Println("Process ", pid)
}

///////////////////////////////////
func Process1() {
	PrintOut(1)
}

///////////////////////////////////
func Process2() {
	PrintOut(2)
}

///////////////////////////////////
func Process3() {
	PrintOut(3)
}

///////////////////////////////////
func main() {

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
