package main

import "fmt"

func main() {
	var locale string = "de"
	var greeting string = "Hello"

	if locale == "en" {
		greeting = "Hello"

	} else if locale == "es" {
		greeting = "Hola"
	} else if locale == "de" {
		greeting = "Guten Tag"
	} else {
		greeting = "Yurrr"
	}

	fmt.Printf(greeting + " Go\n")

}
