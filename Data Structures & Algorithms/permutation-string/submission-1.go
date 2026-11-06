func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	
	s1Freq := make([]int, 26)
	s2Freq := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Freq[s1[i] - 'a']++
		s2Freq[s2[i] - 'a']++
	}

	var matches int
	for i := 0; i < 26; i++ {
		if s1Freq[i] == s2Freq[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		index := s2[r] - 'a'
		s2Freq[index]++
		if s1Freq[index] == s2Freq[index] {
			matches++
		} else if s1Freq[index]+1 == s2Freq[index] {
			matches--
		}

		index = s2[l] - 'a'
		s2Freq[index]--
		if s1Freq[index] == s2Freq[index] {
			matches++
		} else if s1Freq[index]-1 == s2Freq[index] {
			matches--
		}
		l++
	}
	return matches == 26
}
