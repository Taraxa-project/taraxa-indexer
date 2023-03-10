package indexer

type Task[T any] struct {
	f     func(T) error
	param T
	err   *error
}

func MakeTask[T any](f func(T) error, param T, err *error) *Task[T] {
	t := new(Task[T])
	t.f = f
	t.param = param
	t.err = err
	return t
}

func (t *Task[T]) Run() {
	*t.err = t.f(t.param)
}
