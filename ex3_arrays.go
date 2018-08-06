package main

import "fmt"

// 2D matrix type
type Matrix [3][3]int

func main() {
	// array example (type string)
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	m := Matrix{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
	}

	fmt.Println(m)


	// looping over array

	for i, n := range names {
		fmt.Printf("%d - %s\n", i, n)
	}
}
