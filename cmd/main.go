package main

import (
	"fmt"
	"log"

	"github.com/khelechy/argus"
)

func main() {

	argus, err := argus.Connect(&argus.Argus{
		Username: "testuser",
		Password: "testpassword",
	})

	if err != nil {
		fmt.Println()
		return
	}

	for {
		select {
		case event := <-argus.Events:
			fmt.Println(event.ActionDescription)
		case message := <-argus.Messages:
			log.Println(message)
		case err := <-argus.Errors:
			log.Println("Error:", err)
		}
	}
}
