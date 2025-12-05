package main

import (
	"fmt"

	"github.com/lnvn/golang/demo_http"
	"github.com/lnvn/golang/multipleretrun"
	"github.com/lnvn/golang/pointer"
	"github.com/lnvn/golang/struct"
	"github.com/lnvn/golang/variadic"
)

func main() {
	// Struct
	fmt.Println("----- Struct -----")
	mystruct.MyStruct()

	// Call pointer package
	fmt.Println("----- Pointer -----")
	pointer.Pointer()
	pointer.CheckNil()

	fmt.Println("----- Variadic Input -----")
	the_variadic := variadic.VariadicInput(3, 2, 3, 4, 5)
	fmt.Println(the_variadic)

	fmt.Println("----- Variadic Input -----")
	m, n, err := multiplereturn.MultipleReturn(24, 7)
	fmt.Printf("%v, %v, %v", m, n, err)

	fmt.Println("----- Variadic Input -----")
	fmt.Println("----- HTTP Multiplexer -----")
	demo_http.MyMux()
}
