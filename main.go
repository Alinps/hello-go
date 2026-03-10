package main

import "fmt"

func main() {
	var day = "Monday"

	if day == "Monday"{
		fmt.Println("Today is Monday")
	}else if day == "Tuesday"{
		fmt.Println("Today is Tuesday")
	}else if day == "Wednesday"{
		fmt.Println("Today is Wednesday")
	}else if day == "Thursday"{
		fmt.Println("Today is Thursday")
	}else if day == "Friday"{
		fmt.Println("Today is Friday")
	}else if day == "Saturday"{
		fmt.Println("Today is Saturday")
	}else if day == "Sunday"{
		fmt.Println("Today is Sunday")
	}else {
		fmt.Println("Invalid day")
	}
}
