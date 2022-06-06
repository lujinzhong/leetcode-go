package main

import "fmt"

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

// 递归版本
func binarySearchV2(arr []int, val int) int {
	return binaryRecursiveSearch(arr, 0, len(arr) -1, val)
}

func binaryRecursiveSearch(arr[]int, low, high, val int) int{
	mid := low + (high-low)/2
	if arr[mid] == val {
		return mid
	}else if arr[mid] > val {
		return binaryRecursiveSearch(arr, low, mid - 1, val)
	}else {
		return binaryRecursiveSearch(arr, mid + 1, high, val)
	}
}

func main() {
	fmt.Println(binarySearch([]int{1,2,3,4,5,6,7,8}, 8))
	fmt.Println(binarySearchV2([]int{1,3,5,10,12}, 10))
}
