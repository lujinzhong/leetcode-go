package main

//
//给定两个字符串 s 和 t ，它们只包含小写字母。
//
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
//请找出在 t 中被添加的字母。
//0 <= s.length <= 1000
//t.length == s.length + 1
//s 和 t 只包含小写字母

// 1. 遍历 s, 生成map, 再遍历 t,找 map 中不存在的数
// 2. 异或运算，最后留下的一定是单身狗

// 使用异或运算符
// 0^0=0 0^1=1 1^1=0
// a^b^a = (a^a)^b = 0^b = b
func findTheDifference(s string, t string) byte {
	if s == "" {
		return t[0] // 用取下标方式取字符串类型的数据获得的是 byte 类型
	}
	// 位运算异或
	var target rune
	for _, v := range s { // 这里生成的是 rune 类型，相当于 ASCII 码
		target ^= v
	}
	for _, v := range t {
		target ^= v
	}
	return byte(target)
}

func main() {
}
