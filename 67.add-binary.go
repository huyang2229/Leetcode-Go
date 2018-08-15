func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	ret := ""
	tmp := ""
	// 是否进位标志
	in := false
	if aLen > bLen {
		i := aLen
		j := bLen
		for i > 0 {
			if j > 0 {
				tmp, in = calculateRule(string(a[i]), string(b[j]), in)
				ret = ret + tmp
			} else {
				tmp, in = calculateRule(string(a[i]), "0", in)
				ret = ret + tmp
			}
			i--
			j--
		}
	} else {
		i := bLen
		j := aLen
		for i > 0 {
			if j > 0 {
				tmp, in = calculateRule(string(b[i]), string(a[j]), in)
				ret = ret + tmp
			} else {
				tmp, in = calculateRule(string(b[i]), "0", in)
				ret = ret + tmp
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
	} else {
		return "0", false
	}
}