package main

import "fmt"

var b = 10
var a = 12

func demo(){
	a := 9
	fmt.Println("in demo a:", a)       //function will always give preferance to his own variable(local variable)
}

func main() {
	demo()
	//functional scope variable a (so we can't use this variable)
	//fmt.Println(a)
	//package scope variable b
	fmt.Println(b)
	fmt.Println("in main a:", a)
}

