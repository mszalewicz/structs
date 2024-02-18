package structs

type stack[T Ordered] struct {
	values   *[]T
	topIndex int
}

// Initializes stack for given type T
func NewStack[T Ordered]() *stack[T] {
	newStack := new(stack[T])
	newSlice := make([]T, 0, resizeSize)
	newStack.values = &newSlice
	newStack.topIndex = -1
	return newStack
}

// Push new element of type T into stack.
// Automatically resizes stack as needed.
func (s *stack[T]) Push(newElement T) {
	if len(*s.values) == cap(*s.values) {
		currentMaxSize := cap(*s.values)
		newSlice := make([]T, 0, currentMaxSize+resizeSize)
		copy(newSlice, *s.values)
		s.values = &newSlice
	}

	*s.values = append(*s.values, newElement)
	s.topIndex += 1
}

// Returns top value and deletes it from stack.
func (s *stack[T]) Pop() interface{} {

	if len(*s.values) == 0 {
		return nil
	}

	topElement := (*s.values)[len(*s.values)-1]
	*s.values = (*s.values)[:len(*s.values)-1]
	s.topIndex -= 1

	return topElement
}

// Returns value of current stack top.
func (s *stack[T]) Peek() interface{} {
	if s.topIndex == -1 {
		return nil
	}

	return (*s.values)[s.topIndex]
}

// Return current height of stack.
func (s *stack[T]) Len() int {
	return s.topIndex + 1
}
