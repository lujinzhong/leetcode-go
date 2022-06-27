package main

import "fmt"

//给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
// [0] 1
// [1] 0
// [0,2] 1
// [1,2] 0

// 1. 哈希表法
// 2. 位运算
// 3. 排序后，暴力解法

// 暴力解法，排序后，直接遍历判断下标和值不相等的第一个元素
func missingNumber(nums []int) int {
	nums = bubblingSort(nums)
	for i := 0; i < len(nums); i++ { // 这里可以做个简单优化，如果 nums[0] != 0 直接返回0， 如果 nums[len(nums)-1] == len(nums)-1 则返回 len(nums)
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

// 遍历存哈希，然后找 0~n 不在哈希的值
func missingNumber2(nums []int) int {
	temp := make(map[int]bool)
	for _, v := range nums {
		temp[v] = true
	}

	for i := 0; i <= len(nums); i++ {
		if !temp[i] {
			return i
		}
	}
	return -1
}

// 位运算，异或运算
func missingNumber3(nums []int) int {
	target := 0
	for _, v := range nums {
		target ^= v
	}
	// 再异或 0-n
	for i := 0; i <= len(nums); i++ {
		target ^= i
	}
	return target
}

func bubblingSort(arr []int) []int {
	count := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < count-1; j++ {
			if arr[j] > arr[j+1] { // 判断大小决定是否替换
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
		count-- //这是最关键的，每次处理后，最大的元素都会在最右边，因此后续要处理的数组长度均减一
	}
	return arr
}
func main() {
	fmt.Println(missingNumber([]int{0, 2}))
}
