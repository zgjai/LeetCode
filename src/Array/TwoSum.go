package Array

import "log"

/**
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution.
 *
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 */

// 解题思路：
// 借助一个map，遍历数组，使用target值减去当前值作为key，从map中取value，存在即说明找到了符合条件的两个值
// 时间为O(n)，空间为O(n)
// 扩展：如果是求someSum呢？
func twoSum(nums []int, target int) ([]int, bool) {
    m := make(map[int]int)
    arr := make([]int, 2)
    ok := false
    for i, num := range nums {
        v, okay := m[num]
        if okay == true {
            arr[0] = v
            arr[1] = i
            ok = true
            break
        }
        m[target-num] = i
    }
    log.Println(arr, ok)
    return arr, ok
}