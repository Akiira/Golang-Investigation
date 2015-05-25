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
var wg sync.WaitGroup
var Num int
var mutex = &sync.Mutex{}

func IncrementNum() {

	for i := 1; i <= 5; i++ {
		mutex.Lock()
		Num = Num + i
		mutex.Unlock()
	}

	wg.Done()
}

func DecrementNum() {

	for i := 1; i <= 5; i++ {
		mutex.Lock()
		Num = Num - i
		mutex.Unlock()
	}

	wg.Done()
}

func main() {
	Init()

	Num = 10

	go IncrementNum()
	go DecrementNum()

	wg.Wait()

	fmt.Println("The final value of Num is: ", Num)
}

//Do not edit
func Init() {
	runtime.GOMAXPROCS(4)
	wg.Add(2)
}
