package main

import (
	"fmt"
	"greetings/greetings"
}

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
