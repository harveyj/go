package main

import "code.google.com/p/go-tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	lchan := make(chan int)
	if t.Left != nil {
		go Walk(t.Left, lchan)
		for v := range lchan {
			ch <- v
		}
	}

	ch <- t.Value

	rchan := make(chan int)
	if t.Right != nil {
		go Walk(t.Right, rchan)
		for v := range rchan {
			ch <- v
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1chan := make(chan int)
	t2chan := make(chan int)
	go Walk(t1, t1chan)
	go Walk(t2, t2chan)
	for {
		v1, ok1 := <-t1chan
		v2, ok2 := <-t2chan

		if ok1 != ok2 {
			return false
		} else if v1 != v2 {
			return false
		} else if ok1 && ok2 {
			return true
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		print(<-ch)
	}
	print(Same(tree.New(1), tree.New(1)))
	print(Same(tree.New(1), tree.New(2)))
}
