package task2

type Task2 struct {

}

func(t *Task2) Execute() error{
	print("I am task 2 \n")
	return nil
}
