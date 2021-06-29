package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 24

	var result = math.Sqrt(num)
	fmt.Printf("this is output: %g \n", result)
	//g: for bigger number
	fmt.Printf("this is output: %.2g \n", result)
	//f: for float number
	fmt.Printf("this is output: %.2f \n", result)

	var intResult = math.Round(result)	
	fmt.Printf("this is output: %.2f \n", intResult)

	var ceilResult = math.Ceil(result)	
	fmt.Printf("this is output: %.2f \n", ceilResult)

	var floorResult = math.Floor(result)	
	fmt.Printf("this is output: %.2f \n", floorResult)

}

//https://golang.org/pkg/fmt/