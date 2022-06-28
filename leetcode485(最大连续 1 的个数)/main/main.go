package main

//给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
//1 <= nums.length <= 105
//nums[i] 不是 0 就是 1.

func findMaxConsecutiveOnes(nums []int) int {
	num, max := 0, 0
	for _, v := range nums {
		if v == 1 {
			num++
		} else {
			// 判断当前值和最大值的区别
			if num > max {
				max = num
			}
			num = 0
		}
	}
	l := len(nums) - 1
	if nums[l] == 1 && num > max {
		return num
	}
	return max
}

func main() {
}
