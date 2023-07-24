package main

import (
	// "errors"
	"fmt"
)

/*
1. The program has a concept [struct] of a User with an id, name, and email - id is int and name/email are strings

2. There should be a hashmap - that has key of type int which will be the user ID and the value will be the user object itself - make this hashmap outside of main so all function can access it

3. There should be 4 functions
  3a. An add function that takes in as an argument a user object and adds it to the hashmap - [bonus return true if the user did not exist before and false if the user was already present]
  3b. A get function that takes in a user id and returns the user from the hashmap IF exists - if not return an error
  3c. A delete function that takes in a user id and deletes that user from the hashmap
  3d. An export function that returns all the users saved in the hashmap as a slice/array
*/

type user struct {
	id    int
	name  string
	email string
}

var db map[int]user{}

// func addUser(us user) (user, bool) {


// 	return user, bool
// }

func main() {
	us1 := user{id: 0, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 1, name: "Jenny Connel", email: "connel@gmail.com"}
	//db = map[int]user
	db[us1.id] = us1
	db[us2.id] = us2

	//_ = errors.New("fai cacare il cazzo")
	fmt.Println(us1, "\n", us2)
	fmt.Println(db)
}