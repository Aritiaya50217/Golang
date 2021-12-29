package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// set a key/value
	os.Setenv("FOO", "1")
	// get a value for a key
	fmt.Println("FOO:", os.Getenv("FOO"))

	// return " " เพราะไม่มีค่า key
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// Use os.Environ to list all key/value pairs in the environment
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
