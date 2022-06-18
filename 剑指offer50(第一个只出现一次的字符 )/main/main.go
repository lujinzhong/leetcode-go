package main

//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//0 <= s 的长度 <= 50000

// 计数 + 两次循环
func firstUniqChar(s string) byte {
	// 遍历
	var temp [26]int
	for _, v := range s {
		temp[v-'a']++
	}

	for _, v := range s {
		if temp[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}

func main() {

}
