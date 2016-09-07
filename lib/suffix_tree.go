package suffix_tree

import(
	"fmt"
	"log"
	"strings"
	)

type DNA int

const(
  A DNA = iota
  C
  G
  T
 )

func (d DNA) String() string {
	switch d {
	case A:
		return "A"
	case C:
		return "C"
	case G:
		return "G"
	case T:
		return "T"
	}
	return "Err"
}

type node struct {
	a *node
	c *node
	g *node
	t *node
	ids map[int]bool // strings with suffixes that end at this node
}

func newNode() *node {
	n := new(node)
	n.ids = make(map[int]bool)
	return n
}

func (n *node) insert(d DNA, id int) (*node) {
	n.ids[id] = true
	switch d {
	case A:
		if (n.a == nil) {
			n.a = newNode()
		}
		return n.a
	case C:
		if (n.c == nil) {
			n.c = newNode()
		}
		return n.c
	case G:
		if (n.g == nil) {
			n.g = newNode()
		}
		return n.g
	case T:
		if (n.t == nil) {
			n.t = newNode()
		}
		return n.t
	}
	log.Fatal("Not DNA: ", d)
	return nil
}

func (n *node) prettyPrint(depth int, max_id int) {
  none_missing := true
  for i := 0; i <= max_id; i++ {
  	if !n.ids[i] {
  		none_missing = false
  	}
  }
  if none_missing {
	  if n.a != nil {
	  	fmt.Println(strings.Repeat(".", depth), "A")
	  	n.a.prettyPrint(depth+1, max_id)
	  }
	  if n.c != nil {
	  	fmt.Println(strings.Repeat(".", depth), "C")
	  	n.c.prettyPrint(depth+1, max_id)
	  }
	  if n.g != nil {
	  	fmt.Println(strings.Repeat(".", depth), "G")
	  	n.g.prettyPrint(depth+1, max_id)
	  }
	  if n.t != nil {
	  	fmt.Println(strings.Repeat(".", depth), "T")
	  	n.t.prettyPrint(depth+1, max_id)
	  }
	}
}

type Tree struct {
	root node
	next_id int
}

func NewTree() Tree {
	return Tree{*newNode(), 0}
}

func (t *Tree) Insert(ds []DNA) {
	id := t.next_id
	t.next_id += 1
	end := len(ds)
	for i := range ds {
		var current_node *node = &(t.root);
		for _, v := range ds[i:end] {
			current_node = current_node.insert(v, id)
		}
	}
}

func (n node) findLongestSubstring(max_id int) (bool, []DNA) {
	for i := 0; i <= max_id; i++ {
		if n.ids[i] != true {
			return false, nil
		}
	}
	var best_result []DNA
	if n.a != nil {
		flag, result :=  n.a.findLongestSubstring(max_id)
		if flag {
			best_result = append([]DNA{A}, result...)
		}
	}
	if n.c != nil {
		flag, result := n.c.findLongestSubstring(max_id)
		if flag && len(result) + 1 >= len(best_result){
			best_result = append([]DNA{C}, result...)
		}	
	}
	if n.g != nil {
		flag, result := n.g.findLongestSubstring(max_id)
		if flag && len(result) + 1 >= len(best_result){
			best_result = append([]DNA{G}, result...)
		}	}
	if n.t != nil {
		flag, result := n.t.findLongestSubstring(max_id)
		if flag && len(result) + 1 >= len(best_result){
			best_result = append([]DNA{T}, result...)
		}
	}
	return true, best_result
}

func (n node) findCommonSubstrings(max_id int) (bool, [][]DNA) {
	for i := 0; i <= max_id; i++ {
		if n.ids[i] != true {
			return false, nil
		}
	}
	var my_results [][]DNA
	if n.a != nil {
		flag, their_results :=  n.a.findCommonSubstrings(max_id)
		if flag {
			for _, r := range their_results {
				my_results = append(my_results, append([]DNA{A}, r...))
			}
			my_results = append(my_results, []DNA{A})
		}
	}
	if n.c != nil {
		flag, their_results := n.c.findCommonSubstrings(max_id)
		if flag {
			for _, r := range their_results {
				my_results = append(my_results, append([]DNA{C}, r...))
			}
			my_results = append(my_results, []DNA{C})
		}
	}
	if n.g != nil {
		flag, their_results := n.g.findCommonSubstrings(max_id)
		if flag {
			for _, r := range their_results {
				my_results = append(my_results, append([]DNA{G}, r...))
			}
			my_results = append(my_results, []DNA{G})
		}
	}
	if n.t != nil {
		flag, their_results := n.t.findCommonSubstrings(max_id)
		if flag {
			for _, r := range their_results {
				my_results = append(my_results, append([]DNA{T}, r...))
			}
			my_results = append(my_results, []DNA{T})
		}
	}
	return true, my_results
}

func (t *Tree) FindCommonSubstrings() [][]DNA {
	_, results := t.root.findCommonSubstrings(t.next_id - 1)
	return results
}

func (t *Tree) PrettyPrint() {
	t.root.prettyPrint(0, t.next_id - 1)
}


