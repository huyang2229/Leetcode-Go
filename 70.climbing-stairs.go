/*
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (41.77%)
 * Total Accepted:    277.8K
 * Total Submissions: 663.7K
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 * Note: Given n will be a positive integer.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 */
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	res := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n]
}