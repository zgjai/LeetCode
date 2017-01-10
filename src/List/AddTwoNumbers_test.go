package List

import (
	"fmt"
	"testing"
)

func TestNorList(t *testing.T) {
	l1 := construct(2, 4, 3)
	l2 := construct(5, 6, 4)
	expRes := []int{7, 0, 8}
	result := addTwoNumbers(l1, l2)
	judgeResult(result.toSlice(), expRes, t)
}

func TestCarryList(t *testing.T) {
	l1 := construct(8, 3, 7, 0)
	l2 := construct(4, 3, 7, 6)
	expRes := []int{1, 5, 1, 0, 4}
	result := addTwoNumbers(l1, l2)
	judgeResult(result.toSlice(), expRes, t)
}

func TestDifLenList(t *testing.T) {
	l1 := construct(2, 5, 8, 3, 7, 0)
	l2 := construct(4, 3, 7, 6)
	expRes := []int{2, 6, 5, 1, 0, 4}
	result := addTwoNumbers(l1, l2)
	judgeResult(result.toSlice(), expRes, t)
}

func judgeResult(result []int, expRes []int, t *testing.T) {
	fmt.Println(result)
	if result == nil {
		t.Error("result is err")
	}
	if len(result) != len(expRes) {
		t.Error("result is err")
	}

	for i := 0; i < len(expRes); i++ {
		if result[i] != expRes[i] {
			t.Error("result is err")
		}
	}
}
