package main

import (
	"fmt"
)

func Receiver(v interface{}){
	_,ok1 := v.(string)
	_,ok2 := v.(int)
	_,ok3 := v.(bool)
	if ok1{
		fmt.Println("这个是string")
	}
	if ok2{
		fmt.Println("这个是int")
	}
	if ok3{
		fmt.Println("这个是bool")
	}
}


func main() {
	var x interface{}
	fmt.Println("请输入：")
	x = 5
	fmt.Scanln(&x)
	Receiver(x)
}

