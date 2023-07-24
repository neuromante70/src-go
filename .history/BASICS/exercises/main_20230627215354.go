package main

import (
	"fmt"
	// "errors"
	// "sort"
	// "strconv"
)

var fprt = fmt.Println

type user struct {
	id    int
	name  string
	email string
}

var userHM = map[int]user{}
var listOfUsers []users

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}
	us3 := user{id: 3, name: "Carl Soft", email: "soft@yahoo.it"}
	userHM[us1.id] = us1
	userHM[us2.id] = us2
	userHM[us3.id] = us3

	for i := 0; i < 10; i++ {
		listOfUsers = append(listOfUsers, users{id: i, name: "culo", email: "culo@figa.it"})
	}
	fprt(userHM)
	fprt(listOfUsers)
}
