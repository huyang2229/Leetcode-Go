/*
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (24.87%)
 * Total Accepted:    538.8K
 * Total Submissions: 2.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 * Examples:
 *
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 *
 * Given "bbbbb", the answer is "b", with the length of 1.
 *
 * Given "pwwkew", the answer is "wke", with the length of 3. Note that the
 * answer must be a substring, "pwke" is a subsequence and not a substring.
 *
 */
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return sLen
	}

	i, j := 0, 1
	max := 0
	for j < length {
		index, isRepeat := contain(s, i, j)
		if isRepeat {
			if newLen := j - i; max < newLen {
				max = newLen
			}
			i = index + 1
		}
		j++
	}
	if newLen := j - i; max < newLen {
		max = newLen
	}
	return max
}

func contain(s string, i, j int) (int, bool) {
	str := s[i:j]
	for index, val := range str {
		if byte(val) == str[j] {
			return index + i, true
		}
	}
	return 0, false
}


