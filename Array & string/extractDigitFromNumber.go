package main

import (
	"fmt"
)

func separateDigits(nums []int) []int {
    result := []int{}

    for _, num := range nums{
        str := fmt.Sprint(num)
        for _, ch := range str {
            result = append(result, int(ch - '0'))
        }
    }
    return result
}

func main(){
	arr := []int{13,25,10,50}
	fmt.Println(separateDigits(arr))
}