package main

import "fmt"

func main()  {
	// create a slice
	nameSlice := []string{"John", "Paul", "George", "Ringo"}
	for _, name := range nameSlice {
		fmt.Println(name)
	}

	// appending
	names := []string{}
	names = append(names, "Bolek")
	names = append(names, "Lolek")

	fmt.Println(names)

	// append an entire slice to another slice
	fmt.Println("cap:", cap(names))
	names = append(names, "Tosia")

	fmt.Println("cap:", cap(names))

	// using make keyword, explicitly set capacity
	fmt.Println("===")
	a := make([]int, 1, 3)

	fmt.Println(a)
	fmt.Println(len(a)) // 1
	fmt.Println(cap(a))


}
