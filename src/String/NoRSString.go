package String

import "fmt"

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
//another one:
//依然借助于map来实现，同时需要两个下标的index1，index2
//遍历list，index1先保持不动，index2向后移动，依次取出value，并判断value是否存在于map中，不存在则当作key put进去，否则
//计算两个index的差，与当前保存的最长长度比较，保留大者，清空map，然后将index2赋值给index1，开始新一轮的遍历
func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	//最长长度，若最终该值依然为0，表示字符串中没有重复的字符
	var maxLen int = 0
	firInd, secInd := 0, 0
	m := make(map[string]int)
	for ; secInd < len(s); secInd++ {
		char := string(s[secInd])
		i, ok := m[char]
		if !ok {
			m[char] = secInd
			continue
		}
		fmt.Println(firInd, secInd)
		m = make(map[string]int)
		m[char] = secInd
		tmpLen := secInd - firInd
		firInd = i + 1
		if tmpLen > maxLen {
			maxLen = tmpLen
		}
		fmt.Println("maxLen: ", maxLen)
	}
	tmpLen := secInd - firInd
	if tmpLen > maxLen {
		maxLen = tmpLen
	}
	return maxLen
}
