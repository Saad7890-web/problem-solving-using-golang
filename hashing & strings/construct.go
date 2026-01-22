package main

func canConstruct(r string, m string) bool {
	freq := make([]int, 26)

	for _, c := range m {
		freq[c-'a']++
	}

	for _, c := range r {
		freq[c - 'a']--
		if freq[c - 'a'] < 0 {
			return false
		}
	}
	return true
}