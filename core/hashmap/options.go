package hashmap

type options[K comparable, V any] struct {
	m map[K]V
}

func (op *options[K, V]) init() {
	if op.m == nil {
		op.m = make(map[K]V)
	}
}
