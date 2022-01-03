package thing

type Thing[T any] struct {
	value T
}

func MakeThing[T any](value T) Thing[T] {
	return Thing[T]{value: value}
}

func (t *Thing[T]) Value() T {
	return t.value
}
