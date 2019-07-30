package workerpool

type TaskState string

const (
	Pending         = "pending"
	Running         = "running"
	Finished        = "finished"
	Cancelled       = "cancelled"
	FailedWithError = "failed"
)
