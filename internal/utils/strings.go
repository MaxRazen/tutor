package utils

func GetLastSegment(str, sep string) string {
	l := len(sep)

	for i := len(str) - 1; i > 0; i-- {
		if str[i-l:i] == sep {
			return str[i:]

		}
	}

	return str
}
