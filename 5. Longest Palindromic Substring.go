package main


import (
	"fmt"
	shared "github.com/santoshpanna/leetcode/shared"
)

func max(a int, b int) int {
	if a > b{
		return a
	}
	return b

}

func longestPalindrome(s string) string {
	// expand around center approach
	var startIndex, endIndex, maxLength = 0, 0, 0

	if len(s) <= 1 {
		return s
	} else if len(s) == 2 && s[0] == s[1]{
		return s
	} else if len(s) == 2 && s[0] != s[1] {
		return string(s[0])
	}

	for idx := 0; idx < (2*len(s)-1); idx++ {
		var leftIndex = idx / 2
		var rightIndex = leftIndex

		if idx % 2 == 1 {
			leftIndex = (idx - 1) / 2
			rightIndex = (idx + 1) / 2
		}



		//fmt.Println("-------------------")
		//fmt.Println("idx== ", idx)


		var temp = 0

		//fmt.Println("s[leftIndex] == s[rightIndex] == ", leftIndex >= 0 && rightIndex < len(s) && s[leftIndex] == s[rightIndex])
		//fmt.Println("leftIndex < rightIndex == ", leftIndex < rightIndex)
		//fmt.Println("leftIndex >= 0 && rightIndex < len(s) == ", leftIndex >= 0 && rightIndex < len(s))

		for leftIndex >= 0 && leftIndex < len(s) && rightIndex >= 0 && rightIndex < len(s) && s[leftIndex] == s[rightIndex]{
			leftIndex = leftIndex - 1
			rightIndex = rightIndex + 1
		}

		//fmt.Println("out of loop, leftIndex = ", leftIndex, " | rightIndex = ", rightIndex, " | idx = ", idx)

		leftIndex = leftIndex + 1
		rightIndex = rightIndex - 1

		//fmt.Println("out of loop, after index correction, leftIndex = ", leftIndex, " | rightIndex = ", rightIndex, " | idx = ", idx)

		temp = rightIndex - leftIndex
		maxLength = endIndex - startIndex

		//if leftIndex < 0 {
		//	leftIndex = 0
		//}
		//
		//if rightIndex >= len(s) {
		//	rightIndex = len(s) - 1
		//}


		//fmt.Println("maxLength ", maxLength, " temp == ", temp, " wow -- ",  temp > maxLength)
		if temp > maxLength {
			startIndex = leftIndex
			endIndex = rightIndex
			//fmt.Print("Final : maxLength ", maxLength, " temp == ", temp)
			//fmt.Print(" startIndex ", startIndex)
			//fmt.Println(" endIndex ", endIndex)
			temp = 0
		}


		//fmt.Println("-------------------")
	}

	//fmt.Println(" startIndex ", startIndex)
	//fmt.Println(" endIndex ", endIndex)

	return s[startIndex:endIndex+1]

}

func main() {
	var result1 = longestPalindrome("babad")
	fmt.Println(shared.Cyan,"babad == ", result1, " should be == bab", shared.Reset)

	var result2 = longestPalindrome("aaaa")
	fmt.Println(shared.Cyan, "aaaa == ", result2, " should be == aaaa", shared.Reset)

	var result3 = longestPalindrome("aa")
	fmt.Println(shared.Cyan,"aa == ", result3, " should be == aa", shared.Reset)

	var result4 = longestPalindrome("ac")
	fmt.Println(shared.Cyan,"ac == ", result4, " should be == a", shared.Reset)

	var result6 = longestPalindrome("ccc")
	fmt.Println(shared.Cyan,"ccc == ", result6, " should be == ccc", shared.Reset)

	var result5 = longestPalindrome("cbbd")
	fmt.Println(shared.Cyan, "cbbd == ", result5, " should be == bb", shared.Reset)
}
