package main

import (
	"sync"
	"taskmanager/manager"
	"taskmanager/task1"
	"taskmanager/task2"
)

func main() {
	tm := manager.GetTaskManager()
	tm.ExecuteTasks()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		// mimicing multiple clients trying to add task 1
		for i := 0; i < 10; i++ {
			t1 := task1.Task1{}
			tm.Push(&t1)
		}
		wg.Done()
	}()

	go func() {
		//mimicing multiple clients trying to add task 2
		for i := 0; i < 10; i++ {
			t2 := task2.Task2{}
			tm.Push(&t2)
		}
		wg.Done()
	}()

	wg.Wait()
}
