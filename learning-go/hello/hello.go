package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

import "rsc.io/quote"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())
	// handling with log
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Basuko")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Println("=====")
	names := []string{"Chlara", "Hadian", "Yosuke"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
