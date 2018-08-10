/*
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (40.86%)
 * Total Accepted:    345.9K
 * Total Submissions: 845.3K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, value := range nums {
		if sum+value < value {
			sum = value
		} else {
			sum += value
		}

		if sum > max {
			max = sum
		}
	}
	return max
}
