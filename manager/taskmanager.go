package manager

import (
	"sync"
)

var taskManagerInstance *TaskManager
var taskManagerOnce sync.Once

type Task interface {
	Execute() error
}

type TaskManager struct {
	taskList chan Task
}

func GetTaskManager() *TaskManager {
	taskManagerOnce.Do(func() {
		if taskManagerInstance == nil {
			taskManagerInstance = &TaskManager{
				taskList: make(chan Task),
			}
		}
	})
	return taskManagerInstance
}

func (t *TaskManager) AddTask(task Task) {
	t.taskList <- task
}

func (t *TaskManager) ExecuteTasks() {
	// Maybe add a sync.Once here as well which ensures only one reader for the queue
	// or make this function non-public and start this in a init() method
	go func() {
		for {
			task := <-t.taskList
			err := task.Execute()
			if err != nil {
				println("failed to execute task")
			} else {
				println("executed task")
			}
		}
	}()
}
