package main

import "fmt"

// creating stack using generics

type Stack[T comparable] struct {
	vals []T
}

// Push method
func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

// Pop method
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true

}

// Contain method
func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {

	// Int Stack
	var intStack Stack[int]

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println(intStack)

	fmt.Println("is intStack contains 20: ", intStack.Contains(20))

	v, ok := intStack.Pop()
	fmt.Println(v, ok)

	// String Stack
	var stringStack Stack[string]

	stringStack.Push("Hello")
	stringStack.Push("Hi")
	stringStack.Push("Bye")

	fmt.Println("is stringStack contains 'Bye': ", stringStack.Contains("Bye"))

	fmt.Println(stringStack)

	stringStack.Pop()

	fmt.Println(stringStack)

}
