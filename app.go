package todo

type DataBaseService interface {
	AddTask(task string) (int, error)
	GetTask(key int) string
	ListTasks() ([]Task, error)
	RemoveTask(key int) error
	DoTask(key int) (string, error)
	CompletedTasks() ([]Task, error)
}

type Task struct {
	Key   int
	Value string
}
