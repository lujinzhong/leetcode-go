package main

import "fmt"

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{
		inStack:  nil,
		outStack: nil,
	}
}

// 入队
func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

// 出队，先判断是否为空，是的话做一次搬运动作
func (this *CQueue) DeleteHead() int {
	outLength := len(this.outStack)
	if outLength == 0 { // 走搬运动作，逆反向搬运
		inLength := len(this.inStack)
		if inLength == 0 { // 此时输入输出栈都是空的
			return -1
		}
		for i := inLength - 1; i >= 0; i-- { // 这里用的是栈，那么要遵守每次出栈都是最后一个元素，不然就没有价值了，所以这里要逆序
			this.outStack = append(this.outStack, this.inStack[i])
		}
		// 搬运完清空输入栈
		this.inStack = nil
	}
	// 然后再判断一下长度
	outLength = len(this.outStack)
	if outLength == 0 {
		return -1
	}
	temp := this.outStack[outLength-1] // 取最后一个元素
	if len(this.outStack) >= 2 {
		this.outStack = this.outStack[0 : outLength-1]
	} else {
		this.outStack = nil
	}

	return temp
}

func main() {
	obj := Constructor()
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}
