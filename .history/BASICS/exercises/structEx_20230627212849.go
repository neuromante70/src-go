package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

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

func userExist(userID int) (user, bool) {
	user, ok := userHM[userID]
	return user, ok
}

func addUser(us user) bool {
	_, exist := userExist(us.id)
	if exist {
		fmt.Println("The user", us.name, "is already present in userHM")
		return false
	}
	userHM[us.id] = us //add the user to the userHM after a check
	fmt.Println("The user", us.name, "wasn't present in userHM and now is correctly inserted")
	return true
}

/*
3b. A get function that takes in a user id and returns the user from the hashmap IF exists - if not return an error
*/
func findUser(userID int) (user, error) {
	us, exists := userExist(userID)
	if exists {
		return us, nil
	}
	return user{}, errors.New("the user with ID " + strconv.Itoa(userID) + " is not in the database")
}

// 3c. A delete function that takes in a user id and deletes that user from the hashmap
func deleteUser(userID int) {
	_, ok := userExist(userID)
	if ok {
		delete(userHM, userID)
		fmt.Println("user", userID, "correctly deleted")
		return
	}
	fmt.Println("user non present, cannot delete it")
}

/*
3d. An export function that returns all the users saved in the userHM as a slice/array
The array is simply an array of user structs []user{} - just how the map has values of user structs
If the DB map can hold user objects - an array can too - once a struct has been defined it can be used as a regular go type
1. the export function needs to return []user not []string
2. The for loop wont work if the ids are not sequential and have gaps - you want to use a special for loop for maps which is
for key, element := range userdb {
// add element to user array
}
When dealing with maps we think in terms of key and value - the loop is also over the key and value in the map not from i=0 to some index
*/

func exportAsSlice() []user {
	userSL := make([]user, len(userHM))
	for key, element := range userHM {
		userSL[key] = userHM[]
	}
	return userSL
}

// ancillary functions
// sort and print the users Hashmap
func printUserHM() {
	keys := make([]int, 0, len(userHM))
	for k := range userHM {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, userHM[k])
	}
}

// print the users slice
func printUserSL(userSL []user) {
	for id, user := range userSL {
		fmt.Println(id, "user is", user)
	}
}

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}

	addUser(us1) // instead of userHM[us1.id] = us1, but with check of validity
	addUser(us2)

	// user us3 is not added in userHM
	us3 := user{id: 3, name: "Carl Soft", email: "soft@yahoo.it"}
	_, err := findUser(us3.id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(us3.name, "present in database")
	}

	// this time us3 is present in userHM
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
		fmt.Println("duplicate user")
	} else {
		fmt.Println(usDupl.name, "present in database")
	}

	// now we delete a user
	us4 := user{id: 4, name: "John Smith", email: "smith@yahoo.com"}
	addUser(us4)
	deleteUser(5)      //it returns an error
	deleteUser(us4.id) //this should be work
	printUserHM()

	userSlice := exportAsSlice()
	printUserSL(userSlice)

}
