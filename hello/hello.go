package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greeings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
