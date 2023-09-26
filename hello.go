package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 0

func main() {
	for {
		fmt.Println("")
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
		showLogs()
	case 0:
		fmt.Println("Exiting the program...")
		os.Exit(0)
	default:
		fmt.Println("I do not recognize your command.")
		os.Exit(-1)
	}
}

func startMonitoring() {
	// sites := []string{"https://www.google.com", "https://www.youtube.com/", "https://twitter.com/home"}
	// sites = append(sites, "https://www.twitch.tv/")
	fmt.Println("Starting monitoring...")
	sites := readFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		// Pause the loop for 5 seconds before continuing.
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Everything is fine with the website:", site)
			registerLog(site, true)
		} else {
			fmt.Println("The website", site, "isn't working well. The status code is:", resp.Status)
			registerLog(site, false)
		}
	}

}

func readFile() []string {
	sites := []string{}
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error occurred:", err.Error())
	}

	reader := bufio.NewReader(file)
	for {
		line, err1 := reader.ReadString('\n')

		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			fmt.Println("An error occurred:", err1.Error())
		}
		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	time := time.Now().Format("02/01/2006 15:04:05")

	if status {
		file.WriteString(time + " " + site + " IS online\n")
	} else {
		file.WriteString(time + " " + site + " IS NOT online\n")
	}

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Println(string(file))
}
