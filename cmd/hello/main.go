package main

import (
	"fmt"
	utils "github.com/Zenrock-Foundation/go-mod-test/v5/utils"
)

func main() {
	fmt.Println("Hello world from the fork")
	f := utils.Fibonacci(10)
	fmt.Println("Fibonacci for number 10 is: ", f)
}
