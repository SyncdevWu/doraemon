package util

func IsNil[K comparable, V any](m map[K]V) bool {
	if m == nil {
		return true
	} else {
		return false
	}
}

func IsNotNil[K comparable, V any](m map[K]V) bool {
	return !IsNil(m)
}

func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}
