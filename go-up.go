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
}
