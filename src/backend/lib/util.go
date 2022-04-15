package lib

func KmpTableGenerator(str string) (result []int) {
	result = append(result, 0)

	for i := 1; i < len(str); i++ {
		start := 0
		end := i - 1

		for start < end && (str[start] == str[end]) {
			start++
			end--
		}

		result = append(result, start)
	}

	return result
}

func BuildLast(pattern string) (result map[byte]int) {
	result = make(map[byte]int)

	for i := 0; i < len(pattern); i++ {
		result[pattern[i]] = i
	}

	return result
}

func Min(first, second int) int {
	if first < second {
		return first
	}
	return second
}
