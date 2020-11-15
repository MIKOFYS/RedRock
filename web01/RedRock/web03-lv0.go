package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
)

var ch1 = make(chan int , 1)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch1 <- res
	myres[n] = <- ch1

}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}