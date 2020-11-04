package main

import (
	"fmt"
)

var max = 100

func OddNumber(ch chan int) {
	for i := 1; i <= max; i++ {
		ch <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}

func EvenNumber(ch chan int) {
	for i := 1; i <= max; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}

func main() {
	ch := make(chan int)
	go OddNumber(ch)
	go EvenNumber(ch)
}
