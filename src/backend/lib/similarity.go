package lib

func Similarity(str, pattern string) float64 {
	jumpTable := KmpTableGenerator(pattern)
	p := 0
	highestMatch := 0.0

	for i := 0; i < len(str); {
		if str[i] == pattern[p] {
			p++

			if p == len(pattern){
				return 1.0
			}

			i++
		} else {
			matchNum := p
			j := p + 1

			if p == 0 {
				i++
			}

			for k := i + 1; j < len(pattern) && k < len(str); k++ {
				if str[k] == pattern[j] {
					matchNum++
				}

				j++
			}

			newSimilarity := float64(matchNum)/float64(len(pattern))
			if newSimilarity > highestMatch {
				highestMatch = newSimilarity
			}

			p = jumpTable[p]
		}
	}

	return highestMatch
}
