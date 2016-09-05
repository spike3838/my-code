package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (

	v1 = Vertex{1, 2} //has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // Y:0, X:0 implicit
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {

	fmt.Println(v1, p, v2, v3)
}
