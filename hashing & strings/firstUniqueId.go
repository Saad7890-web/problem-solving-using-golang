package main

func firstUnique(req []int) int {
	freq := make(map[int]int)

	for _,id := range req{
		freq[id]++
	}

	for _,id := range req {
		if req[id] == 1{
			return id
		}
	}
	return -1
}