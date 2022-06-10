package main

///设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//实现 MinStack 类:
//
//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//-23 <= val <= 231- 1
//pop、top 和 getMin 操作总是在 非空栈 上调用
//push,pop,top, and getMin最多被调用 3 * 104 次

type Node struct {
	Val  int
	Next *Node
}

type MinStack struct {
	topNode  *Node
	minValue int
}

//执行用时：16 ms, 在所有 Go 提交中击败了 72.43% 的用户
//内存消耗：7.7 MB , 在所有 Go 提交中击败了 97.25%的用户

func Constructor() MinStack {
	return MinStack{topNode: nil, minValue: 2147483647}
}

func (m *MinStack) Push(val int) {
	// 每次都判断一下当前值 和 minValue 的大小，实时更改
	if m.minValue > val {
		m.minValue = val
	}
	// 不为空的话，放到链表头部
	m.topNode = &Node{val, m.topNode}
}

func (m *MinStack) Pop() {
	// 出栈后，如果出的是最小值，需要变更最小值
	if m.minValue == m.topNode.Val { // 获取最小的值
		// 遍历下剩余的链表，获取最小值
		cur := m.topNode.Next
		if cur == nil { // 出的是最后一个元素, 恢复初始化
			m.minValue = 2147483647
		} else {
			minValue := cur.Val // 取第一个作为基准值
			for cur != nil {
				if cur.Val < minValue {
					minValue = cur.Val
				}
				cur = cur.Next
			}
			m.minValue = minValue
		}

	}
	m.topNode = m.topNode.Next

}

func (m *MinStack) Top() int {
	return m.topNode.Val
}

func (m *MinStack) GetMin() int {
	return m.minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
