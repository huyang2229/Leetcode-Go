/*
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (29.31%)
 * Total Accepted:    258.6K
 * Total Submissions: 880.3K
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */
func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	if x == 1 {
		return 1
	}
	low, high := 0, x
	for high-low > 1 {
		mid := int(float64(high)/2 + float64(low)/2)
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			high = mid
		} else {
			low = mid
		}
	}
	return low
}
