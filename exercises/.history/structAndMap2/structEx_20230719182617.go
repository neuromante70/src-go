package main

import (
	user "BASICS/exercises/structAndMap2/models"
	"errors"
	"fmt"
	"strconv"
)

var prt = fmt.Println // just for fun

// ancillary functions:
func (u) userExist(usID int) (u user.User, exists bool) {
	// I tried to convert the function as method of user.User but golang doesn't allow it
	// I like a lot named return values in functions
	u, exists = user.UserHM[usID]
	return u, exists
}

/*
3a. An add function that takes in as an argument a user object and adds it to the hashmap
[bonus return true if the user did not exist before and false if the user was already present]
*/
func addUserToHM(u user.User) (exists bool) {
	// _, exists := userExist(u)
	// if !exists {
	// I don't like the syntax below, I prefer the other version above.
	if _, exists := userExist(u.Id); exists {
		prt("user ", u.Id, "correctly added")
		return !exists // the ok, comma idiom in Map return true if the key exists,
	} // and I want true if the key doesn't exist because this prove the record will be registered
	user.UserHM[u.Id] = u
	prt("user id", u.Id, "correctly added")
	return !exists
}

/*
3b. A get function that takes in a user id and returns the user from the hashmap IF exists - if not return an error
*/
func returnUserFromHM(usID int) (u user.User, notFound error) {
	u, exists := user.UserHM[usID]
	if exists {
		return u, nil
	}
	return user.User{}, errors.New("User with id " + strconv.Itoa(usID) + " is not present in DB")
}

/*
3c. A delete function that takes in a user id and deletes that user from the hashmap
*/
func deleteUserFromHM(usID int) (result bool) {
	_, exists := user.UserHM[usID]
	if !exists {
		prt("user ", usID, " not present in DB or already deleted, unable to delete it")
		return !exists
	} else {
		delete(user.UserHM, usID)
		prt("user id ", usID, " correctly deleted")
		return exists
	}
}

/*
3d. An export function that returns all the users saved in the hashmap as a slice/array
*/
var UserSL []user.User

// var userSL []user.User
func exportAsSlice() []user.User {
	for i := range user.UserHM {
		UserSL = append(UserSL, user.UserHM[i])
	}
	return UserSL
}

func main() {
	us1 := user.User{Id: 1,
		Name:  "Phil Cazzaniga",
		Email: "cazzaniga@gmail.com",
	}
	us2 := user.User{Id: 2, Name: "Jenny Connelly", Email: "connelly@gmail.com"}
	us3 := user.User{Id: 3, Name: "Carl Soft", Email: "soft@yahoo.it"}

	_ = addUserToHM(us1)

	_ = addUserToHM(us1)

	u, err := returnUserFromHM(us1.Id)
	if err != nil {
		prt(err)
	} else {
		prt("The searched user is:", u)
		prt(user.UserHM[us1.Id])
	}

	u, err = returnUserFromHM(us2.Id) // us2 exist as user but not in Map
	if err != nil {
		prt(err)
	} else {
		prt("The searched user is:", u)
		prt(user.UserHM[us2.Id])
	}
	_ = addUserToHM(us2)

	_ = deleteUserFromHM(us2.Id)
	_ = deleteUserFromHM(us2.Id)
	_ = addUserToHM(us2)
	_ = addUserToHM(us3)

	mySlice := exportAsSlice()
	prt("This is the exported slice: ", mySlice)

	for i := range user.UserHM {
		prt("This is the Map user: ", user.UserHM[i])
	}
}
