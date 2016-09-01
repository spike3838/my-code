package main

import (

	"fmt"
	"math/rand"

)

func main()  {

	rand.Seed(65535)
	fmt.Println("Random number: ", rand.Intn(65535))

}
