package main

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

//输入：head = [1,3,2]
//输出：[2,3,1]
//
//0 <= 链表长度 <= 10000

type ListNode struct {
	Val  int
	Next *ListNode
}

//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：3 MB, 在所有 Go 提交中击败了 90.12% 的用户

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	if head.Next == nil {
		return []int{head.Val}
	}

	var arr []int
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	// 根据获取的数组，进行原地反序 [1,3,2,5,8] => [8,5,2,3,1]
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = temp
	}
	return arr
}
func main() {
}
