package main

import (
	"errors"
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

var db = map[int]user{}

/* equivalent to:
var db = make(map[int]user)
Why in this case I don't use the curly braces I don't know.
This doesn't work:
var db map[int]user
because the Map isn't initialized and a nil Map with no kyes return a runtime panic error.
*/

func addUser(us user) (map[int]user, bool) {
	if _, ok := db[us.id]; ok {
		return nil, false
	}
	db[us.id] = us
	return db, true
}

func main() {
	us1 := user{id: 0, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 1, name: "Jenny Connel", email: "connel@gmail.com"}
	//db = map[int]user
	db[us1.id] = us1
	db[us2.id] = us2

	_ = errors.New("Not the expected result")
	fmt.Println(us1, "\n", us2)
	fmt.Println(db)
	us3 := user{id:2, name: "Carl Soft", email: "soft@yahoo.it"}
	db, err := addUser(us3)
	if !(err) {
		fmt.Println(err.Error())
	return
	}
	fmt.Println(resultDivide.answer)

}
