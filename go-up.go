package main

import "fmt"

func main() {
	welcome()
}
func welcome() {
	var name string
	var version float32 = 1.1
	fmt.Print("Say ur name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Welcome", name, "this app is on version:", version)
	fmt.Println(action())
}
func action() int {
	fmt.Println("--------------------------------------------------")
	fmt.Println("1- Begin Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("3- Exit program")
	var action int
	fmt.Scanf("%d", &action)
	return action
}
