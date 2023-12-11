package main

import (
	"fmt"
	"sync"
)

func printNumbers2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters2(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := 'a'; c <= 'j'; c++ {
		fmt.Printf("%c ", c)
	}
}

func task2() {
	fmt.Println("\nTask 2")

	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers2(&wg)
	go printLetters2(&wg)

	wg.Wait()

	fmt.Println("\nDone")
}
