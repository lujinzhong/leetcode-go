package main

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//例如，121 是回文，而 123 不是。
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	res := strconv.Itoa(x)
	var newStr string
	for _, v := range res {
		if newStr == "" {
			newStr = string(v)
		} else {
			newStr = string(v) + newStr
		}
	}
	if newStr == res {
		return true
	}
	return false
}

// 直接反转数字
func isPalindromeV2(x int) bool {
	// 2312
	// 121
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	num := x
	res := 0
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	return x == num
}

// 只反转一半的数字
func isPalindromeV3(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	num := x
	res := 0
	for num > res {
		res = res*10 + num%10
		num = num / 10
	}
	// 此时如果是偶数，则是相等，如果是奇数位，则结果/10 再判断
	return res == num || res/10 == num
}

func main() {
	fmt.Println(isPalindromeV2(121))
}
