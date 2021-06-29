package main

import (
	"fmt"
	"time"
)

func main() {
	num := 2
	switch num{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("None")
	}

	//

	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday{
	case today + 0:
		fmt.Println("Today😄.")
	case today + 1:
		fmt.Println("Tommorow😊.")
	case today + 2:
		fmt.Println("In two days🤨.")
	default:
		fmt.Println("Too far away😐.")
	}

	fmt.Println(time.Now())
}
