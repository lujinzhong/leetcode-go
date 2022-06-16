package main

import "fmt"

//一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
//1 <= 数组长度 <= 10000

// 1. 暴力解法，遍历找到下标值和元素值不相等的元素，为O(n)
// 2. 二分法，找到不符合元素的左边，加一返回
func missingNumber(nums []int) int {
	// 处理边界问题
	if nums[len(nums)-1] == len(nums)-1 { // 全部有序的
		return len(nums)
	}

	//二分法找区间,找到不符合的左边元素
	low := 0
	high := len(nums) - 1
	for high >= low {
		mid := low + (high-low)/2
		if nums[mid] == mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high + 1
}
func main() {
	fmt.Println(missingNumber([]int{1}))
	fmt.Println(missingNumber([]int{0}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3}))

}
