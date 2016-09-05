package main

import "fmt"

type Vertex struct {

	X int
	Y int
}

func main(){

	p := &Vertex{3, 5}
//	p := &v
	p.X = 1e9
	p.Y = 17
	fmt.Println(p.X)
	fmt.Println(p.Y)
}
