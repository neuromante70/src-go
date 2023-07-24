package main

import "fmt"

/*
3. There should be 4 functions
  3a. An add function that takes in as an argument a user object and adds it to the hashmap - [bonus return true if the user did not exist before and false if the user was already present]
  3b. A get function that takes in a user id and returns the user from the hashmap IF exists - if not return an error
  3c. A delete function that takes in a user id and deletes that user from the hashmap
  3d. An export function that returns all the users saved in the hashmap as a slice/array
*/

// 1. The program has a concept [struct] of a User with an id, name, and email - id is int and name/email are strings
type user struct {
	id          int
	name, email string
}




func main() {
	fmt.Println("vim-go")
}
