package String

import (
	"testing"
)

func TestNoRepeatingString(t *testing.T) {
	var tS string = "abc"
	l := lengthOfLongestSubstring(tS)
	if l != 3 {
		t.Error("length is incorrect")
	}
}

func TestSameString(t *testing.T) {
	var tS string = "aaaaaaaa"
	l := lengthOfLongestSubstring(tS)
	if l != 1 {
		t.Error("length is incorrect")
	}
}

func TestRepeatingString(t *testing.T) {
	var tS string
	var l int
	tS = "abcabcbb"
	l = lengthOfLongestSubstring(tS)
	if l != 3 {
		t.Error("length is incorrect")
	}

	tS = "pwwkew"
	l = lengthOfLongestSubstring(tS)
	if l != 3 {
		t.Error("length is incorrect")
	}
	tS = "abcabcbbfdsa"
	l = lengthOfLongestSubstring(tS)
	if l != 5 {
		t.Error("length is incorrect")
	}
}

func TestLongString(t *testing.T) {
	var tS string = "fdgregoeogitouojfgghkrjerioguhiuyetuegnjkjnkjndckjvndfjbnkjndgklashdglkhjeiourtghjrkhgnkjasngkhndfkjghnbdkfjgiehrutyhuihrykfjdkjgkasdghfjyehiutrkgjhd"
	l := lengthOfLongestSubstring(tS)
	if l != 14 {
		t.Error("length is incorrect")
	}
}
