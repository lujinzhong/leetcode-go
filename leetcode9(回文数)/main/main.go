package main

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//例如，121 是回文，而 123 不是。
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

import "strconv"

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

func main() {
}
