package hashmap

type Options[K comparable, V any] struct {
	m map[K]V
}

func (op *Options[K, V]) init() {
	if op.m == nil {
		op.m = make(map[K]V)
	}
}
