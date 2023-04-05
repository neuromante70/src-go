package main

/*
fizzBuzz
• If the number is divisible by 3, print Fizz.
• If the number is divisible by 5, print Buzz.
• If the number is divisible by 15, print FizzBuzz.
*/
import (
	"fmt"
	"strconv"
)

func fizzBuzz(value int) string {
	if value%15 == 0 {
		return "FizzBuzz"
	} else if value%3 == 0 {
		return "Fizz"
	} else if value%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(value)
	}
}

func main() {
	v := 14
	buzzFizz := fizzBuzz(v)
	fmt.Println("Your value is", v, buzzFizz)
}
