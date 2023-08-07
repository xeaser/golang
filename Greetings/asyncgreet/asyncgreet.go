package asyncgreet

import (
	"fmt"
	"sync"
)

func Greet() {
	fmt.Println("=============================================================")
	fmt.Println("Greet")
	fmt.Println()

	var wg sync.WaitGroup
	msgChan := make(chan string, 2)

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go asyncGreet(&wg, msgChan, fmt.Sprintf("Hello %d!", i))
	}

	wg.Wait()
	close(msgChan)

	for i := range msgChan {
		fmt.Println(i)
	}

	fmt.Println()
}

func asyncGreet(wg *sync.WaitGroup, c chan string, msg string) {
	defer wg.Done()
	c <- msg
}
