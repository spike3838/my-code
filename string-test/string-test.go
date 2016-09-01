package main

import(

	"fmt"
)

func main() {

	var test string
	test = "Hello"
	for i := 0; i < 5; i++ {
		fmt.Printf( "% x\n",test[i])
	}
}
