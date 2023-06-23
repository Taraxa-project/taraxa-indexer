package utils

import (
	"runtime"

	"github.com/spiretechnology/go-pool"
)

type Task[P any] struct {
	f      func(P) error
	params P
	err    *error
}

func MakeTask[P any](f func(P) error, params P, err *error) *Task[P] {
	t := new(Task[P])
	t.f = f
	t.params = params
	t.err = err
	return t
}

func (t *Task[P]) Run() {
	exec_err := t.f(t.params)
	if exec_err != nil {
		*t.err = exec_err
	}
}

type TaskWithResult[P, R any] struct {
	f      func(P) (R, error)
	params P
	result *R
	err    *error
}

func MakeTaskWithResult[P, R any](f func(P) (R, error), params P, result *R, err *error) *TaskWithResult[P, R] {
	t := new(TaskWithResult[P, R])
	t.f = f
	t.result = result
	t.params = params
	t.err = err
	return t
}

func (t *TaskWithResult[P, R]) Run() {
	var exec_err error
	*t.result, exec_err = t.f(t.params)
	if exec_err != nil {
		*t.err = exec_err
	}
}

// isn't creating threads, but limiting goroutines count. Mostly used for RPC and db related tasks
func MakeThreadPool() pool.Pool {
	return pool.New(uint(runtime.NumCPU()))
}
