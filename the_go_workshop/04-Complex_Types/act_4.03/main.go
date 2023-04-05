package main

// Creating a Locale Checker
import (
	"fmt"
	"os"
)

func getPassedArgs() string {
	// var args string
	args := os.Args[1]
	return args
}

type locales struct {
	language string
	country  string
}

func newLocales(language string, country string) string {
	l := locales{language: language, country: country}
	return l.language + "_" + l.country
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	australia := newLocales("en", "AU")
	usa := newLocales("en", "US")
	canada := newLocales("en", "CN")
	canadaFrench := newLocales("fr", "CN")
	russia := newLocales("ru", "RU")

	var countries []string
	countries = append(countries, australia, usa, canada, canadaFrench, russia)

	fmt.Println("The locales for my countries are:", countries)
	countryToSearch := getPassedArgs()
	if contains(countries, countryToSearch) {
		fmt.Println("You have this locale")
	} else {
		fmt.Println("Sorry, this locale is not implemented yet")
	}

}
