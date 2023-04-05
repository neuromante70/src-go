package main

// Printing a User's Name Based on User Input

import (
	"fmt"
	"os"
	"strconv"
)

func getPassedArgs() int {
	if len(os.Args) > 2 {
		fmt.Printf("Only one arguments is admitted.\n")
		os.Exit(1)
	}
	number, _ := strconv.Atoi(os.Args[1])
	return number
}

func main() {
	// employees := make(map[int]string)
	// employees[73] = "Tracy"
	// employees[201] = "Bob"
	// employees[305] = "Sue"
	// employees[631] = "Jake"

	employees := map[int]string{
		73:  "Tracy",
		201: "Bob",
		305: "Sue",
		631: "Jake",
	}
	// fmt.Println(employees)
	idEmployee := getPassedArgs()
	fmt.Println("Your ID is:", idEmployee)
	value, ok := employees[idEmployee]
	if ok {
		fmt.Println("and the employee is:", value)
	} else {
		fmt.Println("There's  no employee with this ID.")
	}
}
