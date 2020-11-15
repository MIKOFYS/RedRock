package main

import ("fmt"
	"math"
)

var i int
var min float64

func minimum (x [9]float64) (mini float64) {
	m := x[0]
	for i = 0 ; i < 9 ; i++ {
		if math.Abs(m) <= math.Abs(x[i]) {
			m = math.Abs(m)
		} else
		{
			m = x[i]
		}
		mini = m
	}

	return

}

func main() {
	var array = [9]float64{-88.0, -25.0, -12.0,  -8.0, -0.5, 1.0, 5.0, 6.0, 77.0}
	min = minimum(array)
	fmt.Println(min)
}