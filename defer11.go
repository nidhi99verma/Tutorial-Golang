package main

import "fmt"

// func main() {
// 	fun1()
// }

// func fun1() {
// 	fmt.Println("fun1 start")
// 	//fun2()                                      //defer will run last
// 	defer fun2()
// 	fmt.Println("fun1 end")
// }

// func fun2() {
// 	fmt.Println("In fun2")
// }


func main() {
	for i := 0; i < 10; i++{
		defer fmt.Println(i)                    //defer store loop's output in stack(lifo) and print all out after 'bye' so we get loop output revers
	}
	fmt.Println("Bye...!")
}

