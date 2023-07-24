package main

import (
	"fmt"
	// "errors"
	// "strconv"
)

type myUser struct {
	id    int
	name  string
	email string
}

var myUserHM = map[int]myUser{}

func myUserExist(myUserID int) myUser {
	return myUserHM[myUserID]
}

func main() {
	us1 := myUser{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := myUser{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}
	myUserHM[us1.id] = us1
	myUserHM[us2.id] = us2

	culo := myUserExist(us1.id)
	fmt.Println(culo)
	culo2

}
