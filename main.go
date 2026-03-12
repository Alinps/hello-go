package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Alin p s"
	var firstname string = "Aneesha"
	var lastname string = "Alin"
	var text string = "      hello world    "
	var language string = "I love, python"
	fmt.Println("length:", len(name))
	fmt.Println(firstname + " " + lastname)
	fmt.Println(firstname == lastname)
	fmt.Println(firstname != lastname)
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	fmt.Println("before triming whitespace: ", text)
	fmt.Println("After triming whitespace: ", strings.TrimSpace(text))
	fmt.Println(strings.Replace(language, "python", "golang", 1))
	fmt.Println(strings.Split(language, ""))

}
