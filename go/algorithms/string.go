package algorithms

func StringNaiveMatch(haystack, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle) && j+i < len(haystack); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}
