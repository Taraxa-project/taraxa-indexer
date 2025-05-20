package common

import (
	"reflect"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func GetFunctionName(temp any) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

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
		log.WithError(exec_err).WithField("func", GetFunctionName(t.f)).Error("Task fn returned an error")
		*t.err = exec_err
	}
}

type TaskWithoutParams struct {
	f   func() error
	err *error
}

func MakeTaskWithoutParams(f func() error, err *error) *TaskWithoutParams {
	t := new(TaskWithoutParams)
	t.f = f
	t.err = err
	return t
}

func (t *TaskWithoutParams) Run() {
	exec_err := t.f()
	if exec_err != nil {
		log.WithError(exec_err).WithField("func", GetFunctionName(t.f)).Error("TaskWithoutParams fn returned an error")
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
	start := time.Now()
	*t.result, exec_err = t.f(t.params)
	elapsed := time.Since(start)
	// log execution time
	log.WithFields(log.Fields{"func": GetFunctionName(t.f), "elapsed": elapsed, "params": t.params}).Debug("TaskWithResult fn execution time")
	if exec_err != nil {
		log.WithFields(log.Fields{"func": GetFunctionName(t.f), "params": t.params}).WithError(exec_err).Error("TaskWithResult fn returned an error")
		*t.err = exec_err
	}
}
