package main

import (
	"fmt"
)

func main() {

	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			if i == 9 {
				fmt.Println(i)
				over <- true
			}
		}(i)
	}
	<-over
	fmt.Println("over!!!")
}