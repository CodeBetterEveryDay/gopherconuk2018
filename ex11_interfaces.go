package main

import "fmt"

type User struct {
	ID    int
	First string
	Last  string
}

func (u User) String() string {
	return fmt.Sprintf("User %d is %v %v", u.ID, u.First, u.Last)
}

// Task:
// Satisfy the stringer interface (https://golang.org/pkg/fmt/#Stringer)
// so that the User struct will print
// User <ID> is <First> <Last>
//
// example:
//      User 1 is Rob Pike

func main() {
	u := User{ID: 1, First: "Rob", Last: "Pike"}
	fmt.Println(u)
}
