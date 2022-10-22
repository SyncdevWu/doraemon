package hashmap

type MapWrapper[K comparable, V any] struct {
	raw map[K]V
}

func NewMapWrapper[K comparable, V any](op *options[K, V]) *MapWrapper[K, V] {
	op.init()
	return &MapWrapper[K, V]{
		raw: op.m,
	}
}

func (w *MapWrapper[K, V]) GetRaw() map[K]V {
	if w.CheckNull() {
		return nil
	}
	return w.raw
}

func (w *MapWrapper[K, V]) size() int {
	if w.CheckNull() {
		return 0
	}
	return len(w.GetRaw())
}

func (w *MapWrapper[K, V]) CheckNullOrEmpty() bool {
	if w.CheckNull() || w.size() == 0 {
		return true
	}
	return false
}

func (w *MapWrapper[K, V]) CheckNull() bool {
	if w == nil {
		return true
	}
	return false
}

func (w *MapWrapper[K, V]) CheckNotNull() bool {
	if w == nil {
		return false
	}
	return true
}

func (w *MapWrapper[K, V]) Get(key K) (V, bool) {
	if w.CheckNull() {
		return nil, false
	}
	value, exists := w.raw[key]
	return value, exists
}
