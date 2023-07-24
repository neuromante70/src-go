package main

import (
	"fmt"
)

m := map[string]int{
	"hello": 5,
	"world": 0,
}

v, ok := m["hello"]
fmt.Println(v, ok)

v, ok = m["world"]
fmt.Println(v, ok)

v, ok = m["world"]
fmt.Println(v, ok)