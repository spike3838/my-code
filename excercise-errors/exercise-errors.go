package main

import "fmt"

var step float64 = 0

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	err := float64(e)
	return fmt.Sprint("cannot Sqrt negative Number: ", err)
}

func Sqrt(x float64) (float64, error){

        if x < 0 {

		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
        count := 0

        for step - z != 0{

                count++
                if count > 100{ return z, nil}

                z = newtons(x, z)
                step = newtons(x, z)
        }
return z, nil

}


func newtons(x, z float64) float64{

        return z - (z * z - x) / (2 * z)
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
