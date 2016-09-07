package main

import(
	suffix_tree "github.com/astrieanna/substrings/lib"
	"fmt"
	"log"
	"os"
	"bufio"
	"sort"
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

func parseDNA(text string) []suffix_tree.DNA {
	output := []suffix_tree.DNA{}
	for _, c := range text {
		switch c {
		case 'A':
			output = append(output, suffix_tree.A)
		case 'C':
			output = append(output, suffix_tree.C)
		case 'G':
			output = append(output, suffix_tree.G)
		case 'T':
			output = append(output, suffix_tree.T)
		default:
			fmt.Println("Unknown: ", c)
		}
	}
	return output
}

type ByLength [][]suffix_tree.DNA

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Too few args; expected 1")
	}
	filename := os.Args[1]
	log.Print("Reading from ", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var t suffix_tree.Tree = suffix_tree.NewTree()

	scanner := bufio.NewScanner(file)
	is_label := true
	count := 0
	for scanner.Scan() {
		if !is_label{
			count += 1
			dna := parseDNA(scanner.Text())
			t.Insert(dna)

			if count % 10 == 0 {
				log.Print("count: ", count)
			}
		}
		is_label = !is_label
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print("total count: ", count)

	t.PrettyPrint()
	ss := t.FindCommonSubstrings()
	sort.Sort(ByLength(ss))
	for _, s := range ss {
		fmt.Printf("%v\n", s)
	}
}
