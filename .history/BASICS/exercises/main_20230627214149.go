package main

import (
	"fmt"
// "errors"
// "sort"
// "strconv"
)

type fprt = fmt.Println

type users struct {
	id    int
	name  string
	email string
}

var listOfUsers []users

func main() {
	for i := 0; i < 10; i++ {
		listOfUsers = append(listOfUsers, user{})
	}
	fprt(listOfUsers)
}
