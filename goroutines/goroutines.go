package main

import (

	"fmt"
	"time"
)

func say(s string) {

	for i := 0; i < 5; i++{

		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	say("World")  //if you don't wait for the routine, main exits before it could do smth
	go say("Hello")
	//say("World")
}
