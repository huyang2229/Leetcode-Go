/*
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (35.14%)
 * Total Accepted:    221.8K
 * Total Submissions: 628.6K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	ret := ""
	tmp := ""
	// 是否进位标志
	in := false
	if aLen > bLen {
		i := aLen - 1
		j := bLen - 1
		for i >= 0 {
			if j >= 0 {
				tmp, in = calculateRule(string(a[i]), string(b[j]), in)
				ret = tmp + ret
			} else {
				tmp, in = calculateRule(string(a[i]), "0", in)
				ret = tmp + ret
			}
			i--
			j--
		}
	} else {
		i := bLen - 1
		j := aLen - 1
		for i >= 0 {
			if j >= 0 {
				tmp, in = calculateRule(string(b[i]), string(a[j]), in)
				ret = tmp + ret
			} else {
				tmp, in = calculateRule(string(b[i]), "0", in)
				ret = tmp + ret
			}
			i--
			j--
		}
	}
	if in {
		ret = "1" + ret
	}
	return ret
}

func calculateRule(a string, b string, in bool) (string, bool) {
	if a == "1" && b == "1" && in {
		return "1", true
	} else if a == "1" && b == "1" && !in {
		return "0", true
	} else if (a == "1" && b == "0" && !in) || (a == "0" && b == "1" && !in) {
		return "1", false
	} else if (a == "1" && b == "0" && in) || (a == "0" && b == "1" && in) {
		return "0", true
	} else if a == "0" && b == "0" && in {
		return "1", false
	} else {
		return "0", false
	}
}
