package main

import "fmt"

var step float64 = 0

func Sqrt(x float64) (float64, int){

	z := 1.0
	count := 0

	for step - z != 0{

	count++
	if count > 100{ return z, count}

		z = newtons(x, z)
		step = newtons(x, z)
	}
return z, count

}


func newtons(x, z float64) float64{

	return z - (z * z - x) / (2 * z)
}

func main(){

	fmt.Println(Sqrt(7))
}
