package String

import (
	"fmt"
)

/**
 * Given a string, find the length of the longest substring without repeating characters.
 *
 * Examples:
 *
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 * Given "bbbbb", the answer is "b", with the length of 1.
 * Given "pwwkew", the answer is "wke", with the length of 3.
 * Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */


//解题思路:
//借助于map来实现，map中key是字符，value是一个list，list中存放该字符在字符串中的index
//以上思路忽略了重要的问题：比如"abba"这种，得到的结果将会是3...
//sad :(
//
func lengthOfLongestSubstring(s string) int {
	//最长长度，若最终该值依然为0，表示字符串中没有重复的字符
	var maxLen int = 0
	//var firInd, secInd int = 0, len(s)
	fmt.Println(s)

	m := make(map[byte][]int)
	for i:=0; i<len(s); i++ {
		tmpS := s[i]
		//print(tmpS, " ")
		list, ok := m[tmpS]
		fmt.Println("tmpS: ", string(tmpS), "list: ", list)
		if ok {
			//println("tmpS: ", tmpS, "list: ", list)
			last := list[len(list)-1]
			fmt.Println("char: ", string(tmpS), "lastIndex: ", last)
			l := i - last
			if l >= maxLen {
				maxLen = l
				fmt.Println("max: ", maxLen)
			}
			m[tmpS] = append(list, i)
			continue
		}
		list = []int{i}
		m[tmpS] = list
	}

	if maxLen == 0 {
		maxLen = len(s)
	}
	return maxLen
}