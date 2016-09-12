package main

import (
	"fmt"
	"math"
)

type myfloat float64

func (f myfloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f := myfloat(-math.Sqrt(2))
	fmt.Println(f.Abs())
}
