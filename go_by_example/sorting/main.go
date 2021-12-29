package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	// เรียง string
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// เรียง int
	ints := []int{1, 2, 9, 4, 0}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
