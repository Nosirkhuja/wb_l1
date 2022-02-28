package main

import (
	"fmt"
	"sync"
)

func wait_group(numbers []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))
	for i := 0; i < len(numbers); i++ {
		go func(id int) {
			fmt.Println(numbers[id] * numbers[id])
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func channel(numbers []int) {

	message := make(chan int, len(numbers))
	go func() {
		for i := 0; i < len(numbers); i++ {
			message <- numbers[i] * numbers[i]
		}
		close(message)
	}()

	for number2 := range message {
		fmt.Println(number2)
	}
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	fmt.Println("Wait group:")
	wait_group(numbers[:])
	fmt.Println("Channel:")
	channel(numbers[:])
}
