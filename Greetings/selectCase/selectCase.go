package selectCase

import (
	"fmt"
	"os"
	"strings"
)

func fibonnaci(c <-chan int, done <-chan int) {
	fmt.Println("=============================================================")
	fmt.Println("select")
	fmt.Println()

	x, y := 0, 1
	for {
		select {
		case <-c:
			fmt.Println(x)
			x, y = y, x+y
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func GetFibonnaci() {
	c := make(chan int)
	done := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		done <- 0
	}()

	fibonnaci(c, done)
}

func GetGoPath() {
	fmt.Println("=============================================================")
	fmt.Println("select")
	fmt.Println()
	envs := os.Environ()
	str := make(chan string, len(envs))
	done := make(chan int)
	for _, s := range envs {
		str <- s
	}

	check(str, done)

	select {
	case <-str:
		fmt.Println(<-str)
	case <-done:
		fmt.Println("done")
	}
}

func check(s chan string, done chan int) {
	pathKey := "GOPATH"
	str := <-s
	if strings.Contains(str, pathKey) {
		fmt.Println(str)
		d := strings.Split(str, pathKey+"=")
		s <- d[1]
		done <- 0
	} else {
		fmt.Println(str)
	}
}
