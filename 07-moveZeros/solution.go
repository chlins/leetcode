package moveZeros

/**
 * https://leetcode.com/problems/move-zeroes/description/
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
 * For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].
 */

func moveZeroes(nums []int) {
	lastIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			lastIndex++
			nums[i], nums[lastIndex] = nums[lastIndex], nums[i]
		}
	}
}
