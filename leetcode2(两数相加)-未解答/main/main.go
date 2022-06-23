package main

import (
	"fmt"
	"math"
)

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := l1
	h2 := l2
	var n1, n2 int64
	i1, i2 := 0, 0
	for h1 != nil {
		if i1 == 0 {
			n1 = int64(h1.Val)
		} else {
			n1 += int64(math.Pow10(i1)) * int64(h1.Val)
		}
		i1++
		h1 = h1.Next
	}
	for h2 != nil {
		if i2 == 0 {
			n2 = int64(h2.Val)
		} else {
			n2 += int64(math.Pow10(i2)) * int64(h2.Val)
		}
		i2++
		h2 = h2.Next
	}
	total := n1 + n2
	// 换成数组
	var arr []int64
	for {
		if total < 10 {
			arr = append(arr, total)
			break
		}
		arr = append(arr, total%10)
		total = total / 10
	}
	maxIndex := len(arr) - 1
	// 开始搞新链表
	newHead := &ListNode{
		0,
		nil,
	}
	h3 := newHead
	for i := 0; i <= maxIndex; i++ {
		h3.Next = &ListNode{
			int(arr[i]),
			nil,
		}
		h3 = h3.Next
	}
	return newHead.Next

}

func main() {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				3,
				nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}
