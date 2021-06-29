package main

import "fmt"

func main() {
	i := 1                        //initialization
	for i <= 5 {				  //condition
		fmt.Println("Nidhi", i)
		i++						 //increment
	}
	
	//or

	for i := 1; i <= 5; i++{
		fmt.Println("Ram", i)
	}
}
