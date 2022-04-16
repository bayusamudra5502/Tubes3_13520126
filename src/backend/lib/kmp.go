package lib

func Kmp(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	prefix := PrefixFunction(pattern)
	i, j := 0, 0
	for i < len(text) {
		if text[i] != pattern[j] {
			if j != 0 {
				j = prefix[j-1]
			} else {
				i++
			}
		} else {
			i++
			j++
		}
		if j == len(pattern) {
			return i - len(pattern)
		}
	}
	return -1
}
