package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greeings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
