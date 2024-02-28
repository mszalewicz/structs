package structs

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {

	functionName := "Insert"

	tests := []struct {
		expected []int
	}{
		{[]int{4, 9, 7, 1, 4}},
		{[]int{8, 11, 1, 0, 3, 4}},
		{[]int{2, 2, 4, 0, 711, 1}},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.expected {
				ll.Insert(element)
			}

			currentNode := ll.head

			for i := range ll.length - 0 {
				result := currentNode.value

				if result != test.expected[i] {
					t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected[i], result)
				}

				currentNode = currentNode.next
			}

		})
	}

}

func TestGetElementAtIndex(t *testing.T) {

	functionName := "getElementAtIndex"

	tests := []struct {
		input    []int
		index    int
		expected interface{}
	}{
		{[]int{5, 9, 7, 1, 4}, 2, 7},
		{[]int{0, 3, 2, 4}, 0, 0},
		{[]int{0, 3, 2, 4}, 7, nil},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.input {
				ll.Insert(element)
			}

			var result interface{}

			if ll.getElementAtIndex(test.index) != nil {
				result = ll.getElementAtIndex(test.index).value
			} else {
				result = nil
			}

			if result != test.expected {
				t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected, result)
			}
		})
	}
}

func TestSwap(t *testing.T) {

	functionName := "swap"

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 9, 7, 1, 4}, []int{7, 9, 5, 1, 4}},
		{[]int{0, 3, 4}, []int{4, 3, 0}},
		{[]int{0, 3, 2, 4}, []int{2, 3, 0, 4}},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.input {
				ll.Insert(element)
			}

			ll.swap(0, 2)
			currentNode := ll.head

			for i := range ll.length - 1 {
				result := currentNode.value

				if result != test.expected[i] {
					t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected[i], result)
				}

				currentNode = currentNode.next
			}

		})
	}

}

func TestQuicksort(t *testing.T) {

	functionName := "quicksort"

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 9, 7, 1, 4}, []int{1, 4, 5, 7, 9}},
		{[]int{9, 11, 1, 0, 3, 4}, []int{0, 1, 3, 4, 9, 11}},
		{[]int{3, 2, 4, 0, 711, 1}, []int{0, 1, 2, 3, 4, 711}},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.input {
				ll.Insert(element)
			}

			ll.Sort()
			currentNode := ll.head

			for i := range ll.length - 1 {
				result := currentNode.value

				if result != test.expected[i] {
					t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected[i], result)
				}

				currentNode = currentNode.next
			}

		})
	}

}

func TestDeleteBeginning(t *testing.T) {

	functionName := "DeleteBeginning"

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 9, 7, 1, 4}, []int{9, 7, 1, 4}},
		{[]int{9, 11, 1, 0, 3, 4}, []int{11, 1, 0, 3, 4}},
		{[]int{3, 2}, []int{2}},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.input {
				ll.Insert(element)
			}

			ll.DeleteBeginning()
			currentNode := ll.head

			for i := range ll.length - 1 {
				result := currentNode.value

				if result != test.expected[i] {
					t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected[i], result)
				}

				currentNode = currentNode.next
			}

		})
	}
}

func TestDeleteEnd(t *testing.T) {

	functionName := "DeleteEnd"

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 9, 7, 1, 4}, []int{5, 9, 7, 1}},
		{[]int{9, 11, 1, 0, 3, 4}, []int{9, 11, 1, 0, 3}},
		{[]int{3, 2}, []int{3}},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.input {
				ll.Insert(element)
			}

			ll.DeleteEnd()
			currentNode := ll.head

			for i := range ll.length - 1 {
				result := currentNode.value

				if result != test.expected[i] {
					t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected[i], result)
				}

				currentNode = currentNode.next
			}

		})
	}
}

func TestDeleteAtIndex(t *testing.T) {

	functionName := "AtIndex"

	tests := []struct {
		index    int
		input    []int
		expected []int
	}{
		{0, []int{5, 9, 7, 1, 4}, []int{9, 7, 1, 4}},
		{4, []int{9, 11, 1, 0, 3, 4}, []int{9, 11, 1, 0, 4}},
		{1, []int{3, 2}, []int{3}},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			ll := NewLinkedList[int]()
			for _, element := range test.input {
				ll.Insert(element)
			}

			ll.DeleteAtIndex(test.index)
			currentNode := ll.head

			for i := range ll.length - 1 {
				result := currentNode.value

				if result != test.expected[i] {
					t.Fatalf("%s fail - expected: %v | result: %v", functionName, test.expected[i], result)
				}

				currentNode = currentNode.next
			}

		})
	}
}
