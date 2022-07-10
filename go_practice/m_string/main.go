package main

import (
	"fmt"
)

func main() {
	TestString()
}

func TestString() {
	a := "我是MCLINK"
	for _, v := range a { // int32 == rune
		fmt.Println(v)
	}
	for i := 0; i < len(a); i++ { //unit8
		fmt.Println(a[i])
	}
}
