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




func main() {

}
