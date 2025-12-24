package main

type MinStack struct {
	minStk   []int
	storeStk []int
}

func Constructor() MinStack {
	var stk MinStack
	return stk
}

func (this *MinStack) Push(val int) {
	if len(this.minStk) == 0 || this.minStk[len(this.minStk)-1] >= val {
		this.minStk = append(this.minStk, val)
	}
	this.storeStk = append(this.storeStk, val)
}

func (this *MinStack) Pop() {
	if len(this.storeStk) != 0 {
		ls, lm := len(this.storeStk), len(this.minStk)
		val := this.storeStk[ls-1]
		if val == this.minStk[lm-1] {
			this.minStk = this.minStk[:lm-1]
		}
		this.storeStk = this.storeStk[:ls-1]
	}
}

func (this *MinStack) Top() int {
	return this.storeStk[len(this.storeStk)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}
