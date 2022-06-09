package main

import "fmt"

//找出数组中重复的数字。
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//限制：
//2 <= n <= 100000

// 遍历 + 哈希
func findRepeatNumber(nums []int) int {
	n := len(nums)
	temp := make(map[int]bool)
	for i := 0; i < n; i++ {
		if temp[nums[i]] {
			return nums[i]
		}
		temp[nums[i]] = true
	}
	return -1
}

// 遍历 + 原地置换
//执行用时：24 ms, 在所有 Go 提交中击败了 98.69% 的用户
//内存消耗：8.5 MB, 在所有 Go 提交中击败了 65.79% 的用户
func findRepeatNumberV2(nums []int) int {
	n := len(nums)
	i := 0 // 指针
	for i < n {
		val := nums[i]
		if val != i {
			// 替换前先判断是不是已经被占坑了
			if nums[val] == val {
				return val
			}
			// 否则替换位置
			temp := nums[val]
			nums[val] = val
			nums[i] = temp
		} else {
			i++
		}
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{1, 2, 4, 2, 4}))
	fmt.Println(findRepeatNumberV2([]int{1, 2, 4, 2, 4}))
}
