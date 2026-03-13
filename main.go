package main

import (
	"fmt"
)

func main() {

	ages := map[string]int{
		"Alin":      24,
		"Aneesha":   22,
		"Aleena":    15,
		"Avanthika": 13,
	}
	headers := map[string]string{
		"Content-type":  "application/json",
		"Accept":        "application/json",
		"Authorization": "Bearer token",
	}
	fmt.Println(ages)
	fmt.Println(headers)

}
