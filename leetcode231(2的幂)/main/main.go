package main

import "fmt"

//给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
//
//如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。
//
//提示：
//
//-231 <= n <= 231 - 1

// 循环判断
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	ans := 1
	for i := 1; i < 32; i++ {
		ans = ans << 1
		if ans == n {
			return true
		}

	}
	return false
}

// 位运算
func isPowerOfTwo2(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}

func main() {
	fmt.Println(1 & 0)
}
