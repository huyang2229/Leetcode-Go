/*
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (37.36%)
 * Total Accepted:    363.1K
 * Total Submissions: 971.9K
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	var result int
	flag := false
	for x > 0 {
		result = result*10 + x%10
		if x == result || x/10 == result {
			flag = true
			break
		}
		x = x / 10
	}
	return flag
}
