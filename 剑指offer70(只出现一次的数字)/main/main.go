package main

import (
	"fmt"
)

//给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
//
//
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁵
//
//
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 { // 长度为一，就一个元素，符合
		return nums[0]
	}
	low := 0
	high := len(nums) - 1
	for high >= low {
		mid := low + (high-low)/2
		if mid == 0 || mid == len(nums)-1 { // 如果是最后一个或者第一个，直接返回
			return nums[mid]
		}

		if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
			return nums[mid]
		}
		// 此时要么前面和他一样，要么后面和他一样， 想办法找到他后面的
		if nums[mid] == nums[mid+1] {
			mid++
		}

		// 包括他，如果和前面有偶数个元素，说明不在前面
		if (mid+1)%2 == 0 { //不在前面
			low = mid + 1
		} else { // 在前面
			high = mid - 1
		}

	}
	return -1
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 2, 2}))
	fmt.Println(singleNonDuplicate([]int{1}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 4, 4}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4}))

}
