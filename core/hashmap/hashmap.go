package hashmap

import (
	"encoding/json"
	"errors"
)

type HashMap[K comparable, V any] struct {
	raw map[K]V
}

func NewHashMap[K comparable, V any](op *Options[K, V]) *HashMap[K, V] {
	op.init()
	return &HashMap[K, V]{
		raw: op.m,
	}
}

func (w *HashMap[K, V]) GetRaw() map[K]V {
	if w.IsNil() {
		return nil
	}
	return w.raw
}

func (w *HashMap[K, V]) size() int {
	if w.IsNil() {
		return 0
	}
	return len(w.GetRaw())
}

func (w *HashMap[K, V]) IsNilOrEmpty() bool {
	if w.IsNil() || w.size() == 0 {
		return true
	}
	return false
}

func (w *HashMap[K, V]) IsNil() bool {
	if w == nil {
		return true
	}
	return false
}

func (w *HashMap[K, V]) IsNotNil() bool {
	return !w.IsNil()
}

func (w *HashMap[K, V]) Get(key K) (V, bool) {
	if w.IsNil() {
		return nil, false
	}
	value, exists := w.raw[key]
	return value, exists
}

func (w *HashMap[K, V]) Put(key K, value V) (V, int) {
	if w.IsNil() {
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

func (w *HashMap[K, V]) Remove(key K) (V, bool) {
	if w.IsNotNil() {
		return nil, false
	}
	old, exists := w.raw[key]
	delete(w.raw, key)
	return old, exists
}

func (w *HashMap[K, V]) ContainsKey(key K) bool {
	if w.IsNotNil() {
		return false
	}
	_, exists := w.raw[key]
	return exists
}

func (w *HashMap[K, V]) PutAll(m map[K]V) int {
	if w.IsNil() {
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

func (w *HashMap[K, V]) Clear() {
	if w.IsNil() {
		return
	}
	w.raw = make(map[K]V)
}

func (w *HashMap[K, V]) Keys() []K {
	if w.IsNil() {
		return make([]K, 0)
	}
	result := make([]K, len(w.raw))
	for key, _ := range w.raw {
		result = append(result, key)
	}
	return result
}

func (w *HashMap[K, V]) Values() []V {
	if w.IsNil() {
		return make([]V, 0)
	}
	result := make([]V, len(w.raw))
	for _, val := range w.raw {
		result = append(result, val)
	}
	return result
}

func (w *HashMap[K, V]) Entries() ([]K, []V) {
	if w.IsNil() {
		return make([]K, 0), make([]V, 0)
	}
	keys := make([]K, len(w.raw))
	values := make([]V, len(w.raw))
	for key, val := range w.raw {
		keys = append(keys, key)
		values = append(values, val)
	}
	return keys, values
}

func (w *HashMap[K, V]) PutIfAbsent(key K, value V) bool {
	if w.IsNil() {
		return false
	}
	if _, exists := w.raw[key]; exists {
		return false
	}
	w.raw[key] = value
	return true
}

func (w *HashMap[K, V]) getOrDefault(key K, defaultValue V) V {
	if w.IsNil() {
		return defaultValue
	}
	if value, exists := w.raw[key]; exists {
		return value
	} else {
		return defaultValue
	}
}

func (w *HashMap[K, V]) MarshalJSON() ([]byte, error) {
	if w.IsNil() {
		return nil, errors.New("map is nil")
	}
	return json.Marshal(w.raw)
}

func (w *HashMap[K, V]) UnmarshalJSON(b []byte) error {
	if w.IsNil() {
		return errors.New("map is nil")
	}
	w.Clear()
	return json.Unmarshal(b, &w.raw)
}
