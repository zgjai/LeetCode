package Array

import "fmt"

/**
 * Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
 *
 * For example,
 * Given [100, 4, 200, 1, 3, 2],
 * The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.
 *
 * Your algorithm should run in O(n) complexity.
 */

//解题思路：
//借助map，遍历输入数组时，查找数值相邻元素，求出最大连续长度
//但是这种方案的最坏时间复杂度是O(n*n)
//

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int, len(nums))
	maxLength := 0
	for _, num := range nums {
		if length := queryMaxLengthByNum(m, num); length > maxLength {
			maxLength = length
			fmt.Println("maxLength", maxLength)
		}
		m[num] = 0
	}
	return maxLength
}

func queryMaxLengthByNum(m map[int]int, num int) int {
	return upperMaxLength(m, num, 1) + lowerMaxLength(m, num, 1) - 1
}

func upperMaxLength(m map[int]int, num int, length int) int {
	if _, ok := m[num+1]; !ok {
		return length
	}
	length = upperMaxLength(m, num+1, length+1)
	fmt.Println("upperMaxLength", length)
	return length
}

func lowerMaxLength(m map[int]int, num int, length int) int {
	if _, ok := m[num-1]; !ok {
		return length
	}
	length = lowerMaxLength(m, num-1, length+1)
	fmt.Println("lowerMaxLength", length)
	return length
}
