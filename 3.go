package main

import "fmt"

func summa1(numbers []int) int {
	var curr int
	channel := make(chan int, len(numbers))

	go func() {
		for i := 0; i < len(numbers); i++ {
			go func(num int) {
				channel <- num * num
			}(numbers[i])
		}
	}()
	for i := 0; i < len(numbers); i++ {
		curr += <-channel
	}
	close(channel)
	return curr
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println("Сумма квадратов", numbers, ": ", summa1(numbers))
}
