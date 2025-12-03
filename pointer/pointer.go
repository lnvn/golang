package pointer

import "fmt"

type Config struct {
	Timeout int
	Server  string
}

func Pointer() {
	// Create a new and zero-value Config struct
	// & return memory address of Config struct
	configPointer := &Config{}
	fmt.Println("Config Timeout: ", configPointer.Timeout)
	fmt.Println("Config Server: ", configPointer.Server)
	SetDefaults(configPointer)
	fmt.Println()
	fmt.Println("Final Timeout: ", configPointer.Timeout)
	fmt.Println("Final Server: ", configPointer.Server)
}

// we have parameter c of type *Config - a pointer to Config
func SetDefaults(c *Config) {
	if c.Timeout == 0 {
		c.Timeout = 30
	}
	if c.Server == "" {
		c.Server = "default server"
	}
}

func CheckNil() {
	// p is initialized to its zero value, which is nil
	var p *int

	if p == nil {
		fmt.Println("p is nil, it does not point to any valid memory address")
		fmt.Println("p: ", p)
	}

	x := 2711
	// p hold the memory address of x
	p = &x
	// *p = 100
	fmt.Println("x memory value: ", p)

	if p != nil {
		fmt.Println("p is not nil, it point to a value")
		fmt.Printf("p: %v\n\n", p)
	}
}
