package workerpool

import "errors"

var (
	TaskPendingError   = errors.New("task is not available")
	TaskCancelledError = errors.New("task was cancelled")
	TaskRunningError   = errors.New("task is currently running")
	TaskExecutionError = errors.New("task execution Failed")
)
