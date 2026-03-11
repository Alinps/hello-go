package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the text: ")
	scanner.Scan()
	text = scanner.Text()
	fmt.Println("You typed: ", text)

}
