package main

import (
	"fmt"
)

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//限制：
//
//0 <= n <= 1000
//0 <= m <= 1000

// 这道题最简单的做法就是两个 for 循环，时间为 O（n*m)
// 第二种做法是，每一行数组都搞个二分，时间为 O （nlogm)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 如果是个空数组，直接GG
	if len(matrix) == 0 {
		return false
	}
	length := len(matrix)
	for i := 0; i < length; i++ {
		if len(matrix[i]) == 0 {
			continue
		}
		high := len(matrix[i]) - 1
		if matrix[i][high] < target { // 如果最大的都比目标小，就没有必要再查了
			continue
		}
		// 走二分
		res := binarySearch(matrix[i], target)
		if res == -1 {
			continue
		} else {
			return true
		}
	}
	return false
}

// 线性探测法
//执行用时：12 ms, 在所有 Go 提交中击败了 99.35% 的用户
//内存消耗： 6.5 MB, 在所有 Go 提交中击败了 53.68% 的用户
func findNumberIn2DArrayV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix) - 1
	n := len(matrix[0]) - 1
	i := 0
	cur := matrix[m][i]
	for {
		if cur > target {
			m--
		} else if cur < target {
			i++
		} else {
			return true
		}
		if i > n || m < 0 {
			return false
		}
		cur = matrix[m][i]
	}
}

// 标准二分法
func binarySearch(arr []int, val int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (high + low) / 2 // low+(high-low)/2 这种可以避免溢出
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Println(findNumberIn2DArray([][]int{{-1}, {-1}}, 0))
	fmt.Println(findNumberIn2DArray([][]int{{}, {}}, 0))
	fmt.Println(findNumberIn2DArray([][]int{{1, 4}, {2, 5}}, 2))
	fmt.Println(findNumberIn2DArray([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, 19))
}
