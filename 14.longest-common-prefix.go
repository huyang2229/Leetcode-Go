/*
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (31.78%)
 * Total Accepted:    300.3K
 * Total Submissions: 945K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	} else if length == 1 {
		return strs[0]
	}
	var res []byte = []byte(strs[0])
	for i := 1; i < length; i++ {
		ele := []byte(strs[i])
		if res = findCommonPrefix(res, ele); len(res) == 0 {
			return ""
		}
	}
	return string(res)
}

func findCommonPrefix(a, b []byte) []byte {
	lenA, lenB, index := len(a), len(b), 0
	if lenA == 0 || lenB == 0 {
		return []byte("")
	}
	for i, v := range a {
		if i >= lenB {
			break
		}
		if v != b[i] && i == 0 {
			return []byte("")
		} else if v != b[i] {
			return b[0:i]
		}
		index = i
	}
	return b[:index+1]
}
