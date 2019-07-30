package workerpool

type Args map[string]interface{}
type Work func(args Args) (interface{}, error) // The function which user can define
type Result interface{}

type Task struct {
	state  TaskState
	Id     int
	Error  error
	Args   Args
	result Result
	Func   Work
}

func NewTask(args Args, f Work) *Task {
	return &Task{
		state:  Pending,
		Id:     0,
		Args:   args,
		result: nil,
		Error:  TaskExecutionError,
		Func:   f,
	}
}

func (t *Task) Execute() (interface{}, error) {
	args := t.Args
	f := t.Func
	r, err := f(args)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *Task) Result() (Result, error) {
	if t.state == Finished {
		return t.result, nil
	}
	if t.state == Cancelled {
		return nil, TaskCancelledError
	}
	if t.state == Pending {
		return nil, TaskPendingError
	}
	if t.state == Running {
		return nil, TaskRunningError
	}
	return nil, t.Error
}

func (t *Task) Cancel() bool {
	if t.state == Pending {
		t.state = Cancelled
		return true
	}
	return false
}

func (t *Task) Done() bool {
	return t.state == Finished || t.state == Cancelled
}

func (t *Task) Cancelled() bool {
	return t.state == Cancelled
}

func (t *Task) Running() bool {
	return t.state == Running
}
