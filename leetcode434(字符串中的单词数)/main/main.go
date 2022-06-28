package main

import "fmt"

//统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
//
//请注意，你可以假定字符串里不包括任何不可打印的字符。
//
//示例:
//
//输入: "Hello, my name is John"
//输出: 5
//解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。

func countSegments(s string) int {
	if s == "" {
		return 0
	}

	num, index, str := 0, 0, ""
	l := len(s)
	// 先去掉前面的空格
	for index < l {
		if s[index] == 32 {
			index++
			continue
		} else {
			break
		}
	}
	// 开始计数，每次遇到空格，且下一个不是空格就加一
	for i := index; i <= l-1; i++ {
		if s[i] == 32 && str != "" {
			num++
			str = ""
		} else if s[i] == 32 && str == "" {
			continue
		} else {
			str += string(s[i])
		}
	}
	// 如果最后一个元素不是空格的话，要加一
	if s[l-1] == 32 {
		return num
	}
	return num + 1
}

//计算字符串中单词的数量，就等同于计数单词的第一个下标的个数。因此，我们只需要遍历整个字符串，统计每个单词的第一个下标的数目即可。
//
//满足单词的第一个下标有以下两个条件：
//
//该下标对应的字符不为空格；
//
//该下标为初始下标或者该下标的前下标对应的字符为空格；
func countSegments2(s string) int {
	ans := 0
	for i, ch := range s {
		if (i == 0 || s[i-1] == ' ') && ch != ' ' {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countSegments("a, bb, cc, cc#, "))

}
