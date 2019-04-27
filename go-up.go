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
	fmt.Println("1- Begin Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("3- Exit program")
	var action int
	fmt.Scan(&action)
	return action
}
func showLogs() {
	fmt.Println("Show logs")
}
func monitoring() {
	fmt.Println("Monitoring...")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	fmt.Println(resp.Status)
}
func exit() {
	fmt.Println("Exit..")
	os.Exit(0)
}
