package main

/*
#cgo  CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L. -lmediator -lstdc++
#include "mediator.h"
*/

import (
	"C"
	"fmt"
)

func main() {
	a, b := 1, 10
	result := C.add(C.int(a), C.int(b))
	fmt.Println("result: ", int(result))
}
