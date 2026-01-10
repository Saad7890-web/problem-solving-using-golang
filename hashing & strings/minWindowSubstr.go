package main

func minWindow(s string, t string) string {
    	if len(t) > len(s) {
		return ""
	}

	
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	required := len(t)
	minLen := len(s) + 1
	start := 0

	
	for right := 0; right < len(s); right++ {
		ch := s[right]

		if need[ch] > 0 {
			required--
		}
		need[ch]--

	
		for required == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			leftChar := s[left]
			need[leftChar]++
			if need[leftChar] > 0 {
				required++
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[start : start+minLen]
}