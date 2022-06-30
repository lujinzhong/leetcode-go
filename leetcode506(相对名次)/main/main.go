package main

//给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。所有得分都 互不相同 。
//
//运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，依此类推。运动员的名次决定了他们的获奖情况：
//
//名次第 1 的运动员获金牌 "Gold Medal" 。
//名次第 2 的运动员获银牌 "Silver Medal" 。
//名次第 3 的运动员获铜牌 "Bronze Medal" 。
//从名次第 4 到第 n 的运动员，只能获得他们的名次编号（即，名次第 x 的运动员获得编号 "x"）。
//使用长度为 n 的数组 answer 返回获奖，其中 answer[i] 是第 i 位运动员的获奖情况。
//
//n == score.length
//1 <= n <= 104
//0 <= score[i] <= 106
//score 中的所有值 互不相同

import (
	"fmt"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	// 然后排序， 递减
	sortScore := make([]int, len(score))
	copy(sortScore, score)
	sortScore = bubblingSort(sortScore)
	temp := make(map[int]string)
	// 记住位置先, 每个分数对应的下标位置或者奖牌
	for k, v := range sortScore {
		if k == 0 {
			temp[v] = "Gold Medal"
		} else if k == 1 {
			temp[v] = "Silver Medal"
		} else if k == 2 {
			temp[v] = "Bronze Medal"
		} else {
			temp[v] = strconv.Itoa(k + 1)
		}
	}
	// 再遍历原数组，做下替换
	res := make([]string, len(score))
	for k, v := range score {
		res[k] = temp[v]
	}
	return res
}

func bubblingSort(arr []int) []int {
	count := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < count-1; j++ {
			if arr[j] < arr[j+1] { // 判断大小决定是否替换
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
		count-- //这是最关键的，每次处理后，最大的元素都会在最右边，因此后续要处理的数组长度均减一
	}
	return arr
}

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
