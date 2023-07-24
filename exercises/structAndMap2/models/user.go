package user

// 1. The program has a concept [struct] of a User with an id, name, and email - id is int and name/email are strings
type User struct {
	Id    int
	Name  string
	Email string
}

// 2. There should be a hashmap - that has key of type int which will be the user ID and the value will be the user object itself - make this hashmap outside of main so all function can access it
var UserHM = map[int]User{}
