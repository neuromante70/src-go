package main

import (
	"errors"
	"fmt"
	"strconv"
)

type user struct {
	id    int
	name  string
	email string
}

var userHM = map[int]user{}

func userExist(userID int) (user, bool) {
	user, ok := userHM[userID]
	return user, ok
}

func findUser(userID int) (user, error) {
	us, exists := userExist(userID)
	if exists {
		return us, nil
	}
	return user{}, errors.New("the user with ID " + strconv.Itoa(userID) + " is not in the database")
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

func deleteUser(userID int) {
	_, ok := userExist(userID)
	if ok {
		delete(userHM, userID)
		fmt.Println("user", userID, "correctly deleted")
		return
	}
	fmt.Println("user non present, cannot delete it")
}


// ancillary function
func printUserHM() {
	for id, user := range userHM {
		fmt.Println(id, "user is", user)
	}
}

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}

	addUser(us1) // instead of userHM[us1.id] = us1
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
	exportAsSlice()

}

