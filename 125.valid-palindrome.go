/*
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (27.97%)
 * Total Accepted:    255.4K
 * Total Submissions: 905.7K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 */
import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	var pos int
	for _, v := range runes {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			runes[pos] = unicode.ToLower(v)
			pos++
		}
	}
	runes = runes[:pos]
	l := len(runes)
	for i := 0; i < l/2; i++ {
		if runes[i] != runes[l-i-1] {
			return false
		}
	}
	return true
}
