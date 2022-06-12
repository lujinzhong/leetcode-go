package main

import "fmt"

//
//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//F(0) = 0,   F(1) = 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1

//递归法，LeetCode 无法通过，时间太长了
func fib(n int) int {
	res := fibSub(n)
	return res % 1000000007
}

func fibSub(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	return fibSub(n-1) + fibSub(n-2)
}

//执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：1.8 MB, 在所有 Go 提交中击败了 73.26% 的用户
func fibV2(n int) int {
	if n < 2 {
		return n
	}
	// 以 f(3) 举例， f(3) = f(2) + f(1) = f(1)+f(0)+f(1) = 2
	// 滑动数组 [0[p],0[q],1[z]] => [0[p],1[q],1[z]] => [1,1,2]
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % (1e9 + 7)
	}
	return r

}

func main() {
	fmt.Println(fibV2(95))
}
