package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i int = 0
	var j int = 0
	var merged []int
	for i < len(nums1) || j < len(nums2){
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				merged = append(merged, nums1[i])
				i++
			} else {
				merged = append(merged, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			merged = append(merged, nums1[i])
			i++
		} else if j < len(nums2){
			merged = append(merged, nums2[j])
			j++
		}
	}

	var medianIdx = float64(len(merged) + 1) / 2.0
	medianIdx = medianIdx - 1

	if math.Floor(medianIdx) == math.Ceil(medianIdx) {
		return float64(merged[int(medianIdx)])
	} else {
		return float64(merged[int(math.Floor(medianIdx))] + merged[int(math.Ceil(medianIdx))]) / 2.0
	}

	return float64(0)
}

func main() {
	var result = findMedianSortedArrays([] int {2}, []int {})
	fmt.Print("jjj == ", result)
}
