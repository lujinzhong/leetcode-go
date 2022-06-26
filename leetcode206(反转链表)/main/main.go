package main

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
// Related Topics 递归 链表 👍 2546 👎 0

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

// 反转链表，迭代版本， O(n)
func reverseList(head *ListNode) *ListNode {
	// nil(pre) 1(head)->2->3
	// nil <-1(pre) 2(head)->3
	// nil<-1<-2(pre) 3(head)->nil
	// nil<-1<-2<-3(pre) nil(head)
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		// 保留被断开的节点
		next := cur.Next
		// 开始移情别恋
		cur.Next = pre
		// 重新开始
		pre = cur
		cur = next
	}
	return pre
}

// 递归版本
// nil(pre) head(1) x
func reverseListV2(head *ListNode) *ListNode {
	var f func(*ListNode, *ListNode) *ListNode
	f = func(pre, head *ListNode) *ListNode {
		if head == nil {
			return pre
		}
		// 保留要抛弃的节点
		next := head.Next
		head.Next = pre
		return f(head, next)
	}

	return f(nil, head)
}

func main() {

}
