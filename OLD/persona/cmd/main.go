package main

import (
	"fmt"
	"persona"
)



func main() {
	fm := persona.Person{"Filippo", "Montecchio", 52}
	fmt.Println(persona.CheckAge(fm))
}
