package main

import (
	"fmt"
)

func sendNumbers6(evenCh, oddCh chan int, errCh chan error) {
	for i := 1; i <= 25; i++ {
		if i > 20 {
			errCh <- fmt.Errorf("error: Number %d is greater than 20", i)
		} else if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
	close(errCh)
}

func task6() {
	fmt.Println("\nTask 6")

	evenCh := make(chan int)
	oddCh := make(chan int)
	errCh := make(chan error)

	go sendNumbers6(evenCh, oddCh, errCh)

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
		case err, ok := <-errCh:
			if !ok {
				fmt.Println("Error channel closed.")
				errCh = nil
			} else {
				fmt.Println(err)
			}
		}

		if evenCh == nil && oddCh == nil && errCh == nil {
			break
		}
	}
}
