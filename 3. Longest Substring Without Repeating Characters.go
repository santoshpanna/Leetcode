package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	// sliding window approach
	var hashMap = make(map[int32] int)

	var winLength, maxResult = 0, 0

	for idx, char := range s {
		if _, ok := hashMap[char]; ok {
			winLength = max(winLength, hashMap[char] + 1)
		}
		hashMap[char] = idx

		maxResult = max(maxResult, idx - winLength + 1)

	}

	return maxResult
}

func main() {
	var result = lengthOfLongestSubstring("pwwkew")

	fmt.Print(result)
}
