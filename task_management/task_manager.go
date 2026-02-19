package taskmanagement

type TaskManagement interface {
	// Add a way to execute tasks
	EnqueueTasks(tasks []Task) 
}