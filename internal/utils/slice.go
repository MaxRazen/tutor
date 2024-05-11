package utils

func InSlice[T comparable](c T, s []T) bool {
	for _, v := range s {
		if c == v {
			return true
		}
	}
	return false
}
