package hashmap

import (
	"encoding/json"
	"errors"
)

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

func (w *MapWrapper[K, V]) Put(key K, value V) (V, int) {
	if w.CheckNull() {
		return nil, 0
	}
	old, exists := w.raw[key]
	w.raw[key] = value
	if exists {
		return old, 0
	} else {
		return old, 1
	}
}

func (w *MapWrapper[K, V]) Remove(key K) (V, bool) {
	if w.CheckNotNull() {
		return nil, false
	}
	old, exists := w.raw[key]
	delete(w.raw, key)
	return old, exists
}

func (w *MapWrapper[K, V]) ContainsKey(key K) bool {
	if w.CheckNotNull() {
		return false
	}
	_, exists := w.raw[key]
	return exists
}

func (w *MapWrapper[K, V]) PutAll(m map[K]V) int {
	if w.CheckNull() {
		return 0
	}
	result := 0
	for k, v := range m {
		if _, exists := w.Get(k); !exists {
			result++
		}
		w.raw[k] = v
	}
	return result
}

func (w *MapWrapper[K, V]) Clear() {
	if w.CheckNull() {
		return
	}
	w.raw = make(map[K]V)
}

func (w *MapWrapper[K, V]) Keys() []K {
	if w.CheckNull() {
		return make([]K, 0)
	}
	result := make([]K, len(w.raw))
	for key, _ := range w.raw {
		result = append(result, key)
	}
	return result
}

func (w *MapWrapper[K, V]) Values() []V {
	if w.CheckNull() {
		return make([]V, 0)
	}
	result := make([]V, len(w.raw))
	for _, val := range w.raw {
		result = append(result, val)
	}
	return result
}

func (w *MapWrapper[K, V]) PutIfAbsent(key K, value V) bool {
	if w.CheckNull() {
		return false
	}
	if _, exists := w.raw[key]; exists {
		return false
	}
	w.raw[key] = value
	return true
}

func (w *MapWrapper[K, V]) getOrDefault(key K, defaultValue V) V {
	if w.CheckNull() {
		return defaultValue
	}
	if value, exists := w.raw[key]; exists {
		return value
	} else {
		return defaultValue
	}
}

func (w *MapWrapper[K, V]) MarshalJSON() ([]byte, error) {
	if w.CheckNull() {
		return nil, errors.New("map is nil")
	}
	return json.Marshal(w.raw)
}

func (w *MapWrapper[K, V]) UnmarshalJSON(b []byte) error {
	if w.CheckNull() {
		return errors.New("map is nil")
	}
	w.Clear()
	return json.Unmarshal(b, &w.raw)
}
