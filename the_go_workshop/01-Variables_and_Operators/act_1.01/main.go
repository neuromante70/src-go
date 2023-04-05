package main

import "fmt"

type patients  struct{
	FirstName string
	FamilyName string
	age int
	peanutsAllergy bool
}

func main() {
	var patient1 = patients{"Filippo", "Montecchio", 52, true}
	fmt.Println(patient1)
}