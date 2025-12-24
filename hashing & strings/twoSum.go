package main

func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)

    for i, num := range nums{
        need := target-num
        if idx, found := indexMap[need]; found {
            return []int{idx, i}
        }
        indexMap[num] = i
    }

    return nil
}