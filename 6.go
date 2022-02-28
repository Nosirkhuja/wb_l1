package main

import "fmt"

// генератор возвращает канал, который производит числа 1, 2, 3…
// чтобы остановить основную горутину, необходимо отправить
// номер этому каналу
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 0 // остановка основной горутины
	//fmt.Println(<-numberм) // ошибка, больше никто не отправляет

}
