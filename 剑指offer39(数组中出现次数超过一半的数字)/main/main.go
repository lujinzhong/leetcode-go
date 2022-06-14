package main

//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
//输出: 2
//1 <= 数组长度 <= 50000

// 1. 暴力解法，先遍历后哈希表计数
// 2. 先排序，取中间数字
// 3.摩恩投票法
func majorityElement(nums []int) int {
	// 先获取数组长度
	length := len(nums)
	tempMap := make(map[int]int)
	for i := 0; i < length; i++ {
		tempMap[nums[i]]++
		if tempMap[nums[i]] > length/2 {
			return nums[i]
		}
	}
	return -1
}

// 我定义为同归于尽法，假设有N个阵营的人，其中有一个阵营超过了一半人（假设为A），如果我们采取一对一的同归于尽，比如说 不是A的人都会找一个A的同归于尽，由于A的人超过了一半，那么最后留下的人肯定是A的人
func majorityElementV2(nums []int) int {
	// 定义得票者
	candidate := 0
	// 定义得票数
	votes := 0
	for _, v := range nums {
		if votes == 0 {
			candidate = v
		}
		if candidate == v {
			votes++
		} else {
			votes--
		}
	}
	return candidate
}

func main() {
}
