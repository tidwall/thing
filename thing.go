package thing

type Thing struct {
	value interface{}
}

func MakeThing(value interface{}) Thing {
	return Thing{value: value}
}

func (t *Thing) Value() interface{} {
	return t.value
}
