package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	q := 0.0
	z := 1.0

	for count := 0 ; count < 100000; count++ {
	z = z-(((z*z) - x) / (2*z))
	}
	q = z

	fmt.Println(Diff(q , z))
	return z

}

func Diff(x , z float64) float64 {

	if x < z {
	return z -x
	} else if x > z {
	return x -z
	} else {
	return 0.0
	}
}

func main(){
	fmt.Println(Sqrt(11))
}
