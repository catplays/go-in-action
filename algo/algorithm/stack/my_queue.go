package stack

type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack: make([]int, 0),
		popStack:  make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack,x)
}

func (this *MyQueue) Pop() int {
	val := this.Peek()
	this.popStack = this.popStack[:len(this.popStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		length :=len(this.pushStack)
		for i:= length-1;i>=0 ;i-- {
			this.popStack = append(this.popStack,this.pushStack[i])
		}
		this.pushStack = this.pushStack[length:length]
	}
	return this.popStack[len(this.popStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.popStack) ==0 && len(this.pushStack) == 0
}
