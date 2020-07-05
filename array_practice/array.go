package array_practice

import (
	"fmt"
)

func CreateArrays() {
	grades := [5]int{87, 68, 97, 92, 83}
	fmt.Printf("Array %v", grades)

	numbers := [...]int{1, 2, 3}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %v, Value: %v", i, numbers[i])
	}

	gpa := []float32{4.00, 3.50, 5.00, 5.00, 4.00}
	for _, val := range gpa {
		fmt.Println(val)
	}

	var x [3]int
	fmt.Println(x)
	x[1] = 100
	x[2] = 200
	fmt.Println(x)

	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1}}
	for _, val := range identityMatrix {
		for _, item := range val {
			fmt.Printf("%v ", item)
		}
		fmt.Println()
	}

	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[0] = [3]int{4, 5, 6}
	matrix[0] = [3]int{7, 8, 9}
	fmt.Println(matrix)

	a := [...]int{1, 2, 5}
	b := a
	b[1] = 3
	fmt.Println(a)
	fmt.Println(b)

	c := &a
	c[1] = 10
	fmt.Println(a)
	fmt.Println(c)
}
