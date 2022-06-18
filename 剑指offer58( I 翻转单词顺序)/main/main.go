package main

import "fmt"

//输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
//无空格字符构成一个单词。
//输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个

func reverseWords(s string) string {
	var stack []string
	// 每次遇到非空格字符开始处理
	var tempStr = ""
	for _, v := range s {
		if tempStr != "" {
			if string(v) != " " {
				tempStr += string(v)
			} else {
				stack = append(stack, tempStr)
				tempStr = ""
			}
		} else {
			if string(v) != " " {
				tempStr += string(v)
			}
		}
	}
	if tempStr != "" {
		stack = append(stack, tempStr)
	}
	length := len(stack)
	newStr := ""
	for i := length - 1; i >= 0; i-- {
		if i != length-1 {
			newStr += " " + stack[i]
		} else {
			newStr += stack[i]
		}
	}
	return newStr
}
func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))

}
