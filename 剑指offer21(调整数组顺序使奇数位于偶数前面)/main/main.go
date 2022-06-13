package main

import "fmt"

//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
//
//
//
//示例：
//
//输入：nums = [1,2,3,4]
//输出：[1,3,2,4]
//注：[3,1,2,4] 也是正确的答案之一。
//
//
//提示：
//
//0 <= nums.length <= 50000
//0 <= nums[i] <= 10000

// 双指针，从头开始，原地算法，时间O(n） 空间O(1)
func exchange(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	// 双指针解法
	maxIndex := len(nums) - 1
	p := 0
	q := 1
	// p先定位到左边的第一个偶数，然后和q遇到的第一个奇数来对调位置 [1,2,3,4]
	for {
		if nums[p]%2 != 0 && p < maxIndex { // (nums[i]&1)==1 为奇数
			p++
			continue
		}
		// 如果 p 此时已经是数组最后一个，那么直接返回
		if p == maxIndex {
			return nums
		}
		// 只在p后面找
		q = p + 1
		for {
			if nums[q]%2 == 0 && q < maxIndex { // 是偶数,继续找奇数
				q++
				continue
			}
			break
		}

		// 此时p为偶数，q为奇数，替换下位置
		temp := nums[q]
		nums[q] = nums[p]
		nums[p] = temp
		// 走到底了结束了
		if p == maxIndex || q == maxIndex {
			break
		}
		p++
	}
	return nums
}

// 头尾双指针，一个左边，一个右边
//执行用时：12 ms, 在所有 Go 提交中击败了 99.44% 的用户
//内存消耗：6.4 MB, 在所有 Go 提交中击败了65.19%的用户
func exchangeV2(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	// 双指针解法
	maxIndex := len(nums) - 1
	p := 0
	q := maxIndex
	for p < q {
		// 左指针找偶数，右指针找奇数，找到就替换
		if nums[p]%2 != 0 && p < q {
			p++
		}
		// 找到了偶数
		for {
			if nums[q]%2 == 0 && p < q {
				q--
			}
			break
		}
		// 找到了奇数，交换啦
		temp := nums[q]
		nums[q] = nums[p]
		nums[p] = temp
	}
	return nums
}

func main() {
	fmt.Println(exchangeV2([]int{9, 9, 14, 1, 2, 12, 8, 14, 7, 20}))
	fmt.Println(exchangeV2([]int{1}))
	fmt.Println(exchangeV2([]int{1, 11, 14}))
	fmt.Println(exchangeV2([]int{1, 2, 3, 4}))
}
