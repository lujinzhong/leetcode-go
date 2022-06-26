package main

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//提示:
//
//1 <= nums.length <= 104
//-231 <= nums[i] <= 231 - 1

// 双指针法
func moveZeroes(nums []int) {
	slow, fast, n := 0, 1, len(nums)-1
	// 每次都一起走，直到 slow 找到零元素
	for fast <= n {
		if nums[slow] != 0 {
			slow++
			fast++
		} else {
			// fast 找到非零元素
			for fast <= n {
				if nums[fast] == 0 {
					fast++
				} else {
					// 替换元素
					nums[slow], nums[fast] = nums[fast], nums[slow]
					slow++
					fast++
					break
				}
			}
		}
	}
}

func main() {

}
