package main

import "fmt"

func main()  {
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawberry"}

	// Use a 'classic' for loop  to print out each fruit in the array, and the
	// corresponding index.

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index: %d, fruit: %s\n", i, fruits[i])
	}


	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.

	for i, fruit := range fruits {
		fmt.Printf("index: %d, fruit name: %s\n", i, fruit)
	}

	// Using the keyword `continue`, skip every other fruit (even ones)

	fmt.Println("===")
	for i, fruit := range fruits {
		if i % 2 != 0 {
			continue
		}
		fmt.Printf("index: %d, fruit name: %s\n", i, fruit)
	}
	
}