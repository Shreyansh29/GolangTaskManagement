package taskmanagement

type Interface job {

	// Job Operations
	InitializeJob()
	StartJobExecution()
	CompleteJob()
	CancelJob()
	JobTimeout()

	// Get Job Details
	GetJobDetails()
	GetJobStatus()
	GetJobResult()
}