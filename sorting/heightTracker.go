package sorting

import "sort"
func heightChecker(heights []int) int {
    exp := make([]int, len(heights))
    copy(exp, heights)
    sort.Ints(exp)
    cnt := 0
    for i:=0; i < len(heights); i++{
        if exp[i] != heights[i]{
            cnt++
        } 
    }
    return cnt
}