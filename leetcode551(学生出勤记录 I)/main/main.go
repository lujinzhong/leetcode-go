package main

import "fmt"

//给你一个字符串 s 表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
//'A'：Absent，缺勤
//'L'：Late，迟到
//'P'：Present，到场
//如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//按 总出勤 计，学生缺勤（'A'）严格 少于两天。
//学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
//如果学生可以获得出勤奖励，返回 true ；否则，返回 false 。

func checkRecord(s string) bool {
	// 如果字符串中不存在两个及以上A, 并且不出现连续三次或者三次以上的L，则符合结果
	var count int
	var b bool
	// b 为 true 表示存在连续的三个L
	l := len(s)

	for i := 0; i < l; i++ {
		if s[i] == 'A' { // A的个数计算
			count++
			if count > 2 {
				return false
			}
		}
		if !b && i < l-2 && s[i] == 'L' && s[i+1] == 'L' && s[i+2] == 'L' {
			b = true
		}
	}
	return !b
}

func main() {
	fmt.Println(checkRecord("PPALALAL"))
}
