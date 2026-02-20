package taskmanagement

type Job interface {
	// Job Operations
	InitializeJob()
	GenerateTasks(request any)
	StartExecution()
	Complete()
	Cancel()
	Timeout()
	EnqueueTasks(tasks []Task)

	// Get Job Details
	GetJobDetails()
	GetJobStatus()
	GetJobResult()
}
