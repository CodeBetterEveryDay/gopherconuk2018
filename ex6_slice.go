package main

import "fmt"

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	// Create a new slice called family by joining the parents and kids slice together
	fmt.Println("=== family  ===")
	family := []string{}
	family = append(family, parents...)
	family = append(family, kids...)

	fmt.Println(family)


	// improved - joining two slices and creating third one
	fmt.Println("=== family2 ===")
	family2 := append(parents, kids...)
	fmt.Println(family2)

	fmt.Println("===")

	// Fix the following bugs - array capacity doesn't match required slice length

	// note: fixed by defining lenght 1 and capacity  1 and adding the str "Alice" to the first elem
	//       of the list
	extras := make([]string, 1, 1)
	extras[0] = "Alice"
	fmt.Println(extras)


}
