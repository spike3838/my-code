package main

import "fmt"

var zahl = 10
var z *int = &zahl

func main() {

	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	fmt.Println(p)  //this is a pointer! it just holds the address of i

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(z)
}
