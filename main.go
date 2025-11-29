package main

import (
	"fmt"

	"github.com/lnvn/golang/pointer"
	"github.com/lnvn/golang/struct"
)

func main() {
	// Struct
	fmt.Println("----- Struct -----")
	mystruct.MyStruct()

	// Call pointer package
	fmt.Println("----- Pointer -----")
	pointer.Pointer()
	pointer.CheckNil()
}

