package main

import (
	"fmt"
	utils "github.com/Zenrock-Foundation/go-mod-test/v4/utils"
)

func main() {
	fmt.Println("Hello world")
	f := utils.Fibonacci(10)
	fmt.Println("Fibonacci for number 10 is: ", f)
}
