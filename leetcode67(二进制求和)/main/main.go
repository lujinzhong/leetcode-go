package main

import (
	"fmt"
	"strconv"
)

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
//输入为 非空 字符串且只包含数字 1 和 0。

//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。

// 本题由于字符串的长度最大是10000，无法通过转数字实现。 uint64 也就 1>>128 ，会溢出的。

// 模拟二进制加法，逢 2 进 1
func addBinary(a string, b string) string {
	var max func(int, int) int
	max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	var carry int
	ans := ""
	for i := 0; i < n; i++ { // 从低位相加，如果是 1+1 ，结果为0，记进1
		if i < lenA { // 此时 a 还没算完
			carry += int(a[lenA-i-1] - '0') // int('1' - '0')=1 int('0'-'0') = 0
		}
		if i < lenB { // 此时 b 还没算完
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry = carry >> 1 // carry = 3  carry >> 1 = 1 (相当于保留了进位)
	}
	if carry > 0 { // 其实就是等于 1 的时候，补上一位
		ans = "1" + ans
	}
	return ans
}

// 不溢出的时候可以使用下面的暴力解法，先转十进制相加后，再转二进制字符串
func addBinary2(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	l1, l2 := len(a), len(b)
	var count1 uint64
	var count2 uint64
	for i := 0; i < l1; i++ {
		if a[i] == '1' {
			count1 += 1 << (l1 - i - 1)
		}
	}
	for i := 0; i < l2; i++ {
		if b[i] == '1' {
			count2 += 1 << (l2 - i - 1)
		}
	}
	sum := count2 + count1
	// 转为字符串
	str := ""
	for sum > 0 {
		if sum%2 > 0 {
			str = "1" + str
		} else {
			str = "0" + str
		}
		sum = sum >> 1
	}
	return str
}

func main() {
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
