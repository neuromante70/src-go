package persona

type Person struct {
	FirstName  string
	SecondName string
	Age        int
}

func CheckAge(p Person) string {
	if p.Age >= 18 {
		return "Your adult"
	}
	return "Your not allowed to buy things here"
}
