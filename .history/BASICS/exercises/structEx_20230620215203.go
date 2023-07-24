package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
NOT IMPLEMENTED YET:
  3d. An export function that returns all the users saved in the hashmap as a slice/array
*/

// 1. The program has a concept [struct] of a User with an id, name, and email - id is int and name/email are strings
type user struct {
	id    int
	name  string
	email string
}

// 2. There should be a hashmap - that has key of type int which will be the user ID and the value will be the user object itself - make this hashmap outside of main so all function can access it
var db = map[int]user{}

/* equivalent to:
var db = make(map[int]user)
Why in this case I don't use the curly braces I don't know.
This doesn't work:
var db map[int]user
because the Map isn't initialized and a nil Map with no keys return a runtime panic error.
*/

/*
3. There should be 4 functions:
3a. An add function that takes in as an argument a user object and adds it to the hashmap -
[bonus return true if the user did not exist before and false if the user was already present]
*/
func addUser(us user) bool {
	if userExist(us) {
		db[us.id] = us
		return true
	}
	return false
}

/*
  3b. A get function that takes in a user id and returns the user from the hashmap IF exists - if not return an error
*/
func userExist(us user) bool {
	_, ok := db[us.id]
	return ok
}

func findUser(us user) (user, error) {
	if userExist(us) {
		return us, nil
	}
	return user{}, errors.New("the user with ID " + strconv.Itoa(us.id) + " is not in the database")
}

// 3c. A delete function that takes in a user id and deletes that user from the hashmap
func deleteUser(us user) {
	if userExist()
}


func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}
	//db = map[int]user
	db[us1.id] = us1
	db[us2.id] = us2

	//fmt.Printf("%v \n %v \n", us1, us2)
	//fmt.Println(db)

	// this user (3) is not already added in db
	us3 := user{id: 3, name: "Carl Soft", email: "soft@yahoo.it"}
	addResult := addUser(us3)
	if !addResult {
		stringError := "the user with ID " + strconv.Itoa(us3.id) + " is already in the database"
		fmt.Println(errors.New(stringError))
	} else {
		fmt.Println(us3.name, "added in database")
	}

	// this user (4) is already present in db
	us4 := user{id: 1, name: "Jenny Connel", email: "connel@gmail.com"}
	addResult = addUser(us4)
	if !addResult {
		stringError := "the user with ID " + strconv.Itoa(us4.id) + " is already in the database"
		fmt.Println(errors.New(stringError))
	} else {
		fmt.Println(us4.name, "added in database")
	}

	// this user (us1) exist in db
	fndUs1, err := findUser(us1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fndUs1.name, "user present in db")
	}

	// the user (us5) is not present in the db, hence the program will return an error
	us5 := user{id: 4, name: "John Smith", email: "smith@yahoo.com"}
	fndUs2, err := findUser(us5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fndUs2.name, "user present in db")
	}

}