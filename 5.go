package main

import (
	"fmt"
	"os"
	"time"
)

func Read(ch chan string) {
	counter := 1
	for data := range ch {
		fmt.Println("read data: ", data, " ", counter, " time(s)")
		counter++
	}
	fmt.Println("channel closed")
}

func main() {

	var t float64

	fmt.Println("Enter t:")

	fmt.Fscan(os.Stdin, &t)

	ch := make(chan string)

	go Read(ch)

	var duration = time.Millisecond * time.Duration(t)
	
	timer := time.NewTimer(duration)

	for {
		select {
		case <-timer.C:
			close(ch)
			fmt.Println("time's out")
			time.Sleep(time.Millisecond * 100)
			return
		default:
			ch <- "Hello World!"
		}
		time.Sleep(time.Millisecond * 100)
	}
}
