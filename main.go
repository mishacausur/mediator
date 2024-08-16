package main

//#include "wrap_sum.hxx"

import (
	"C"
	"fmt"
)

func main() {
	a, b := 1, 10
	result := C.add_from_cplus(C.int(a), C.int(b))
	fmt.Println("result: ", int(result))
}
