package main

import "golang.org/x/tour/tree"
import "fmt"
import "reflect"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	queue := make([]*tree.Tree, 0)
	current := t
	queue = append(queue, current)
	for {
		if len(queue) < 1 { break }
		current = queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = make([]*tree.Tree, 0)
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		} 
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
		ch <- current.Value
	}
	//fmt.Println("Closing channel")
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1d, t2d := make(map[string]int), make(map[string]int)
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	chan2map := func(ch chan int, d map[string]int){
		for i := range ch {
			key := i
			if _, ok := d[string(key)]; ok {
				d[string(key)] += 1	
			} else {
				d[string(key)] = 1	
			}
		}
	}
	chan2map(ch1, t1d)
	chan2map(ch2, t2d)
	return reflect.DeepEqual(t1d, t2d)
}

func main() {
	tree1 := tree.New(10)
	//fmt.Println(tree1)
	tree2 := tree.New(10)
	//fmt.Println(tree1, tree2)
	//ch1 := make(chan int)
	//ch1, ch2 := make(chan int), make(chan int)
	//go Walk(tree1, ch1)
	//for i := range ch1 {
	//	fmt.Println(i)	
	//}
	//go Walk(tree2, ch2)
	fmt.Println(Same(tree1, tree2))
}
