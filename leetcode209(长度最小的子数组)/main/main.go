package main

import (
	"fmt"
)

//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 105

func minSubArrayLen(target int, nums []int) int {
	left, right, sum, res := 0, 0, 0, 0
	for right <= len(nums)-1 { // 右指针最大长度为 len(nums) - 1
		sum += nums[right]
		if sum >= target {
			if res == 0 { // 首次被赋值
				res = right - left + 1
			} else { // 比较一下目前的最小长度
				res = min(res, right-left+1)
			}
			// 开始动左指针
			for left <= right {
				sum -= nums[left]
				left++
				if sum >= target { // 一定不是首次了，这时候再比较一下目前的最小长度
					res = min(res, right-left+1)
				} else {
					right++
					break
				}
			}
		} else {
			right++ // 右指针右移动
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
