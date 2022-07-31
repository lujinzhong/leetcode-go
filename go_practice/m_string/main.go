package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	TestStringReader()
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

func TestByteBuild() {
	var b bytes.Buffer
	b.Grow(20)
	b.WriteString("a")
	fmt.Println(b.String())
}

func TestStringBuild() {
	var s strings.Builder
	s.Grow(200)
	for i := 0; i < 10; i++ {
		s.WriteString(fmt.Sprintf("num %d", i))
	}
	fmt.Println(s.String())
}

func TestStringReader() {
	s := strings.NewReader("abcdjfdksjalkfjd")
	read, err := s.Read([]byte{'c', 'd', 'f'})
	if err != nil {
		return
	}
	fmt.Println(read)
}
