package main

import (
	"fmt"
)

func main() {
	us1 := user{id: 1, name: "Phil Cazzaniga", email: "cazzaniga@gmail.com"}
	us2 := user{id: 2, name: "Jenny Connelly", email: "connelly@gmail.com"}
	us3 := user{id: 3, name: "Carl Soft", email: "soft@yahoo.it"}
	userHM[us1.id] = us1
	userHM[us2.id] = us2
	userHM[us3.id] = us3

	for _, user := range userHM {
		userSL = append(userSL, user)
	}
	fprt(userHM)
	fprt(userSL)
}