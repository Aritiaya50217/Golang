package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// returned values
	_, c := vals()
	fmt.Println("value : ", c)
}
