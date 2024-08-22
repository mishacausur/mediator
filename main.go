package main

//#include "wrap_sum.hxx"

import (
	"C"
	"fmt"
)

func main() {
	a, b := 1, 10
	result := a + b
	fmt.Println("result: ", int(result))
}
