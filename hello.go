package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		command := showIntroduction()
		showMenu(command)
	}
}

func showIntroduction() int {
	/* When working with Go, every function coming from an external package starts with an uppercase letter. */
	/* name := "Allas" This is a way of declaring a variable without specifying its type. */
	fmt.Println("Hello! How's it going?")
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
		startMonitoring()
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

func startMonitoring() {
	fmt.Println("Starting monitoring..")
	site := "https://www.google.com"

	// resp, error := http.Get(site)
	resp, _ := http.Get(site)
	// The _ indicates that you want to ignore that respective variable
	if resp.StatusCode == 200 {
		fmt.Println("Everything is fine with the website:", site)
	} else {
		fmt.Println("Unfortunately, the website ", site, "isn't working well, the status code is: ", resp.StatusCode)
	}
}
