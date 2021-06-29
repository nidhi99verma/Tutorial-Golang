package main

import "fmt"

// func add(x , y int)int{
	//or
func add(x int, y int)int{
	out := x + y
	return out
}

func calc(a, b int)(int, int){
	res1 := a + b
	res2 := a - b
	return res1, res2
}
//or
func solv(a, b int)(res1, res2 int){
	res1 = a * b
	res2 = a / b
	return 
}

func main(){

	num1 := 5
	num2 := 8

	result := add(num1, num2)
	fmt.Println(result)

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)

	result3, result4 := solv(num1, num2)
	fmt.Println(result3, result4)
}