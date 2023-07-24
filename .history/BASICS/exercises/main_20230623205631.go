package main

import (
	"errors"
	"fmt"
	"strconv"
)

type myUser struct {
	id    int
	name  string
	email string
}

var myUserHM = map[int]user{}

func myUserExist(myUserID int) user {
	return myUserHM[myUserID]
}

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}
	myUserHM[us1.id] = us1


culo := myUserExist(us1.id)
fmt.Println(culo)



}