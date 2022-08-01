package main

import "fmt"

//实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
//-100.0 < x < 100.0
//-231 <= n <= 231-1
//-104 <= xn <= 104

// 递归做法
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 { //  兼容负次幂，取倒数
		return 1 / mySubPow(x, -n)
	} else {
		return mySubPow(x, n)
	}
}

// 非递归写法
func myPow2(x float64, n int) float64 {

	if n < 0 { // 小于0，取倒数，n 转正
		x = 1 / x
		n = -n
	} else if n == 0 {
		return x
	}

	var res float64 = 1
	for n != 0 {
		if n&1 == 1 { // 是奇的话，先乘x
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func mySubPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	temp := mySubPow(x, n>>1)
	if n&1 == 0 { // 偶数
		return temp * temp
	} else {
		return temp * temp * x
	}
}

func main() {
	//fmt.Println(myPow(2, -3))
	//fmt.Println(myPow2(2, -3))
	fmt.Println(myPow2(2, 10))

}
