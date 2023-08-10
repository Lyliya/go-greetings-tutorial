package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Lyliya", "Glados", "Chell"}

	messages, err := greetings.Hellos(names)

	if (err != nil) {
		log.Fatal(err)
	}

	for name, greet := range messages {
		fmt.Println(fmt.Sprintf("%v: %v", name, greet))
	}
}