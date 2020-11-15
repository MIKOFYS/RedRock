package main

import (
	"fmt"
)

func Receiver(v interface{}) {
	switch v.(type)  {
	case int: fmt.Println("这个是int")
	break
	case float64: fmt.Println("这个是float64")
	break
	case bool: fmt.Println("这个是bool")
	break
	case string: fmt.Println("这个是string")
	break
	}
}

	func main() {
		var x interface{}
		_,_ = fmt.Scanf("%s",x)
		Receiver(x)
	}
