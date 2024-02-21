package structs

import (
	"testing"
)

type Test struct {
	expected []interface{}
}

func TestPush(t *testing.T) {

	functionName := "Push"

	// Test 1
	s1 := NewStack[int]()
	expected1 := []int{1, 2, 3}

	for _, element := range expected1 {
		s1.Push(element)
	}

	for index, expected := range expected1 {
		if (*(*s1).values)[index] != expected {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, expected1, *s1.values)
		}
	}

	// Test 2
	s2 := NewStack[float64]()
	expected2 := []float64{.3, 78.9, 30.}

	for _, element := range expected2 {
		s2.Push(element)
	}

	for index, expected := range expected2 {
		if (*(*s2).values)[index] != expected {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, expected2, *s2.values)
		}
	}

	// Test 3
	s3 := NewStack[string]()
	expected3 := []string{"z", "c", "t"}

	for _, element := range expected3 {
		s3.Push(element)
	}

	for index, expected := range expected3 {
		if (*(*s3).values)[index] != expected {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, expected3, *s3.values)
		}
	}

	// Test 4
	s4 := NewStack[uint]()
	expected4 := []uint{451, 562, 3}

	for _, element := range expected4 {
		s4.Push(element)
	}

	for index, expected := range expected4 {
		if (*(*s4).values)[index] != expected {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, expected4, *s4.values)
		}
	}
}

func TestPop(t *testing.T) {

	functionName := "Push"

	// Test 1
	s1 := NewStack[int]()
	expected1 := []int{1, 2, 3}

	for _, element := range expected1 {
		s1.Push(element)
	}

	lengthS1 := len(expected1)

	for _ = range lengthS1 {

		var lastElement interface{}

		if len(*s1.values) == 0 {
			lastElement = nil
		} else {
			lastElement = (*(*s1).values)[len(*s1.values)-1]
		}

		poppedElement := s1.Pop()

		if poppedElement != lastElement {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, lastElement, poppedElement)
		}
	}

	// Test 2
	s2 := NewStack[float64]()
	expected2 := []float64{.3, 78.9, 30.}

	for _, element := range expected2 {
		s2.Push(element)
	}

	lengthS2 := len(expected2)

	for _ = range lengthS2 {

		var lastElement interface{}

		if len(*s2.values) == 0 {
			lastElement = nil
		} else {
			lastElement = (*(*s2).values)[len(*s2.values)-1]
		}

		poppedElement := s2.Pop()

		if poppedElement != lastElement {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, lastElement, poppedElement)
		}
	}

	// Test 3
	s3 := NewStack[string]()
	expected3 := []string{"z", "c", "t"}

	for _, element := range expected3 {
		s3.Push(element)
	}

	lengthS3 := len(expected3)

	for _ = range lengthS3 {

		var lastElement interface{}

		if len(*s3.values) == 0 {
			lastElement = nil
		} else {
			lastElement = (*(*s3).values)[len(*s3.values)-1]
		}

		poppedElement := s3.Pop()

		if poppedElement != lastElement {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, lastElement, poppedElement)
		}
	}

	// Test 4
	s4 := NewStack[uint]()
	expected4 := []uint{451, 562, 3}

	for _, element := range expected4 {
		s4.Push(element)
	}

	lengthS4 := len(expected4)

	for _ = range lengthS4 {

		var lastElement interface{}

		if len(*s4.values) == 0 {
			lastElement = nil
		} else {
			lastElement = (*(*s4).values)[len(*s4.values)-1]
		}

		poppedElement := s4.Pop()

		if poppedElement != lastElement {
			t.Fatalf("%s fail - expected: %v | result: %v", functionName, lastElement, poppedElement)
		}
	}

}
