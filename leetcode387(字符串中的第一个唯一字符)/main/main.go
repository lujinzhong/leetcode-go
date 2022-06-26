package main

//给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
//1 <= s.length <= 105
//s 只包含小写字母

// 暴力解法-哈希表计数
func firstUniqChar(s string) int {
	// 先遍历放进哈希表计数
	tempMap := make(map[rune]int)
	for _, v := range s {
		tempMap[v] += 1
	}
	// 在遍历一次
	for k, v := range s {
		if tempMap[v] == 1 {
			return k
		}
	}
	return -1
}

func main() {

}
