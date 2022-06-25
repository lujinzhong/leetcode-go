package main

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 此题和青蛙跳台阶问题相同。 f(n)=f(n-1)+f(n-2)

//1 <= n <= 45
func climbStairs(n int) int {
	// n=1 1
	// n=2 2
	// n=3 3
	// n=4 5
	if n <= 3 {
		return n
	}
	pre := 2
	cur := 3
	for i := 4; i <= n; i++ {
		temp := cur
		cur = cur + pre
		pre = temp
	}
	return cur
}

func main() {
}
