package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for index, num := range nums {
		if value, exist := mp[target-num]; exist {
			return []int{value, index}
		}
		mp[num] = index
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 17
	res := twoSum(nums, target)
	fmt.Println(res)
}
