package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
NOT IMPLEMENTED YET:
3d. An export function that returns all the users saved in the userHM as a slice/array
*/

// 1. The program has a concept [struct] of a User with an id, name, and email - id is int and name/email are strings

type user struct {
	id    int
	name  string
	email string
}

// 2. There should be a hashmap - that has key of type int which will be the user ID and the value will be the user object itself - make this hashmap outside of main so all function can access it

var userHM = map[int]user{}

/* Functions:
3a. An add function that takes in as an argument a user object and adds it to the hashmap -
[bonus return true if the user did not exist before and false if the user was already present]
*/
func addUser(us user) bool {
	if _, exist := userExist(us.id); !exist {
		userHM[us.id] = us
		return true
	}
	return false
}

/*
3b. A get function that takes in a user id and returns the user from the hashmap IF exists - if not return an error
*/
func userExist(usID int) (user, bool) {
	user, ok := userHM[usID]
	return user, ok
}

func findUser(usID int) (user, error) {
	if us, exists := userHM[usID]; exists {
		return us, nil
	}
	return user{}, errors.New("the user with ID " + strconv.Itoa(usID) + " is not in the database")
}

// 3c. A delete function that takes in a user id and deletes that user from the hashmap
func deleteUser(usID int) {
	_, ok := userExist(usID)
	if ok {
		delete(userHM, usID)
		fmt.Println("user ", usID, "correctly deleted")
	} else {
		fmt.Println("user non present, cannot delete it")
	}
}

func printUserHM() {
	for id, user := range userHM {
		fmt.Println(id, "user is", user)
	}
}

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}

	addUser(us1) // instead of userHM[us1.id] = us1
	addUser(us2) // instead of userHM[us2.id] = us2

	// user us3 is not added in userHM
	us3 := user{id: 3, name: "Carl Soft", email: "soft@yahoo.it"}
	_, err := findUser(us3.id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(us3.name, "present in database")
	}

	// now us3 is present in userHM
	addUser(us3)
	_, err = findUser(us3.id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(us3.name, "present in database")
	}

	// we try to add a user already present in userHM
	usDupl := user{id: 3, name: "Carl Soft", email: "soft@yahoo.it"}
	_, err = findUser(usDupl.id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(usDupl.name, "present in database")
	}

	// now we delete a user
	us4 := user{id: 4, name: "John Smith", email: "smith@yahoo.com"}
	addUser(us4)
	fmt.Println(userHM)
	//	fmt.Println(userHM)
	// func deleteUser(usID int) {
	deleteUser(5)      //it should return an error
	deleteUser(us4.id) //this should be work
	fmt.Println(userHM)
	printUserHM()

	/*
		// this user (4) is already present in userHM
		usDupl := user{id: 1, name: "Jenny Connel", email: "connel@gmail.com"}
		addResult = addUser(usDupl)
		if !addResult {
			stringError := "the user with ID " + strconv.Itoa(usDupl.id) + " is already in the database"
			fmt.Println(errors.New(stringError))
		} else {
			fmt.Println(usDupl.name, "added in database")
		}

		// this user (us1) exist in userHM
		fndUs1, err := findUser(us1)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(fndUs1.name, "user present in userHM")
		}

		// the user (us5) is not present in the userHM, hence the program will return an error
		us4 := user{id: 4, name: "John Smith", email: "smith@yahoo.com"}
		fndUs2, err := findUser(us4)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(fndUs2.name, "user present in userHM")
		}

		addUser(us4)
		fmt.Println(userHM, us4)
		deleteUser(us4)
		fmt.Println(userHM)
	*/
}
