package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	// use []byte(s) to coerce it to bytes
	h.Write([]byte(s))
	fmt.Println(h)

	bs := h.Sum(nil)

	// Use the %x format verb to convert a hash results to a hex string
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
