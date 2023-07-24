package main

import (
	"fmt"
	"path"
)

/* path.Split()
func Split(path string) (dir, file string)
*/
func main() {
	// var dir, file string
	_, file := path.Split("css/main.css")

	// fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
