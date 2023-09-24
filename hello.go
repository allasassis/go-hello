package main

import (
	"fmt"
	"os"
)

func main() {
	command := showIntroduction()
	showMenu(command)
}

func showIntroduction() int {
	/* When working with Go, every function coming from an external package starts with an uppercase letter. */
	/* name := "Allas" This is a way of declaring a variable without specifying its type. */

	fmt.Println("Hey, what's your name?")
	var name string
	fmt.Scan(&name)

	fmt.Println("Hello,", name, "! How's it going?")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command)

	return command
}

func showMenu(command int) {
	switch command {
	case 1:
		fmt.Println("Starting monitoring!")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting the program...")
		os.Exit(0)
	default:
		fmt.Println("I do not recognize your command.")
		os.Exit(-1)
	}
}
