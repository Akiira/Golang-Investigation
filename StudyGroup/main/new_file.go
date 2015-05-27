package main

import (
	"cmpsc"
	"fmt"
	"sync"
)

var (
	student_occupiers  int        = 0
	room_enter_or_exit sync.Mutex = sync.Mutex{}

	dangle_free       bool       = true
	dangle_free_mutex sync.Mutex = sync.Mutex{}
)

func main() {
	cmpsc.CoBegin(
		Student,
		Dangle,
	)
}

func Student() {
	cmpsc.DelayRandom()
	
	room_enter_or_exit.Lock()
	dangle_free_mutex.Lock()
	student_occupiers++
}

func Dangle() {
	room_enter_or_exit.Lock()

	if student_occupiers == 0 || student_occupiers > 5 {

		dangle_free_mutex.Lock()

		fmt.Println("Officer Dangle entered the room.")
		room_enter_or_exit.Unlock()

		//Waiting for students to leave
		for student_occupiers != 0 {}

		cmpsc.DelayRandom()

		room_enter_or_exit.Lock() //get enter/exit cool so dangle can leave
		dangle_free_mutex.Unlock()
		fmt.Println("Officer Dangle left the room.")		
	}
	
	room_enter_or_exit.Unlock()
}

//func Dangle() {
//	room_enter_or_exit.Lock()
//
//	if student_occupiers == 0 || student_occupiers > 5 {
//
//		dangle_free_mutex.Lock()
//		dangle_free = false
//		dangle_free_mutex.Unlock()
//
//		fmt.Println("Officer Dangle entered the room.")
//		room_enter_or_exit.Unlock()
//
//		//Waiting for students to leave
//		for student_occupiers != 0 {}
//
//		cmpsc.DelayRandom()
//
//		room_enter_or_exit.Lock() //get enter/exit cool so dangle can leave
//
//		dangle_free_mutex.Lock()
//		dangle_free = true
//		dangle_free_mutex.Unlock()
//
//		fmt.Println("Officer Dangle left the room.")		
//	}
//	
//	room_enter_or_exit.Unlock()
//}
