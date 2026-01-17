package main

func countResidue(s string) int {
	seen := make([]bool, 26)
	ans := 0
	count := 0

	for i:=0; i < len(s); i++{
		index := s[i]-'a'
		if !seen[index]{
			seen[index] = true
			count++;
		}
		pref := i+1
		if count == pref%3{
			ans++;
		}
		
	}
	return ans
}