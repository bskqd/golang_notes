package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger,
	//including the log entry prefix and a flag
	//to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Glay", "Name", "Surname"}
	// Get a greeting message and print it.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range messages {
		fmt.Printf("%s's greeting: %s\n", key, value)
	}
}
