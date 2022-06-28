package main

//波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
//F(0) = 0，F(1) = 1
//F(n) = F(n - 1) + F(n - 2)，其中 n > 1
//给定 n ，请计算 F(n) 。

// 递归法
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 迭代法
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	// n=0 sum=0, n=1 sum=1 , n=2 sum=1, n=3 sum=2, n=4 sum=3

	// 从 n=4
	cur := 1
	sum := 2
	for n-3 > 0 {
		temp := sum
		sum = sum + cur
		cur = temp
		n--
	}
	return sum
}

func main() {
}
