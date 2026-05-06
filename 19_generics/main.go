package main

import "fmt"

// Generics in GO:

// Example: Non generic sums
// SumInts adds together the values of m
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Now Generic function to handle both and once (ints and floats)
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	ints := map[string]int64{
		"first":  32,
		"second": 23,
		"third":  22,
	}

	floats := map[string]float64{
		"first":  32.23,
		"second": 23.32,
		"third":  22.11,
	}
	fmt.Printf("Non Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	// or
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

}
