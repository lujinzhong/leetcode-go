package main

import "fmt"

//颠倒给定的 32 位无符号整数的二进制位。
func reverseBits(num uint32) uint32 {
	ans := 0
	for i := 1; i <= 32; i++ {
		if num%2 > 0 { // 最后一位是1
			ans += 1 << (32 - i)
		}
		num = num >> 1
	}
	return uint32(ans)
}

func main() {
	fmt.Println(reverseBits(43261596))
}
