package main

import (
	"fmt"
)

func main() {
	euWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("European week:", euWeek)
	usWeek := append(euWeek[6:], euWeek[:6]...)
	// usWeek := make([]string, 7)
	// usWeek[0] = euWeek[6]
	// for i := 1; i < len(euWeek); i++ {
	// 	usWeek[i] = euWeek[i-1]
	// }
	fmt.Println("American week:", usWeek)
}
