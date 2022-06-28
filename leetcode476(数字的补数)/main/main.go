package main

import (
	"fmt"
)

//对整数的二进制表示取反（0 变 1 ，1 变 0）后，再转换为十进制表示，可以得到这个整数的补数。
//
//例如，整数 5 的二进制表示是 "101" ，取反后得到 "010" ，再转回十进制表示得到补数 2 。
//给你一个整数 num ，输出它的补数。
//
//1 <= num < 231

// 找到最高位，然后的前一位再减一，可以获得连续1的二进制，再和原值异或操作
func findComplement(num int) int {
	count := 0
	n := num
	for n > 0 {
		n = n >> 1
		count++
	}
	m := (1 << count) - 1 // 2^count拿到连续1的二进制
	return num ^ m
}

func main() {
	fmt.Println(findComplement(5))
}
