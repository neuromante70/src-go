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




func main() {

}
