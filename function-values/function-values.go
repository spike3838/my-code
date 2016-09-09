package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64{

	var x, y float64 = 4.44, 5.55
	return fn(x, y)
}

func main(){

	a := func(x , y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(a(3, 4))

	fmt.Println(compute(a))
	fmt.Println(compute(math.Pow))
	fmt.Println("(╯°□°）╯︵ ┻━┻")
}
