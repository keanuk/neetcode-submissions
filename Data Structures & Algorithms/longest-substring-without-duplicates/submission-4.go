func lengthOfLongestSubstring(s string) int {
	var res int
	var sub int
	seen := make(map[rune]int)
	for i, c := range s {
		if v, ok := seen[c]; ok && v >= sub {
			sub = v + 1
		}
		seen[c] = i
		res = max(res, i - sub + 1)
	}
	return res
}
