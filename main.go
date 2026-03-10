package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var name string
	var age int
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	fmt.Print("Enter your age")
	fmt.Scan(&age)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ",age)
}
