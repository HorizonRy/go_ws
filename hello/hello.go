package main

import (
	"fmt"

	"github.com/HorizonRy/go_ws/tree/main/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
