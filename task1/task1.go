package task1

type Task1 struct {

}

func(t *Task1) Execute() error{
	print("I am task 1 \n")
	return nil
}
