package switchCase

import "fmt"

func SwitchCase() {
	fmt.Println("=============================================================")
	fmt.Println("Switch")
	fmt.Println()

	var day = 5

	switch day {
	case 1:
		fmt.Println(1)
	case 5:
		fmt.Println(day)
	default:
		fmt.Println(0)
	}
	fmt.Println()
}
