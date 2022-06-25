package main

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 为 无重复元素 的 升序 排列数组
//-104 <= target <= 104
//

// 二分法
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// 没有找到目标值
	// 1,2 (3) 最终mid=high=low=2 res=low
	// 2,3 (1) 最终mid=high=low=-1 res=low+1
	// 1,3,4 (2) 最终mid=high=low=-1 res=low+2
	if low == len(nums) || nums[low] > target { // 这个顺序不能换，否则会溢出
		return low
	} else {
		return low + 1
	}
}

func main() {
}
