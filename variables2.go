package main

import "fmt"

func main() {
	var num1 = 2
	var num2 = 4
	var result1 = num1 + num2
	fmt.Println(result1);

	// default int variable value 0
	var num3 int            
	var num4 int
	var result2 = num3 + num4
	fmt.Println(result2);
	
	num3 = 5
	num4 = 7
	result2 = num3 + num4
	fmt.Println(result2);

	var num5, num6 int
	num5, num6 = 6, 8
	var result3 = num5 + num6
	fmt.Println(result3);

	result4 := 16            //sorthand of var result = 16
	fmt.Println(result4);
}