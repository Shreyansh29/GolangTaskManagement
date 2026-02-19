package taskmanagement

type Job interface {
	// Job Operations
	InitializeJob()
	StartJobExecution()
	CompleteJob()
	CancelJob()
	JobTimeout()
	EnqueueTasks(tasks []Task)

	// Get Job Details
	GetJobDetails()
	GetJobStatus()
	GetJobResult()
}
