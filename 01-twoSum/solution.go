package twoSum

/**
 * https://leetcode.com/problems/two-sum/description/
 *
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 */

/**
 * @param {[]int} nums
 * @param {int} target
 * @return {[]int}
 */
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}
