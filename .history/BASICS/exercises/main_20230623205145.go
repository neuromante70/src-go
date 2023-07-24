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

var userHM = map[int]user{}
