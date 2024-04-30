//@Author: wulinlin
//@Description: 两个栈实现一个队列  同leetcode https://leetcode.cn/problems/implement-queue-using-stacks/description/
//@File:  ans_v1
//@Version: 1.0.0
//@Date: 2024/04/30 00:57

package _2_two_stack_make_queue

type MyQueue struct {
	addStack []int
	delStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		addStack: make([]int, 0),
		delStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.addStack = append(this.addStack, x)
}

// Pop
//
//	@Description: 移除队头元素并返回
//	@receiver this
//	@return int
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if len(this.delStack) == 0 {
		this.moveToDelStack()
	}
	// 出栈
	tmp := this.delStack[len(this.delStack)-1]
	this.delStack = this.delStack[:len(this.delStack)-1]
	return tmp
}

func (this *MyQueue) moveToDelStack() {
	for len(this.addStack) > 0 {
		// 弹出一个元素
		tmp := this.addStack[len(this.addStack)-1]
		this.addStack = this.addStack[:len(this.addStack)-1]
		this.delStack = append(this.delStack, tmp) // 将这个元素加入到删除栈中
	}
}

// Peek
//
//	@Description: 返回队列开头的元素 不会移除这个元素，仅仅获取并且打印
//	@receiver this
//	@return int
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if len(this.delStack) == 0 {
		this.moveToDelStack()
	}
	tmp := this.delStack[len(this.delStack)-1]
	return tmp
}

// Empty
//
//	@Description: 如果队列为空，返回 true ；否则，返回 false
//	@receiver this
//	@return bool
func (this *MyQueue) Empty() bool {
	if len(this.delStack) == 0 && len(this.addStack) == 0 {
		return true
	}
	return false
}
