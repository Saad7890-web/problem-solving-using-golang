package main

import "sort"

func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string)

	for _, str := range strs {
		chars := []rune(str)

		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)

		anagramMap[key] = append(anagramMap[key], str)

	}

	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}