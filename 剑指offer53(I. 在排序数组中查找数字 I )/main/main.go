package main

import "fmt"

//
//统计一个数字在排序数组中出现的次数。
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2
//提示：
//
//0 <= nums.length <= 105
//-109<= nums[i]<= 109
//nums是一个非递减数组
//-109<= target<= 109

// 二分法+循环
func search(nums []int, target int) int {
	//先二分查找
	if len(nums) == 0 {
		return 0
	}
	low := 0
	high := len(nums) - 1
	index := -1
	count := 0

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			index = mid
			count = 1
			break
		}
	}
	if index == -1 { // 没有找到
		return 0
	}
	index2 := index
	// 由于是有序的，从找到的值这里左右去移动判断
	for {
		index++
		if index > len(nums)-1 {
			break
		}
		if nums[index] == target {
			count++
		} else {
			break
		}
	}

	for {
		index2--
		if index2 < 0 {
			break
		}
		if nums[index2] == target {
			count++
		} else {
			break
		}
	}
	return count
}

// 二分求差值法
func searchV2(nums []int, target int) int {
	// 先找到目标最左边的下标
	leftIndex := binarySearch(nums, target)
	// 再找到目标最右边的下标
	rightIndex := binarySearch(nums, target+1) - 1
	// 两个下标的差值就是结果， 这里要判断特殊情况，就是没查到的时候和索引溢出的情况
	if (leftIndex < 0 || leftIndex > len(nums)-1) || (rightIndex > len(nums)-1 || rightIndex < 0) || nums[leftIndex] != target || nums[rightIndex] != target {
		return 0
	}
	// 正常情况
	return rightIndex - leftIndex + 1
}

// 查出符合目标值的最左边的元素
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low := 0
	high := len(nums) - 1
	ans := len(nums)
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid - 1
			ans = mid
		} else {
			low = mid + 1
		}
	}
	return ans // 如果返回 len(nums) 说明没有查到值
}

func main() {
	//fmt.Println(searchV2([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchV2([]int{0, 0, 1, 2, 3, 3, 4}, 2))
	//fmt.Println(searchV2([]int{1}, 1))

}
