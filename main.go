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
	fmt.Println(ages)
	fmt.Println("Age of Aneesha: ", ages["Aneesha"])
	fmt.Println("Age of Alin: ", ages["Alin"])
	ages["Aleena"] = 16
	ages["Roshni"] = 23
	fmt.Println("Updated map: ", ages)

	languages := map[string]string{
		"Python":     "Django",
		"Java":       "Spring boot",
		"GoLang":     "Gin",
		"JavaScript": "Express.js",
	}
	languages["C#"] = "ASP.net"

	value, exists := ages["Alinps"]
	fmt.Println("Value: ", value)
	fmt.Println("Exists", exists)
	fmt.Println(len(ages))
	fmt.Println("Iterating map")
	for key, value := range ages {
		fmt.Println(key, value)
	}

}
