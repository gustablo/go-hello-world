package main

import (
	"example/hello/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	names := []string{"Gustavo", "Igor", "Luska"}

	messages, err := greetings.MultipleHellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
