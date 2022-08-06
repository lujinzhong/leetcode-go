package main

import "fmt"

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5
//0 <= x <= 231 - 1

// 二分法求近似解
func mySqrt(x int) int {
	low, high := 1, x
	for low <= high {
		mid := low + (high-low)/2
		temp := mid * mid
		if temp > x {
			high = mid - 1
		} else if temp <= x && (mid+1)*(mid+1) > x {
			return mid
		} else {
			low = mid + 1
		}

	}
	return 0
}

func main() {
	fmt.Println(mySqrt(4), mySqrt(8), mySqrt(0), mySqrt(1))
}
