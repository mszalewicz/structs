package structs

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {

	functionName := "Enqueue"

	tests := []struct {
		expected []int
	}{
		{[]int{1, 8, 9}},
		{[]int{489, 90, 2, 1}},
		{nil},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {

			q := NewQueue[int]()

			for _, element := range test.expected {
				q.Enqueue(element)
			}

			for i, element := range *q.values {
				if element != test.expected[i] {
					t.Fatalf("%s fail - expected: %d | result: %d", functionName, test.expected, *q.values)
				}
			}
		})
	}
}

func TestDequeue(t *testing.T) {

	functionName := "Dequeue"

	tests := []struct {
		expected []int
	}{
		{[]int{1, 8, 9}},
		{[]int{489, 90, 2, 1}},
		{nil},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {

			q := NewQueue[int]()

			for _, element := range test.expected {
				q.Enqueue(element)
			}

			for i, _ := range *q.values {

				firstElement := test.expected[i:][0]
				dequeuedElement := q.Dequeue()

				if firstElement != dequeuedElement || len(test.expected)-i-1 != len(*q.values) {
					t.Fatalf("%s fail - expected: %d | result: %d", functionName, test.expected[i+1:], *q.values)
				}
			}
		})
	}

}
