package main

import(
	suffix_tree "github.com/astrieanna/substrings/lib"
	"fmt"
	)


var example1 []suffix_tree.DNA = []suffix_tree.DNA{
	suffix_tree.A, 
	suffix_tree.C,
	suffix_tree.C,
	suffix_tree.G,
	suffix_tree.T}
var example2 []suffix_tree.DNA = []suffix_tree.DNA{
	suffix_tree.C,
	suffix_tree.C,
	suffix_tree.A,
	suffix_tree.G,
	suffix_tree.T}


func main() {
	var t suffix_tree.Tree = suffix_tree.NewTree()
	t.Insert(example1)
	t.Insert(example2)
	fmt.Println(&t)
	ss := t.FindLongestSubstring()

	fmt.Println("Example1: ", example1)
	fmt.Println("Example2: ", example2)
	fmt.Printf("Longest Substring: %v\n", ss)
}
