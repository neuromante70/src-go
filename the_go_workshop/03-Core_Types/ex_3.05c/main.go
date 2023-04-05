package main

// TEXT: RUNE

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	fmt.Println()
	// Convert out bytes back to strings

	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}
	/* It doesn't work as expected, because...
	   OUTPUT: S i r _ K i n g _ Ã  b e r %
	   That's because Ü needs more than one byte to be encoded
	   To safely work with interindividual characters of a multi-byte string,
	   you first must convert the strings of slice of byte types to a slice
	   of rune types.
	*/
	fmt.Println()
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}

	fmt.Println()
	// Length of a string
	fmt.Println("Bytes: ", len(username))
	fmt.Println("Runes: ", len([]rune(username)))

	// Limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))

}
