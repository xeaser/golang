package slices

import "fmt"

func SliceOptn() {
	fmt.Println("=============================================================")
	fmt.Println("Slices")
	fmt.Println()

	a := [3]int{1, 2}
	fmt.Println(len(a), cap(a), a)

	s := []int{1, 2}
	fmt.Println(len(s), cap(s), s)

	s = s[len(s)-1:]
	fmt.Println(len(s), cap(s), s)

	fmt.Println()
}
