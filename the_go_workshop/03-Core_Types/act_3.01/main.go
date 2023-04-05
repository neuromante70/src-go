package main

/*
|-------|-------|---------------|
|Item	|Cost	|Sales Tax Rate |
|-------|-------|---------------|
|Cake	|0.99	|7.5%			|
|Milk	|2.75	|1.5%			|
|Butter	|0.87	|2%				|
|_______|_______|_______________|
*/

import (
	"fmt"
)

type item struct {
	name    string
	cost    float64
	taxRate float64
}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}

func main() {
	var item1 = item{"Cake", 0.99, 0.075}
	var item2 = item{"Milk", 2.75, 0.015}
	var item3 = item{"Butter", 0.87, 0.02}

	sumTax := salesTax(item1.cost, item1.taxRate)
	sumTax += salesTax(item2.cost, item2.taxRate)
	sumTax += salesTax(item3.cost, item3.taxRate)

	fmt.Println("Sales Tax Total:", sumTax)
}
