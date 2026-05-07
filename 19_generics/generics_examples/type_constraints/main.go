package main

import "fmt"

// * Type Constraints defines : what types are allowed for generic type parameter
// * Without Type Constraints, generic code would be too unsafe because the compiler would not know which operations are valid.

// For Example:

// func Add[T any](a, b T) T {
// 	return a + b
// }

// ! NOTE: The compiler yell at above code because we use any type and it is safe when we pass int or string but what about struct (it doesn't support +) so the correct version

type Number interface {
	int | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

// ! One interesting things: The tilde ~ problem
// * ~ this means match any type whose underlying type is this type.

// Example: without this ~
type MyInt int

/* Without ~
type Num interface {
	int
}
*/

// With ~
type Num interface {
	~int
}

func Double[T Num](v T) T {
	return v * 2
}

// REASON:
/*
* MyInt != int
* Even though: type MyInt int
* MyInt is the NEW named type
* Go treats it as distinct from int
 */

// Exercise: 1 calculate the largest stock of items on a ranch
func FindLargestRanchStock[K comparable, V int | float64](m map[K]V) K {

	var stock V
	var name K

	for k, v := range m {
		if v > stock {
			stock = v
			name = k
		}
	}
	return name

}

func main() {

	fmt.Println(Add(12, 13))
	fmt.Println(Add(12.2, 13.3))

	// Without this ~ Failed
	// With this ~ works
	var x MyInt = 10
	fmt.Println(Double(x))

	// Exercise 1.
	animalStock := map[string]int{
		"chicken": 10,
		"cattle":  20,
		"horses":  5,
	}

	miscStock := map[string]float64{
		"Hay":        5.5,
		"Feed":       1.4,
		"Fertilizer": 4.5,
	}

	largestStockOnRanchInt := FindLargestRanchStock(animalStock)
	fmt.Printf("The largest stock item on the ranch is %s\n", largestStockOnRanchInt)

	largestStockOnRanchFloat := FindLargestRanchStock(miscStock)
	fmt.Printf("The largest stock item on the ranch is %s\n", largestStockOnRanchFloat)
}
