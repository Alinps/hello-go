package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text: ")
	text, _ = reader.ReadString('\n')
	fmt.Println("You typed: ", text)

}
