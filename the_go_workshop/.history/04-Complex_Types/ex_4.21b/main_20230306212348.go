package main

import "fmt"

// Type definition
type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

type printer interface {
	printInfo()
}

type Toy struct {
	Name  string
	Price float32
}

// Methods definition
func (d Drink) printInfo() {
	fmt.Printf("Drink: %s -- Price: %.2f\n", d.Name, d.Price)
}

func (b Book) printInfo() {
	fmt.Printf("Book: %s -- Price: %.2f\n", b.Title, b.Price)
}

func (t Toy) printInfo() {
	fmt.Printf("Toy: %s -- Price: %.2f\n", t.Name, t.Price)
}

func main() {

	gunslinger := Book{
		Title: "The Gunslinger",
		Price: 4.75,
	}

	coffee := Drink{
		Name:  "Coffee",
		Price: 2.50,
	}

	duck := Toy{
		Name:  "Paperino",
		Price: 1.35,
	}

	// gunslinger.printInfo()
	// coffee.printInfo()
	// duck.printInfo()
	// info := []printer{gunslinger, coffee, duck}
	// info[0].printInfo()
	// info[1].printInfo()
	// info[2].printInfo()

}
