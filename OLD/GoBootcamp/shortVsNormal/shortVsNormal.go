package main

import "fmt"

func main() {
	// var widht, height = 100, 50 // DON'T!
	widht, height := 100, 50 // 👍
	fmt.Println(widht, height)

	// DON'T!
	// width = 50
	// color := "red"
	widht, color := 50, "red" // idiomatic Go redeclaration 👍
	fmt.Println(widht, color)
}
