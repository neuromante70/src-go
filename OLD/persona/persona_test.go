package persona_test

import (
	"persona"
	"testing"
)

func TestPerson(t *testing.T) {
	t.Parallel()

	_ = persona.Person{
		FirstName: "Culo",
		SecondName: "di Merda",
		Age: 17,
	}
}

func TestCheckAge(t *testing.T) {
	p := persona.Person{
		FirstName: "Culo",
		SecondName: "di Merda",
		Age: 18,
	}
	
	want := "Your adult"
	got := persona.CheckAge(p)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

}