//*********************
//** HW#4 Problem 1  **
//**   ShareNum.go   **
//*********************

package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* Global Variables  */
var Num int

func IncrementNum() {
	for i := 1; i <= 3; i++ {
		Num = Num + i
	}
}

func DecrementNum() {
	for i := 1; i <= 3; i++ {
		Num = Num - i
	}
}

func main() {
	Num = 10
	
	CoBegin(
		IncrementNum,
		DecrementNum,
	)

	fmt.Println("The final value of Num is: ", Num)
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

