package embedding

import "fmt"

type describer interface {
	describe() string
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func Embed() {
	fmt.Println("=============================================================")
	fmt.Println("Embed")
	fmt.Println()

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	fmt.Println()
}
