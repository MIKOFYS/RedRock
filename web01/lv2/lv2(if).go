package main

import "fmt"

func sumFn(x , y int) int{
	sum := x + y
	return sum

}

func subFn(x , y int) int {
	sub := x - y
	return sub
}

func mulFn(x , y int) int {
	mul := x * y
	return mul
}

func divFn(x , y int )int {
	div := x / y
	return div
}

func main(){
	var a,b int
	var c string
	fmt.Println("INPUT:")
	scan: fmt.Scan(&a,&c,&b)
	fmt.Println("OUTPUT:")

	if c == "+"{
		fmt.Println(sumFn(a , b))
		goto scan
	}else
	if c== "-"{
		fmt.Println(subFn(a , b))
		goto scan
	}else
	if c == "*"{
		fmt.Println(mulFn(a , b))
		goto scan
	}else
	if c == "/"{
		if b == 0{
			fmt.Println("ERROR!")
		}
		if b != 0{
			fmt.Println(divFn(a , b))
		}
		goto scan
	}

}
