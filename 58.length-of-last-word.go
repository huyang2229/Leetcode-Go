/*
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.08%)
 * Total Accepted:    205.9K
 * Total Submissions: 642K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 */
import (
	"unicode"
)

func lengthOfLastWord(s string) int {
	count := 0
	for len(s) > 0 {
		if unicode.IsSpace(rune(s[len(s)-1])) {
			s = s[:len(s)-1]
		} else {
			break
		}
	}

	if s == "" {
		return 0
	}

	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsSpace(rune(s[i])) {
			count++
		} else {
			break
		}
	}

	return count
}
