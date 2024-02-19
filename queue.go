package structs

import "cmp"

type queue[T cmp.Ordered] struct {
	values *[]T
}

// Initializes queue for given type T.
func NewQueue[T cmp.Ordered]() *queue[T] {
	newQueue := new(queue[T])
	newSlice := make([]T, 0, resizeSize)
	newQueue.values = &newSlice
	return newQueue
}

// Adds new element T to queue.
func (q *queue[T]) Enqueue(newElement T) {
	if len(*q.values) == cap(*q.values) {
		currentMaxSize := cap(*q.values)
		newSlice := make([]T, 0, currentMaxSize+resizeSize)
		copy(newSlice, *q.values)
		q.values = &newSlice
	}

	*q.values = append(*q.values, newElement)
}

// Reads first element T and deletes it from queue.
func (q *queue[T]) Dequeue() interface{} {

	if len(*q.values) == 0 {
		return nil
	}

	firstInQueue := (*q.values)[0]
	*q.values = (*q.values)[1:]

	return firstInQueue
}

// Returns value of current queue front.
func (q *queue[T]) Peek() interface{} {
	if len(*q.values) == 0 {
		return nil
	}
	return (*q.values)[0]
}

// Return current length of queue.
func (q *queue[T]) Len() int {
	return len(*q.values)
}
