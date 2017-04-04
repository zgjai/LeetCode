package Array

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 1, 5, -1, 8, 10, 20}
	target := 19
	arr, ok := twoSum(nums, target)
	result := []int{3, 6}
	if ok && !isEqual(arr, result) {
		t.Error("result is err")
	}
}

func TestNonTwoSum(t *testing.T) {
	nums := []int{}
	target := 0
	arr, ok := twoSum(nums, target)
	result := []int{}
	if ok && !isEqual(arr, result) {
		t.Error("result is err")
	}
}

func isEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
