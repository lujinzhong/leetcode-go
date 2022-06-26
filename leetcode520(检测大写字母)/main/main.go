package main

//我们定义，在以下情况时，单词的大写用法是正确的：
//
//全部字母都是大写，比如 "USA" 。
//单词中所有字母都不是大写，比如 "leetcode" 。
//如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
//给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。

// 统计大小写字母个数，再判断
func detectCapitalUseV2(word string) bool {
	// 1. 全大写
	// 2. 全小写
	// 3. 有大写和小写 （除非大写只有一个且在头，否则返回false）
	up, lower := 0, 0
	for k, _ := range word {
		if isUp(word[k]) {
			up++
		} else {
			lower++
		}
	}
	// 全大写或者全小写
	if up == len(word) || lower == len(word) {
		return true
	}
	// 有大写和小写
	if up == 1 && isUp(word[0]) {
		return true
	}
	return false
}

func detectCapitalUse(word string) bool {
	n := len(word)
	// 首先判断第一个字符是不是大写，是的话，要求剩下的都是小写或者都是大写
	// 不是的话，要求剩下的都是小写
	// A-Z  65-90
	// a-z 97-122

	if isUp(word[0]) { // 是大写
		// 要么都是大写，要么都是小写
		if n >= 2 {
			if isLow(word[1]) { // 剩下的都要是小写
				return isAllLow(word, 2)
			} else { // 剩下的都要是大写
				return isAllUp(word, 2)
			}
		}
	} else { // 是小写,要求剩下的都是小写
		return isAllLow(word, 1)
	}
	return true
}

func isAllUp(s string, start int) bool {
	for i := start; i < len(s); i++ {
		if !isUp(s[i]) {
			return false
		}
	}
	return true
}

func isAllLow(s string, start int) bool {
	for i := start; i < len(s); i++ {
		if !isLow(s[i]) {
			return false
		}
	}
	return true
}

func isUp(r byte) bool {
	return r >= 65 && r <= 90
}

func isLow(r byte) bool {
	return r >= 97 && r <= 122
}

func main() {

}
