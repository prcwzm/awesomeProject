package algorithm

type AQueue struct {
	Head AStack
	Tail AStack
}

type AStack struct {
	Stack []int
}

func CStack() *AStack {
	cs := new(AStack)
	cs.Stack = []int{}
	return cs
}

func (s *AStack) Pop() (val int) {
	if len(s.Stack) == 0 {
		return -1
	}
	val = s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	return val
}

func (s *AStack) Push(val int) {
	s.Stack = append(s.Stack, val)
}

func (a *AQueue) Len() int {
	return len(a.Tail.Stack) + len(a.Head.Stack)
}

func CQueue() (a *AQueue) {
	a = new(AQueue)
	a.Tail = *CStack()
	a.Head = *CStack()
	return a
}

func (a *AQueue) Pop() (head int) {
	for len(a.Head.Stack) > 0 {
		a.Tail.Push(a.Head.Pop())
	}
	return a.Tail.Pop()
}

func (a *AQueue) Push(head int) {
	for len(a.Tail.Stack) > 0 {
		a.Head.Push(a.Tail.Pop())
	}
	a.Head.Push(head)
}
