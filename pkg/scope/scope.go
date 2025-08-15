package scope

// file block scope
import "fmt"

// package block scope
var x = 42

func Scope() {
	fmt.Println("Scope package")
	fmt.Println(x)

	// blog scope
	y := 43
	println(y)

	// variable shadowing
	x := 100
	println(x)
}