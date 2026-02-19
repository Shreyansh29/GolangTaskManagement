package taskmanagement

type Task interface{
	// Task Operations
	InitializeTask()
	StartTaskExecution()
	CompleteTask()
	CancelTask()
	TaskTimeout()

	// Get Task Details
	GetTaskDetails()
	GetTaskStatus()
	GetTaskResult()
}