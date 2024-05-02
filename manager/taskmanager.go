package manager

import (
	"sync"
)

var taskManagerInstance *TaskManager
var taskManagerOnce sync.Once

type Task interface {
	Execute()
}

type TaskManager struct {
	taskList []Task
	lock     sync.RWMutex
}

func GetTaskManager() *TaskManager {
	taskManagerOnce.Do(func() {
		if taskManagerInstance == nil {
			taskManagerInstance = &TaskManager{}
		}
	})
	return taskManagerInstance
}

func (t *TaskManager) Push(task Task) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.taskList = append(t.taskList, task)
	return nil
}

func (t *TaskManager) ExecuteTasks() {
	// Maybe add a sync.Once here as well which ensures only one reader for the queue
	// or make this function non-public and start this in a init() method
	go func() {
		for {
			if len(t.taskList) > 0 {
				t.lock.Lock()
				task := t.taskList[0]
				t.taskList = t.taskList[1:]
				t.lock.Unlock()
				task.Execute()
			}
		}
	}()
}
