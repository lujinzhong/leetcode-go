package main

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
func sortedSquares(nums []int) []int {
	l := len(nums)
	arr := make([]int, l, l)
	max := l - 1
	left, right := 0, l-1
	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			arr[max] = nums[left] * nums[left]
			left++
			max--
		} else {
			arr[max] = nums[right] * nums[right]
			right--
			max--
		}
	}
	return arr
}
func main() {
}
