package main

import "gopkg.in/gookit/color.v1"

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	welcome()

	switch action() {
	case 1:
		monitoring()
	case 2:
		showLogs()
	case 3:
		exit()
	default:
		fmt.Println("Command not found")
		os.Exit(-1)
	}
}
func welcome() {
	var name string
	var version float32 = 1.1
	fmt.Print("Say ur name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Welcome", name, "this app is on version:", version)
}
func action() int {
	fmt.Println("--------------------------------------------------")
	fmt.Println("1- Begin Monitoring - CTRL + C to stop")
	fmt.Println("2- Show Logs")
	fmt.Println("3- Exit program")
	var action int
	fmt.Scan(&action)
	return action
}
func monitoring() {
	for {
		fmt.Println("Monitoring...")
		file, err := os.Open("./sites.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// each line is a url to verify

		for scanner.Scan() {
			site := scanner.Text()
			resp, _ := http.Get(site)
			red := color.FgRed.Render
			green := color.FgGreen.Render
			if resp.StatusCode == 200 {
				fmt.Println(site, green(resp.Status))
				writeLog(site, resp.Status)
			} else {
				fmt.Println(site, red(resp.Status))
				writeLog(site, resp.Status)
			}
		}
		file.Close()
		fmt.Println("-------------------")
		fmt.Println("Will verify again....")
		time.Sleep(2 * time.Second)
	}
}
func exit() {
	fmt.Println("Exit..")
	os.Exit(0)
}

func writeLog(site string, status string) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	t := time.Now()
	file.WriteString("Site: " + site + " - " + status + "; " + t.Format("2006-01-02 15:04:05") + "\n")
	file.Close()
}
func showLogs() {
	if _, err := os.Stat("logs.txt"); os.IsNotExist(err) {
		fmt.Println("No logs to show")
		os.Exit(0)
	}
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
