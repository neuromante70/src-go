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

