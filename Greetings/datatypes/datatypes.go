package datatypes

import "fmt"

func Types() {
	fmt.Println("=============================================================")
	fmt.Println("Data types")
	fmt.Println()

	var i = 0
	var s string = "s"

	//float32/64
	//bool

	fmt.Printf("%T \n", i)
	fmt.Printf("%T \n", s)
}

func ComplexTypes() {
	s := []int{1, 2}

	fmt.Printf("%T \n", s)
	fmt.Println(len(s), cap(s))

	m := make(map[int]string, 0)

	m[1] = ""
	m[2] = "s"

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	var refI i = s
	fmt.Println(refI, &refI, &m)

	fmt.Println()
}

type i interface{}
