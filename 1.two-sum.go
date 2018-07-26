/*
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (38.59%)
 * Total Accepted:    999.9K
 * Total Submissions: 2.6M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 *
 *
 */
func twoSum(nums []int, target int) []int {
	res := []int{}
	for i1, v1 := range nums {
		for i2, v2 := range nums {
			if (i1 != i2) && (v1+v2 == target) {
				res = []int{i1, i2}
			}
		}
	}
	return res
}
