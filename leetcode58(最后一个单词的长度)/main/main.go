package main

//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
//1 <= s.length <= 104
//s 仅有英文字母和空格 ' ' 组成
//s 中至少存在一个单词

func lengthOfLastWord(s string) int {
	l := len(s) - 1
	ans := 0
	for { // 干掉右侧空格
		if s[l] == 32 {
			l--
		} else {
			break
		}
	}
	for l >= 0 {
		if s[l] != 32 {
			ans++
			l--
		} else {
			return ans
		}
	}
	return ans
}
func main() {
}
