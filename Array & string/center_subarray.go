package main

func countCenteredSubarrays(nums []int) int {
	n := len(nums)
	count := n 

	for i := 0; i < n; i++ {
		sum := 0
		elements := make(map[int]bool)

		for j := i; j < n; j++ {
			sum += nums[j]
			elements[nums[j]] = true

			
			if j > i {
				if elements[sum] {
					count++
					break 
				}
			}
		}
	}

	return count
}