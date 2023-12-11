package main

import (
	"fmt"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for c := 'a'; c <= 'j'; c++ {
		fmt.Printf("%c ", c)
	}
}

func task1() {
	fmt.Println("\n==============Task 1==============")

	go printNumbers()
	go printLetters()

	fmt.Println("\n===============Done===============")
}
