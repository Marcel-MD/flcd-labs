package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	t string
	n string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter productions of rules separated by spaces (S-aS S-bS ...): ")
	// S-aS S-bS S-cD D-dD D-bF D-a F-bS F-a
	scanner.Scan()

	// productions of rules
	p := strings.Split(scanner.Text(), " ")

	graph := make(map[string][]Pair)

	for _, s := range p {
		arr := strings.Split(s, "")

		if len(arr) == 3 {
			arr = append(arr, "")
		}

		if graph[arr[0]] == nil {
			graph[arr[0]] = []Pair{{t: arr[2], n: arr[3]}}
		} else {
			graph[arr[0]] = append(graph[arr[0]], Pair{t: arr[2], n: arr[3]})
		}
	}

	// non-terminal symbols
	vn := make([]string, 0)
	dfs(graph, vn, "S")

	fmt.Println()
	printGraph(graph, vn)

	fmt.Println("\nGive me a string to check if it's accepted by FA: ")
	// acdbbca
	scanner.Scan()
	str := strings.Split(scanner.Text(), "")

	rsp := checkString(str, graph)

	if rsp {
		fmt.Println("Your string was accepted!")
	} else {
		fmt.Println("Your string was rejected!")
	}
}

// Check if string is accepted by FA
func checkString(str []string, graph map[string][]Pair) bool {
	pos := "S"

	for _, c := range str {

		arr := graph[pos]
		check := false

		for _, p := range arr {
			if p.t == c {
				pos = p.n
				check = true
				break
			}
		}

		if !check {
			return false
		}
	}

	if pos == "" {
		return true
	}

	return false
}

// Prints the graph
func printGraph(graph map[string][]Pair, vn []string) {
	for s, arr := range graph {
		fmt.Print(convertSymbol(s, vn), ": [")
		for _, p := range arr {
			fmt.Printf("[%s: %s]", p.t, convertSymbol(p.n, vn))
		}
		fmt.Print("]\n")
	}
}

// Depth First Search
func dfs(graph map[string][]Pair, vn []string, start string) {
	if start == "" {
		return
	}

	vn = append(vn, start)

	for _, p := range (graph)[start] {
		if indexOf(p.n, vn) != -1 {
			continue
		}
		dfs(graph, vn, p.n)
	}
}

// Converts nonterminal symbols to q0, q1...
func convertSymbol(s string, vn []string) string {
	index := indexOf(s, vn)

	if index == -1 {
		return "q" + strconv.Itoa(len(vn))
	}

	return "q" + strconv.Itoa(index)
}

// Because there is no indexOf in Go
func indexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1
}
