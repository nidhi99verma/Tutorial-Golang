package main

import "fmt"

func main() {
	num := 5
	if num < 5 {
		fmt.Println("num is smaller then 5:", num)
	} else { //}else{  in same line
		fmt.Println("num is greater then 5:", num)
	}

	num1 := 5
	if num <= 5 {
		fmt.Println("num is smaller then 5:", num1)
	} else {
		fmt.Println("num is greater then 5:", num1)
	}

	num2 := 5
	if num2 == 1 {
		fmt.Println("One")
	} else if num2 == 2 {
		fmt.Println("Two")
	} else if num2 == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("None")
	}
 
	//even odd
	num3 := -5
	if (num3 % 2 == 0){
		fmt.Println("even number : ", num3)
	}else{
		fmt.Println("odd number : ", num3)
	}

	//greater number
	num4 := [4]int{4, 5, 6, 7}
	var grtNum int;
	for i := 0; i < 4; i++ {
		if num4[i] > grtNum{
			grtNum = num4[i]
		}
	}
	fmt.Println("Greater number : ",grtNum)
}
