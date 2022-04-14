package lib

func Similarity(str, pattern string) float64 {
	jumpTable := KmpTableGenerator(pattern)
	p := 0
	highestMatch := 0.0

	for i := 0; i < len(str); {
		var matchNum int 

		if str[i] == pattern[p] {
			p++
			matchNum = p

			if p == len(pattern){
				return 1.0
			}

			i++
			} else {
			matchNum = p
			j := p + 1

			for k := i + 1; j < len(pattern) && k < len(str); k++ {
				if str[k] == pattern[j] {
					matchNum++
				}
				
				j++
			}
			
			if p == 0 {
				i++
			}
			
			p = jumpTable[p]
		}

		newSimilarity := float64(matchNum)/float64(len(pattern))
		if newSimilarity > highestMatch {
			highestMatch = newSimilarity
		}
	}

	return highestMatch
}
