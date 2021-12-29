package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// Type
	fmt.Printf("Type :%T\n", f)

	// 0 means infer the base from the string
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// Type
	fmt.Printf("Type :%T\n", i)

	// hex-formatted (เลขฐาน 16)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Type
	fmt.Printf("Type :%T\n", u)

	// Atoi is a convenience function for basic base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
