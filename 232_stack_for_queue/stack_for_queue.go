package lc0232

type Stack struct {
	element []int
}

func NewStack() *Stack {
	return &Stack{element: []int{}}
}

func (s *Stack) Push(n int) {
	s.element = append(s.element, n)
}
func (s *Stack) Pop() int {
	res := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return res
}

func (s *Stack) Len() int {
	return len(s.element)
}

func (s *Stack) IsEmpty() bool {
	return len(s.element) == 0
}

type MyQueue struct {
	front, back *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{front: NewStack(), back: NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.front.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.back.Len() == 0 {
		for this.front.Len() > 0 {
			this.back.Push(this.front.Pop())
		}
	}
	return this.back.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	res := this.Pop()
	this.back.Push(res)
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.front.Len() == 0 && this.back.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
