package mystruct

import "fmt"

// define a struct: use to create a collection of data types in one variable
type Info struct {
	fistname string
	lastname string
	age      int
}

// Embeded struct
type Person struct {
	Info
	Gender string
}

func MyStruct() {
	// Struct
	p1 := Info{
		fistname: "Go",
		lastname: "Lang",
		age:      16,
	}

	// Embeded struct example
	p2 := Person{
		Info: Info{
			fistname: "Embeded",
			lastname: "Struct",
			age:      1,
		},
		Gender: "male",
	}

	// Anonymous struct
	p3 := struct {
		Server string
		Port   int
	}{
		Server: "localhost",
		Port:   8080,
	}

	fmt.Printf("Type: %T, value: %v\n", p1, p1)
	fmt.Printf("Type: %T, value: %v\n", p2, p2)
	fmt.Printf("Type: %T, value: %v\n", p3, p3)
}
