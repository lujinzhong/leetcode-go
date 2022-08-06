package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//提示：
//
//1 <= n <= 8

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	if n == 0 {
		return res
	}
	getParenthesis("", n, n)
	return res
}

func getParenthesis(str string, left, right int) {
	if left == 0 && right == 0 {
		res = append(res, str)
		return
	}
	if left == right { // 左右数目相等，使用左括号
		getParenthesis(str+"(", left-1, right)
	} else if left < right {
		if left > 0 {
			getParenthesis(str+"(", left-1, right)
		}
		getParenthesis(str+")", left, right-1)
	}
}

func main() {

}
