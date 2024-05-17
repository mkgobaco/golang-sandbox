package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left == nil && t.Right == nil {
		//wg.Add(1)
		ch <- t.Value
		return
	}

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	//wg.Add(1)
	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the treesf
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	/*
		var s1 string
		for i := range ch1 {
			s1 += strconv.Itoa(i)
		}

		var s2 string
		for i := range ch2 {
			s2 += strconv.Itoa(i)
		}
		return strings.Compare(s1, s2) == 0
	*/

	for i := range ch1 {
		j := <-ch2
		if i != j {
			return false
		}
	}

	return true

}

func walkChannelRange() {
	fmt.Println("walkChannelRange start")
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("walkChannelRange finish")
}

func walkChannelSelect() {
	fmt.Println("walkChannelSelect start")
	ch := make(chan int)
	end := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		end <- 1
	}()

	for {
		select {
			case i := <- ch:
				fmt.Println(i)
			case <- end:
				fmt.Println("walkChannelSelect finish")	
				return
		}
	}

	
}

func main() {

	walkChannelRange()
	walkChannelSelect()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
