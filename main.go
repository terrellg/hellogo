package main

import "fmt"

func main() {
	var locale, greeting string
	var languages = [4]string{"en", "es", "de", "fr"}
	locale = languages[0]
	switch locale {
	case "en":
		greeting = "Hello"
	case "de":
		greeting = "Guten Tag"
	case "es":
		greeting = "Hola"

	case "fr":
		greeting = "Bonjour"
	default:
		greeting = "Yurrrrrr"
	}
	/*
		if locale == "en" {
			greeting = "Hello"

		} else if locale == "es" {
			greeting = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else {
			greeting = "Yurrr"
		}
	*/
	fmt.Printf(greeting + " Go\n")

}
