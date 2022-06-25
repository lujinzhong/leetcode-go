package main

//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//1 <= digits.length <= 100
//0 <= digits[i] <= 9

func plusOne(digits []int) []int {
	l := len(digits) - 1
	if digits[l] == 9 { // 最后一个是不是9
		digits[l] = 0
		for j := l - 1; j >= 0; j-- { // 这个时候要区分开，前一位是不是9，是的话，继续归0进1
			if digits[j] == 9 {
				digits[j] = 0
			} else {
				digits[j] = digits[j] + 1
				return digits
			}
		}
		// 这个时候没有返回，说明可能是 999 的结构，相当于 1+ 000
		digits = make([]int, l+1)
		digits[0] = 1
	} else {
		digits[l] = digits[l] + 1
		return digits
	}
	return []int{}

}
func main() {
}
