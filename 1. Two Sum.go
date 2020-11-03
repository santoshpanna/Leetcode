package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var hashMap = make(map[int] int)
	var finalArr []int
	for i:=0; i < len(nums); i++ {
		var diff =  target - nums[i]
		if _, ok := hashMap[diff]; ok {
			return append(finalArr, hashMap[diff], i)
		} else {
			hashMap[nums[i]] = i
		}
	}

	return finalArr
}
func main() {
	var nums[] int
	var target int = 9
	nums = append(nums, 2,7,11,15)
	fmt.Println(twoSum(nums, target))
}
