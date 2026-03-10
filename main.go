package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var name string
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	fmt.Println("Name: ", name)
}
