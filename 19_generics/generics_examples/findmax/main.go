package main

import (
	"cmp"
	"fmt"
)

// Writing find max integer/float function without generics and with generics

// Without Generics (Repetitive code)

func maxInt(intSlice []int) int {

	maxInt := intSlice[0]

	for _, num := range intSlice {
		if num > maxInt {
			maxInt = num
		}
	}

	return maxInt

}

func maxFloat(floatSlice []float64) float64 {
	maxFloat := floatSlice[0]

	for _, floatNum := range floatSlice {
		if floatNum > maxFloat {
			maxFloat = floatNum
		}
	}

	return maxFloat

}

// With generics
func maxIntOrFloat[T int | float64](s []T) T {

	max := s[0]

	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max

}

// Another example find min

func findMinGeneric[T cmp.Ordered](s []T) T {

	min := s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func main() {

	fmt.Println("Without generics")
	fmt.Println(maxInt([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxFloat([]float64{1.2, 3.34, 4.2, 5.2}))

	fmt.Println("With generics")
	fmt.Println(maxIntOrFloat([]int{12, 34, 21, 24}))
	fmt.Println(maxIntOrFloat([]float64{12.32, 34.22, 21.45, 24.33}))

	fmt.Println("Find Min Generics")
	fmt.Println(findMinGeneric([]int{12, 34, 21, 24}))
	fmt.Println(findMinGeneric([]float64{12.32, 34.22, 21.45, 24.33}))
	fmt.Println(findMinGeneric([]string{"apple", "banana", "cherry"}))
	
}
