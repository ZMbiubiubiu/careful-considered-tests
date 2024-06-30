// 用两个栈实现一个队列

package main

type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: nil,
		s2: nil,
	}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	var result int
	if len(this.s2) != 0 {
		result = this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
		return result
	}
	// 将s1腾挪
	for i := len(this.s1) - 1; i >= 0; i-- {
		this.s2 = append(this.s2, this.s1[i])
	}
	this.s1 = nil

	if len(this.s2) != 0 {
		result = this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
		return result
	}
	return result
}

func (this *MyQueue) Peek() int {
	var result int
	if len(this.s2) != 0 {
		result = this.s2[len(this.s2)-1]
		return result
	}
	// 将s1腾挪
	for i := len(this.s1) - 1; i >= 0; i-- {
		this.s2 = append(this.s2, this.s1[i])
	}
	this.s1 = nil

	if len(this.s2) != 0 {
		result = this.s2[len(this.s2)-1]
		return result
	}
	return result
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
