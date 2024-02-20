package structs

import (
	"cmp"
	"errors"
	"fmt"
	"log"
)

type list[T Ordered] struct {
	head   *node[T]
	length int
}

type node[T Ordered] struct {
	value T
	next  *node[T]
}

func NewLinkedList[T Ordered]() *list[T] {
	newHead := new(node[T])
	newList := new(list[T])
	newList.head = newHead
	newList.length = 0

	return newList
}

func (l *list[T]) getElementAtIndex(index int) *node[T] {
	if index+1 > l.length {
		return nil
	}

	nodeAtIndex := l.head

	for _ = range index {
		nodeAtIndex = nodeAtIndex.next
	}

	return nodeAtIndex
}

func (l *list[T]) swap(x int, y int) error {

	if x == y {
		return nil
	}

	if x < 0 || y < 0 || x >= l.length || y >= l.length {
		errorMessage := fmt.Sprintf("Function structs.swap(): Index %d and/or %d outside of list bounds. Expecting values in range [%d, %d]", x, y, 0, l.length-1)
		return errors.New(errorMessage)
	}

	if x > y {
		tmpIntSwap := x
		x = y
		y = tmpIntSwap
	}

	var nodeBehindX *node[T]
	nodeX := l.getElementAtIndex(x)
	nodeAfterX := nodeX.next

	nodeBehindY := l.getElementAtIndex(y - 1)
	nodeY := l.getElementAtIndex(y)
	nodeAfterY := nodeY.next

	if y-x == 1 {
		nodeY.next = nodeX
	} else {
		nodeY.next = nodeAfterX
		nodeBehindY.next = nodeX
	}

	nodeX.next = nodeAfterY

	if x == 0 {
		l.head = nodeY
	} else if x == 1 {
		l.head.next = nodeY
	} else {
		nodeBehindX = l.getElementAtIndex(x - 1)
		nodeBehindX.next = nodeY
	}

	return nil
}

func partition[T Ordered](list *list[T], low int, high int) (error, int) {
	if low > high {
		errorMessage := fmt.Sprintf("Function structs.partition(). Expected low index to be smaller than index high. Found low = %d, high = %d", low, high)
		return errors.New(errorMessage), -1
	}

	pivot := list.getElementAtIndex(high)
	positionSmallerThanPivot := low - 1

	for i := low; i < high; i++ {
		if node := list.getElementAtIndex(i); node.value < pivot.value {
			positionSmallerThanPivot++
			list.swap(positionSmallerThanPivot, i)
		}
	}
	list.swap(positionSmallerThanPivot+1, high)

	return nil, positionSmallerThanPivot + 1
}

func quicksort[T Ordered](list *list[T], low int, high int) {
	if low < high {
		err, pivot := partition[T](list, low, high)
		if err != nil {
			log.Fatal(err)
		}

		quicksort[T](list, low, pivot-1)
		quicksort[T](list, pivot+1, high)
	}
}

func (l *list[T]) Insert(newElement T) {
	if l.length == 0 {
		l.head.value = newElement
		l.length += 1
		return
	}

	newNode := new(node[T])
	newNode.value = newElement

	if l.length == 1 {
		l.head.next = newNode
	} else {
		tail := l.getElementAtIndex(l.length - 1)
		tail.next = newNode
	}

	l.length += 1
}

func (l *list[T]) DeleteBeginning() {

	switch lenght := l.length; lenght {
	case 0:
		return
	case 1:
	default:
		l.head = l.head.next
	}

	l.length -= 1

}

func (l *list[T]) DeleteEnd() {

	switch lenght := l.length; lenght {
	case 0:
		return
	case 1:
	default:
		nextToLast := l.getElementAtIndex(l.length - 2)
		nextToLast.next = nil
	}

	l.length -= 1
}

func (l *list[T]) DeleteAtIndex(index int) {
	if l.length-1 < index {
		return
	}

	switch index {
	case 0:
		l.DeleteBeginning()
	case l.length - 1:
		l.DeleteEnd()
	default:
		nodeBeforIndex := l.getElementAtIndex(index - 1)
		nodeAfterIndex := l.getElementAtIndex(index + 1)

		nodeBeforIndex.next = nodeAfterIndex
	}
}

func (l *list[T]) Sort() {
	if l.length < 2 {
		return
	}

	currentNode := l.head
	sorted := true

	for _ = range l.length - 1 {
		if sorted && cmp.Less[T](currentNode.next.value, currentNode.value) {
			sorted = false
			break
		}
		currentNode = currentNode.next
	}

	if sorted {
		return
	}

	quicksort[T](l, 0, l.length-1)
}

func (l *list[T]) Print() {

	if l.length == 0 {
		fmt.Println("List is empty.")
		return
	}

	currentNode := l.head

	for {
		fmt.Printf("%v ", currentNode.value)

		if currentNode.next == nil {
			break
		} else {
			fmt.Print("-> ")
			currentNode = currentNode.next
		}
	}

	fmt.Println()
}
