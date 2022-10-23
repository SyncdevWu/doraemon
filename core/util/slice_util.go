package util

func IsNilOrEmpty[T any](b []T) bool {
	if b == nil || len(b) == 0 {
		return true
	}
	return false
}
