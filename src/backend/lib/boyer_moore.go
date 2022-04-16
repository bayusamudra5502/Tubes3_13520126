package lib

func BoyerMoore(text, pattern string) int {
	last := BuildLast(pattern)
	textLength := len(text)
	patternLength := len(pattern)
	textIndex := patternLength - 1

	if textIndex > textLength-1 {
		return -1
	}

	patternIndex := patternLength - 1

	for textIndex <= textLength-1 {
		if pattern[patternIndex] == text[textIndex] {
			if patternIndex == 0 {
				return textIndex
			} else {
				textIndex--
				patternIndex--
			}
		} else {
			lo, exists := last[text[textIndex]]
			if !exists {
				lo = -1
			}

			textIndex = textIndex + patternLength - Min(patternIndex, 1+lo)
			patternIndex = patternLength - 1
		}
	}

	return -1
}
