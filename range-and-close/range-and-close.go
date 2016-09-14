package main

import (

	"fmt"
)

func fibonacci(n int, c chan int) {

	x, y := 0, 1
	for i := 0; i < n; i++{

		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	c := make(chan int, 50)
	go fibonacci(cap(c), c)
	for i := range c {

		fmt.Println(i)
	}
}

/* channels dont need to be closed, like files, they only do need to be closed to tell the recaiver that no more files are going to come*/
