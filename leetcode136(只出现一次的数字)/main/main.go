package main

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

// 暴力解法，用map 来计数
func singleNumber(nums []int) int {
	tempMap := make(map[int]int)
	for _, v := range nums {
		tempMap[v] = tempMap[v] + 1
	}
	for k, v := range tempMap {
		if v != 2 {
			return k
		}
	}
	return -1
}

// 使用异或运算符
// 0^0=0 0^1=1 1^1=0
// a^b^a = (a^a)^b = 0^b = b
func singleNumberV2(nums []int) int {
	target := 0
	for _, v := range nums {
		target ^= v
	}
	return target
}

func main() {
}
