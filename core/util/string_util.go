package util

func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsBlank(str string) bool {
	for _, runeVal := range str {
		if runeVal != ' ' {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
