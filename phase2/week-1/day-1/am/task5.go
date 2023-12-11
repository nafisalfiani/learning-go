package main

import (
	"fmt"
)

func sendNumbers(evenCh, oddCh chan int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func task5() {
	fmt.Println("\nTask 5")

	evenCh := make(chan int)
	oddCh := make(chan int)

	go sendNumbers(evenCh, oddCh)

	for {
		select {
		case even, ok := <-evenCh:
			if !ok {
				fmt.Println("Even channel closed.")
				evenCh = nil
			} else {
				fmt.Printf("Received even number: %d\n", even)
			}
		case odd, ok := <-oddCh:
			if !ok {
				fmt.Println("Odd channel closed.")
				oddCh = nil
			} else {
				fmt.Printf("Received odd number: %d\n", odd)
			}
		}

		if evenCh == nil && oddCh == nil {
			break
		}
	}
}
