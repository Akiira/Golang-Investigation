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

func IncrementNum() {

	for i := 1; i <= 5; i++ {
		Num = Num + i
	}

	wg.Done()
}

func DecrementNum() {

	for i := 1; i <= 5; i++ {
		Num = Num - i
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
