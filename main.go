package main

import (
	"fmt"
	"go-common/number"
	"go-common/string"
)
import "go-common/greeting"

func main() {
	fmt.Println(greeting.SayHello("Khang"))

	fmt.Println(string.Trim(" Hello world "))

	fmt.Println(number.IsPrime(17))
}

