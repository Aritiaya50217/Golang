package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("Before :", s)

	const n = 5000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	// sine
	fmt.Println(math.Sin(n))
}
