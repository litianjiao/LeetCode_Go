package lc0155

type MinStack struct {
	unit []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{unit: []int{}}

}

func (this *MinStack) Push(x int) {
	this.unit = append(this.unit, x)
}

func (this *MinStack) Pop() {
	//	res := this.unit[len(this.unit)-1]
	this.unit = this.unit[:len(this.unit)-1]
}

func (this *MinStack) Top() int {
	return this.unit[0]
}

func (this *MinStack) GetMin() int {

	for i, _ := range this.unit {
		if this.unit[i] > this.unit[i+1] {
			this.Pop()
		}
	}
	return this.unit[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
