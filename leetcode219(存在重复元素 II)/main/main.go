package main

//给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
//1 <= nums.length <= 105
//-109 <= nums[i] <= 109
//0 <= k <= 105

func containsNearbyDuplicate(nums []int, k int) bool {
	set := map[int]struct{}{} // 用 struct{} 可以减少内存
	for i, num := range nums {
		if i > k { //超过 K 大小后，每次遍历删除最左边元素，小于K时，不做删除，先填充满窗口
			delete(set, nums[i-k-1]) // 删除指定 key 元素
		}
		if _, ok := set[num]; ok { // 存在重复元素，则返回 true
			return true
		}
		set[num] = struct{}{} // 将当前元素增加，相当于滑动窗口右移增加了新元素
	}
	return false
}

func main() {
}
