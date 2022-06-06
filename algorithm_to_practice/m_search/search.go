package main

import "fmt"

// 标准二分法
func binarySearch(arr [8]int, val int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (high + low) / 2
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
	arr := [8]int{1,2,3,4,5,6,7,8}
	fmt.Println(binarySearch(arr, 8))

}
