package main

import "fmt"

//实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//
//说明：
//
//当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
//对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	a := 0
	l1 := len(haystack)
	l2 := len(needle)
	if l2 > l1 || (l1 == l2 && haystack != needle) { //子串比父串大或者是父子串长度相等但是值不相同，返回 -1
		return -1
	}
	for a <= l1-1 { // a 不能大于父串长度，否则溢出
		// 先找到b
		if haystack[a] == needle[0] {
			// 判断剩余的父串长度是否比子串还小，是的话，那后面不用查了
			if l1-a < l2 {
				return -1
			}
			if haystack[a:a+l2] == needle { // 利用切片的原理，获取子串长度的切片判断是否和子串相等，相等则符合，否则父串继续缩小
				return a
			}
		}
		a++
	}
	return -1
}
func main() {
	fmt.Println(strStr("abcdefg", "bcd"))

}
