package main

import "fmt"

func main() {
	// make(map[key-type]value-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("map :", m)

	// get value
	v1 := m["k1"]
	fmt.Println("value :", v1)

	// length
	fmt.Println("len:", len(m))

	// delete
	delete(m, "k2")
	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
