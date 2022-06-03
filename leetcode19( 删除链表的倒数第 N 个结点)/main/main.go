package main

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz
type ListNode struct {
	Val  int
	Next *ListNode
}


//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//输入：head = [1], n = 1
//输出：[]
//输入：head = [1,2], n = 1
//输出：[1]

//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：2.4 MB, 在所有 Go 提交中击败了 5.98% 的用户

// 暴力破解，先遍历一遍，获取长度，然后再处理第N个节点，只要找到倒数第N+1 个节点，就可以处理了
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 考虑几种临界情况
	// 1. 只有一个节点的时候
	// 2. 要干掉第一个节点的时候
	// 3. 要干掉最后一个节点的时候


	if  head.Next == nil { // 1.只有一个节点
		return nil
	}
	// 获取链表长度
	i := 0
	tempMap := make(map[int]interface{})
	cur := head
	for cur != nil {
		i = i+ 1
		tempMap[i] = cur
		cur = cur.Next
	}

	// 2.如果N=1 特殊处理，直接干掉最后一个
	if n == 1{
		// 获取倒数第二个节点
		cur2 := tempMap[i-1].(*ListNode)
		cur2.Next = nil
		return head
	}


	// 算出倒数第N+1节点的位置
	temp := i-n
	// 3.说明要干掉第一个
	if temp == 0 {
		head = head.Next
		return head
	}
	cur2 :=  tempMap[temp].(*ListNode)
	cur2.Next = cur2.Next.Next

	return head
}

func main() {

}
