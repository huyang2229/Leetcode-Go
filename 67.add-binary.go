func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	ret := []byte{}
	// 是否进位标志
	in := false
	if aLen > bLen {
		i := aLen
		j := bLen
		for i > 0 {
			if j > 0 {
				ret[i], in = calculateRule(string(a[i]), string(b[j]), in)
			} else {
				ret[i], in = calculateRule(string(a[i]), string('0'), in)
			}
			i--
			j--
		}
	} else {
		i := bLen
		j := aLen
		for i > 0 {
			if j > 0 {
				ret[i], in = calculateRule(string(b[i]), string(a[j]), in)
			} else {
				ret[i], in = calculateRule(string(b[i]), string('0'), in)
			}
			i--
			j--
		}
	}
	if in {
		ret = append([]byte{1}, ret...)
	}
	return string(ret)
}

func calculateRule(a string, b string, in bool) (byte, bool) {
	if a == "1" && b == "1" && in {
		return byte('1'), true
	} else if a == "1" && b == "1" && !in {
		return byte('0'), true
	} else if (a == "1" && b == "0" && !in) || (a == "0" && b == "1" && !in) {
		return byte('1'), false
	} else if (a == "1" && b == "0" && in) || (a == "0" && b == "1" && in) {
		return byte('0'), true
	} else {
		return byte('0'), false
	}
}