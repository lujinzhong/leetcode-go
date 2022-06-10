package main

//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
//
// 如果有两个中间结点，则返回第二个中间结点。
//
//
//
// 示例 1：
//
//
//输入：[1,2,3,4,5]
//输出：此列表中的结点 3 (序列化形式：[3,4,5])
//返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
//注意，我们返回了一个 ListNode 类型的对象 ans，这样：
//ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next =
//NULL.
//
//
// 示例 2：
//
//
//输入：[1,2,3,4,5,6]
//输出：此列表中的结点 4 (序列化形式：[4,5,6])
//由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
//
//
//
//
// 提示：
//
//
// 给定链表的结点数介于 1 和 100 之间。
//
// Related Topics 链表 双指针 👍 579 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
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

// 双指针，遍历一次获得长度，再遍历一半长度
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil { // 只有两个节点
		return head.Next
	}
	// 有三个节点
	cur := head
	length := 1
	// 遍历计数链表的长度
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	time := length / 2
	// 前置指针回归 time 次
	for i := 0; i < time; i++ {
		head = head.Next
	}
	return head
}

// 快慢指针
func reverseListV2(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		// 1,2,3
		// 1,2,3,4
		// 1,2,3,4,5
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
}
