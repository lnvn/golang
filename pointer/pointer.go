// Example of pointer

package pointer

import "fmt"

type Config struct {
	Timeout int
	Server string
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

// we have parameter c of type *Config - apointer to Config
func SetDefaults(c *Config) {
	if c.Timeout == 0 {
		c.Timeout = 30
	}
	if c.Server == "" {
		c.Server = "default server"
	}
}