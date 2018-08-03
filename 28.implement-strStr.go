/*
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (29.59%)
 * Total Accepted:    298.5K
 * Total Submissions: 1M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}
	i, index := 0, 0
	for i < haystackLen {
		if index == needleLen {
			return i - index
		}
		if haystack[i] == needle[index] {
			index++
		} else {
			i -= index
			index = 0
		}
		i++
	}
	if index == needleLen {
		return i - index
	}
	return -1
}
