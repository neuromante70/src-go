package main

import (
	"fmt"
// "errors"
// "sort"
// "strconv"
)

var fprt = fmt.Println

type users struct {
	id    int
	name  string
	email string
}

var listOfUsers []users

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}
	for i := 0; i < 10; i++ {
		listOfUsers = append(listOfUsers, users{id: i, name: "culo", email: "culo@figa.it"})
	}
	fprt(listOfUsers)
}
