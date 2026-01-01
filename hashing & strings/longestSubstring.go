package main
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	mx := 0
	left := 0

	for right, ch := range s{
		if idx, found := charIndex[ch]; found && idx >= left {
			left = idx + 1
		}
		charIndex[ch] = right
		if right - left + 1 > mx {
			mx = right - left + 1
		}
	}

	return mx
}