package util

func CheckNull[K comparable, V any](m map[K]V) bool {
	if m == nil {
		return true
	} else {
		return false
	}
}

func CheckNotNull[K comparable, V any](m map[K]V) bool {
	if m == nil {
		return false
	} else {
		return true
	}
}

func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}
