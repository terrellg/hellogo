package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var locale, greeting string
	var languages = [5]string{"en", "es", "de", "fr", "it"}
	locale = languages[rand.Intn(5)]
	switch locale {
	case "en":
		greeting = "Hello"
	case "de":
		greeting = "Guten Tag"
	case "es":
		greeting = "Hola"

	case "fr":
		greeting = "Bonjour"
	case "it":
		greeting = "Ciao"
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
