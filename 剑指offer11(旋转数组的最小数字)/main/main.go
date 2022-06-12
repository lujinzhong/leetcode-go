package main

import "fmt"

//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//
//给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
//
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

// 双指针法
//执行用时： 4 ms , 在所有 Go 提交中击败了 82.15% 的用户
//内存消耗：3 MB, 在所有 Go 提交中击败了 96.99%的用户
func minArray(numbers []int) int {
	if len(numbers) == 1 { // 只有一个元素，就是它自己了
		return numbers[0]
	}
	// 旋转后，实际上分为两个有序区间，例如 [3,4,5,1,2] => [3,4,,5] + [1,2] 用双指针，一个从左边遍历，一个从右边遍历，如果遍历到突然无序了，那么无序前的值就是最小值
	i := 0
	j := len(numbers) - 1
	for i <= j {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
		if numbers[j] < numbers[j-1] {
			return numbers[j]
		}
		i++
		j--
	}
	// 此时是完全有序的数组，例如 [1,2,3,4,5] 直接返回第一个
	return numbers[0]
}

// 二分法 + 暴力，此题因为有重复数字，无法直接使用二分法
// 使用二分法的条件是：1 通过下标可以直接获取到值 2 通过指定的下标值，可以知道查找值在指定下标值的哪个范围，此题 2 不满足
//执行用时：4 ms, 在所有 Go 提交中击败了 82.15% 的用户
//内存消耗：3 MB, 在所有 Go 提交中击败了56.44%的用户
func minArrayV2(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	low := 0
	high := len(numbers) - 1
	for high > low {
		mid := (high + low) / 2
		if numbers[mid] > numbers[high] {
			low = mid + 1 // 因为这个值比 high 下标的值大，那么这个值不可能是最小的，直接放弃掉
		} else if numbers[mid] == numbers[high] {
			high-- // 由于最右边一定大于等于左边的，所以可以缩小范围
		} else {
			high = mid // 缩小范围
		}
	}
	return numbers[high]
}

func main() {
	fmt.Println(minArrayV2([]int{4, 4, 5, 6, 7, 1, 2, 3}))
	fmt.Println(minArrayV2([]int{1, 2, 3, 4}))
	fmt.Println(minArrayV2([]int{1}))
	fmt.Println(minArrayV2([]int{1, 1, 1, 1}))

}
