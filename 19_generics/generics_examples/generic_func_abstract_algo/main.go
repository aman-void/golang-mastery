package main

import "fmt"

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two params, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {

	r := make([]T2, len(s))

	for i, v := range s {
		r[i] = f(v)
	}

	return r

}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reduce reduces a []T1 to a single value using reduction
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {

	r := initializer

	for _, v := range s {
		r = f(r, v)
	}

	return r

}

func main() {

	// Map usage
	numSlice := []int{12, 13, 14, 15, 16}
	squared := Map(numSlice, func(n int) int {
		return n * n
	})

	fmt.Printf("Square slice of %v = %v\n", numSlice, squared)

	// Filter usage
	evenNumbers := Filter(numSlice, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even numbers in %v = %v\n", numSlice, evenNumbers)

	// Reduce usage
	sum := Reduce(numSlice, 0, func(acc, val int) int {
		return acc + val
	})

	fmt.Println(sum)
}
