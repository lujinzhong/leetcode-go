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

//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：2.1 MB, 在所有 Go 提交中击败了 99.95% 的用户
// 用双指针来实现，让两个指针相差 n 个空间，因此当后续指针走到最后一个时，前续指针刚好在要删除的前一个节点中
func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	if  head.Next == nil { // 1.只有一个节点
		return nil
	}

	first := head
	second := head
	// 先让 second 往后走 n-1 次
	for i := 0; i < n; i++ {
		second =  second.Next
	}
	// 如果 second == nil ，说明要删除的是第一个元素
	if second == nil {
		head = head.Next
		return head
	}

	// 然后两个指针再一起走, 直到 second 指针刚好在最后一个
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	// 此时 first 刚好是在要删除的前一个节点
	first.Next = first.Next.Next
	return head
}

func main() {

}
