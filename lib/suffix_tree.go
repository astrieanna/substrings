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

func (n *node) String(depth int) string {
  var output string;
  for k := range n.ids {
  	output += fmt.Sprintf(" %d", k)
  }
  output += "\n"
  if n.a != nil {
  	output += strings.Repeat(".", depth) + "A" + n.a.String(depth+1)
  }
  if n.c != nil {
  	output += strings.Repeat(".", depth) + "C" + n.c.String(depth+1)
  }
  if n.g != nil {
  	output += strings.Repeat(".", depth) + "G" + n.g.String(depth+1)
  }
  if n.t != nil {
  	output += strings.Repeat(".", depth) + "T" + n.t.String(depth+1)
  }
  return output
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
	fmt.Println(n.ids[0], n.ids[1])
	for i := 0; i <= max_id; i++ {
		if n.ids[i] != true {
			fmt.Println("nil")
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
	fmt.Println(best_result)
	return true, best_result
}

func (t *Tree) FindLongestSubstring() []DNA {
	_, result := t.root.findLongestSubstring(t.next_id - 1)
	return result
}

func (t *Tree) String() string {
	return t.root.String(0)
}


