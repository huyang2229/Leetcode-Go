/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// method1
// func isAnagram1(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	map1 := map[rune]int{}
// 	map2 := map[rune]int{}
// 	for _, key := range s {
// 		map1[key] += 1
// 	}
// 	for _, key := range t {
// 		map2[key] += 1
// 	}

// 	for key, value := range map1 {
// 		if value != map2[key] {
// 			return false
// 		}
// 	}

// 	return true
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var alphabetArr [26]int
	for i, value := range s {
		alphabetArr[value-'a']++
		alphabetArr[t[i]-'a']--
	}

	for _, value := range alphabetArr {
		if value != 0 {
			return false
		}
	}
	return true
}

