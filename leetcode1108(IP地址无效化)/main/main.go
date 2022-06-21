package main

//给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
//
//所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。

func defangIPaddr(address string) string {
	var str string
	for _, v := range address {
		if v == 46 {
			str += "[.]"
		} else {
			str += string(v)
		}
	}
	return str
}

func main() {
}
