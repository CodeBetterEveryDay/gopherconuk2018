package main

import "fmt"

func main()  {
	var command string
	var valid bool

	fmt.Printf("var command is of type %T and has a value %+v\n", command, command)
	fmt.Printf("var valid is of type %T and has a value %+v\n", valid, valid)

	var foo int = 5
	var bar bool = true

	fmt.Printf("var foo is of type %T and has value %+v\n", foo, foo)
	fmt.Printf("var bar is of type %T and has value %+v\n", bar, bar)
}