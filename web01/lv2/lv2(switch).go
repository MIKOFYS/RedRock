package main

import "fmt"

func sumFn1(x , y int) int{
	sum := x + y
	return sum

}

func subFn1(x , y int) int {
	sub := x - y
	return sub
}

func mulFn1(x , y int) int {
	mul := x * y
	return mul
}

func divFn1(x , y int )int {
	div := x / y
	return div
}

func main(){
	var a,b int
	var c string
	fmt.Println("INPUT:")
    scan: fmt.Scan(&a,&c,&b)
	fmt.Println("OUTPUT:")

	switch c {
	case "+": {
			fmt.Println(sumFn1(a , b))
			goto scan
		}
	case "-": {
			fmt.Println(subFn1(a , b))
			goto scan
		}
	case "*":{
			fmt.Println(mulFn1(a , b))
			goto scan
		}
	case "/":{
			if b == 0{
				fmt.Println("ERROR!")
			}
			if b != 0{
				fmt.Println(divFn1(a , b))
			}
			goto scan
		}
	}


}
