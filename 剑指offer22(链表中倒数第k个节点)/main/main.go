package main

//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.

type ListNode struct {
	Val  int
	Next *ListNode
}

//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：2.3 MB, 在所有 Go 提交中击败了 12.14% 的用户
// 1.直接遍历获取节点数后，再遍历到倒数第K个节点
// 2.直接遍历获取节点数并且放到哈希表中，然后可以用O(1)的方式找到倒数第N个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 先获取链表长度，然后通过哈希表来快速找到某个节点
	cur := head
	nums := 0
	hashMap := make(map[int]*ListNode)
	for cur != nil {
		nums = nums + 1
		hashMap[nums] = cur
		cur = cur.Next
	}
	return hashMap[nums+1-k] // 假设有5个节点，倒数第三个，就是 5-3+1 = 3
}

func main() {
}
