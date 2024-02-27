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
