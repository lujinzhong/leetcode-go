package main

import (
	"fmt"
	"strings"
)

//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
// Related Topics 双指针 字符串 👍 530 👎 0

// 注：这道题要注意非字符的元素 ASCII 字符中 A-Z 65 到 90，a-z 97-122
// 1. 剔除无效元素，翻转字符串和原有字符串比较
func a1(s string) bool {
	if s == "" {
		return true
	}
	var sArr []rune
	for _, v := range s {
		if isLetterOrNumber(v) {
			sArr = append(sArr, v)
		}
	}
	if len(sArr) == 1 {
		return true
	}

	str := strings.ToLower(string(sArr))
	r := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		r[i] = str[len(str)-i-1]
	}
	newStr := string(r)
	if str == newStr {
		return true
	}
	return false
}

func isLetterOrNumber(s rune) bool {
	// a ~ z , A ~ Z, 0-9
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (int(s) >= 48 && int(s) <= 57) { // 字符串 0和9分别对应48和57
		return true
	}
	return false
}

// 这个占用内存更小一丢丢，速度是一样的
func isLetterOrNumberV2(s rune) bool {
	// a ~ z , A ~ Z, 0-9
	if (s >= 97 && s <= 122) || (s >= 65 && s <= 90) || (s >= 48 && s <= 57) { // 字符串 0和9分别对应48和57
		return true
	}
	return false
}

// 双指针
func a2(s string) bool {
	if s == "" {
		return true
	}

	length := len(s)
	left := 0           // 左指针
	right := length - 1 // 右指针

	for left <= right {
		// 优先固定左指针
		if !isLetterOrNumber(rune(s[left])) { //左边第一个满足条件的
			left++
			continue
		}

		if !isLetterOrNumber(rune(s[right])) { // 过滤不满足的
			right--
		} else { // 满足，判断是否相等
			if strings.ToLower(string(s[left])) == strings.ToLower(string(s[right])) { // 这里性能比较差
				left++ // 匹配上了，左指针+1， 右指针减一
				right--
				continue // 继续匹配
			} else {
				return false // 没匹配上，直接GG
			}
		}
	}
	return true
}

// 双指针优化, 不用强转，直接数字比较，利用 ASCII 码的特性
func a3(s string) bool {
	if s == "" {
		return true
	}
	length := len(s)
	if length == 1 {
		return true
	}
	left := 0           // 左指针
	right := length - 1 // 右指针

	for left <= right {
		// 优先固定左指针
		if !isLetterOrNumberV2(rune(s[left])) { //左边第一个满足条件的
			left++
			continue
		}

		if !isLetterOrNumberV2(rune(s[right])) { // 过滤不满足的
			right--
		} else { // 满足，判断是否相等
			if s[left] == s[right] { // 如果是同样大小写字符，则可以直接比较
				left++ // 匹配上了，左指针+1， 右指针减一
				right--
				continue // 继续匹配
			} else if (s[right] >= 65 && s[right] <= 90) && // 右指针是 A-Z
				(s[left] >= 97 && s[left] <= 122) && // 左指针是 a-z
				(s[left]-s[right] == 32) { // 相减为 32 ，说明是匹配的
				left++ // 匹配上了，左指针+1， 右指针减一
				right--
				continue // 继续匹配
			} else if (s[left] >= 65 && s[left] <= 90) && // 左指针是 A-Z
				(s[right] >= 97 && s[right] <= 122) && // 右指针是 a-z
				(s[right]-s[left] == 32) { // 相减为 32 ，说明是匹配的
				left++ // 匹配上了，左指针+1， 右指针减一
				right--
				continue // 继续匹配
			} else {
				return false // 没匹配上，直接GG
			}
		}
	}
	return true
}

func main() {
	fmt.Print(a1("0P"))
	fmt.Print(a2("0P"))
	fmt.Print(a3("0P"))
}
