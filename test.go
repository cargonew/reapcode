package main

import "fmt"

func main() {
	fmt.Println("Hello from arch btw")
	res := add()

	fmt.Println("results => " , res)

}

func add() int{ 
	return  3+6
}
