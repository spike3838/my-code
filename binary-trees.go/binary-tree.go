package main

import(
	"golang.org/x/tour/tree"
	"fmt"
)
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){

	walk_rec(t, ch)
	close(ch)
}

func walk_rec(t *tree.Tree, ch chan int){

	if t != nil {
		//ch <- t.Value
		walk_rec(t.Left, ch)
		ch <- t.Value
		walk_rec(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {

	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Print(v)
}
}
