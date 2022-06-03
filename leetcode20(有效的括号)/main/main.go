package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，'] 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合

//1 <= s.length <= 104
//s 仅由括号 '()[]{}' 组成

// 用栈
//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：2.1 MB, 在所有 Go 提交中击败了8.95%
func isValid(s string) bool {

	// 如果是奇数个，直接不通过
	if len(s) % 2 == 1 {
		return false
	}
	// 先搞个 map 映射下关系
	tempMap := map[string]string{")" : "(", "]" : "[", "}" :  "{"}
	// 初始化栈
	stack := NewStack()
	for _, v := range s {
		if stack.topNode == nil { // 第一个元素进栈
			stack.Push(string(v))
			continue
		}
		temp := tempMap[string(v)]
		if temp == stack.topNode.data {
			stack.Pop()
		} else {
			stack.Push(string(v))
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}

// Node 定义节点类型
type Node struct {
	data interface{}
	next *Node
}

// Stack 定义栈类型
type Stack struct {
	topNode *Node
}

// NewStack 实现了新建一个栈的功能
func NewStack() *Stack {
	return &Stack{nil}
}

// Push 实现了push元素的功能
func (s *Stack) Push(value interface{}) {
	s.topNode = &Node{value, s.topNode}
}

// IsEmpty 实现了判断栈是否为空的功能
func (s *Stack) IsEmpty() bool {
	return s.topNode == nil
}

// Pop 实现了pop的功能
func (s *Stack) Pop() interface{} {
	if s.topNode == nil {
		return nil
	}
	res := s.topNode.data
	s.topNode = s.topNode.next
	return res

}



func main() {
	s := "()"
	fmt.Println(isValid(s))
}
