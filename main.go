package main

import (
	"sync"
	"taskmanager/manager"
	"taskmanager/task1"
	"taskmanager/task2"
	"time"
)

func main() {
	tm := manager.GetTaskManager()
	tm.ExecuteTasks()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// mimicing multiple clients trying to add task 1
		for i := 0; i < 10; i++ {
			t1 := task1.Task1{}
			tm.AddTask(&t1)
			time.Sleep(1000)
		}
		wg.Done()
	}()

	go func() {
		//mimicing multiple clients trying to add task 2
		for i := 0; i < 10; i++ {
			t2 := task2.Task2{}
			tm.AddTask(&t2)
			time.Sleep(1000)
		}
		wg.Done()
	}()
	wg.Wait()
}
