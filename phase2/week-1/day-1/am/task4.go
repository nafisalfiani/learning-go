package main

import (
	"fmt"
	"sync"
)

func produce4(ch chan int, wg *sync.WaitGroup) {
	defer close(ch)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	wg.Done()
}

func consume4(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("%d\n", num)
	}
}

func task4() {
	fmt.Println("\nTask 4")

	var wg sync.WaitGroup

	ch := make(chan int, 5)
	wg.Add(2)

	go produce4(ch, &wg)
	go consume4(ch, &wg)

	wg.Wait()

	fmt.Println("\nDone")
}
