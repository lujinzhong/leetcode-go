package main

import "fmt"

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

//限制：
//0 <= s 的长度 <= 10000

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：3.2 MB, 在所有 Go 提交中击败了12.16%的用户
func replaceSpace(s string) string {
	newStr := ""
	if len(s) == 0 {
		return s
	}
	for _, v := range s {
		if string(v) == " " {
			newStr += "%20"
		} else {
			newStr += string(v)
		}
	}
	return newStr
}

func main() {
	fmt.Println(replaceSpace("we are one"))
}
