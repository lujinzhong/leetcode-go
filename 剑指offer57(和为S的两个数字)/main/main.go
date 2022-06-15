package main

import "fmt"

//输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
//输入：nums = [2,7,11,15], target = 9
//输出：[2,7] 或者 [7,2]
//
//限制：
//
//1 <= nums.length <= 10^5
//1 <= nums[i] <= 10^6

// 暴力map
func twoSum(nums []int, target int) []int {
	if len(nums) == 1 {
		return []int{}
	}
	// 把数据放到Map 里，用来判断指定数据是否存储
	tempMap := make(map[int]bool)
	for _, v := range nums {
		tempMap[v] = true
	}
	for _, v := range nums {
		if tempMap[target-v] {
			return []int{v, target - v}
		}
	}
	return []int{}
}

// 双指针
func twoSumV2(nums []int, target int) []int {
	length := len((nums))
	if length == 1 {
		return []int{}
	}
	//双指针
	left := 0
	right := length - 1
	for {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
		if left == right {
			break
		}

	}
	return []int{}
}
func main() {
	fmt.Println(twoSumV2([]int{14, 15, 16, 22, 53, 60}, 76))
}
