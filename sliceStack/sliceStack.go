package sliceStack

type SliceStack[T comparable] struct {
	stack []T
}

func (s *SliceStack[T]) Len() int         { return len(s.stack) }
func (s *SliceStack[T]) Append(element T) { s.stack = append(s.stack, element) }
func (s *SliceStack[T]) PopAndGet() (top T) {
	if s.Len() == 0 {
		return nil
	}
	top = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}
func (s *SliceStack[T]) Pop() {
	if s.Len() == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *SliceStack[T]) GetTop() (top T) {
	if s.Len() == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

func (s *SliceStack[T]) RemoveElement(element T) {
	i := 0
	for _, v := range s.stack {
		if v != element {
			s.stack[i] = v
			i++
		}
	}
}

func (s *SliceStack[T]) RemoveI(i int) {
	copy(s.stack[i:], s.stack[i+1:])
}

func (s *SliceStack[T]) GetList() []T {
	return s.stack
}

func (s *SliceStack[T]) Push(element T) {

}
