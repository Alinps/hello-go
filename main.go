package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter the day: ")
	fmt.Scan(&name)
	switch name {
	case "Monday":
		fmt.Println("Starting of the week")
	case "Friday":
		fmt.Println("Weekend coming")
	case "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

}
