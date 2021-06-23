package main

import "fmt"

func main() {
	var num1 = 4
	const num2 = 2
	num1 = 6
	//num2 = 4   //we can't update const's value
	var result1 = num1 + num2
	fmt.Println(result1)
}
