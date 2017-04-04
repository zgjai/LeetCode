package Array

import "testing"

func TestEmpty(t *testing.T) {
	arr := []int{}
	maxL := longestConsecutive(arr)
	expL := 0
	if maxL != expL {
		t.Error("result is err")
	}
}

func TestCommon(t *testing.T) {
	arr := []int{100, 4, 200, 1, 3, 2}
	maxL := longestConsecutive(arr)
	expL := 4
	if maxL != expL {
		t.Error("result is err")
	}
}

func TestCommon2(t *testing.T) {
	arr := []int{100, 4, 200, 1, 3, 2, 6, 5, 11, 10, 29, 30, 89, 79, 90, 89, -1, -2, -3, -0, 10, 90, 119, 89, 2130, 32, 54, 07, 213, 456, 78, 21, 22, 23, 24, 56, 78, 23}
	maxL := longestConsecutive(arr)
	expL := 11
	if maxL != expL {
		t.Error("result is err")
	}
}
