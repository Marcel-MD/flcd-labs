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

	fmt.Println("Enter nonterminal symbols separated by spaces: ")
	scanner.Scan()
	vn := strings.Split(scanner.Text(), " ")

	// fmt.Println("Enter terminal symbols separated by spaces: ")
	// scanner.Scan()
	// vt := strings.Split(scanner.Text(), " ")

	fmt.Println("Enter derivatives separated by spaces (S-aS S-bS): ")
	scanner.Scan()
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

	fmt.Println()
	printGraph(&graph, &vn)

	fmt.Println("\nGive me a string to check if it's accepted by FA: ")
	scanner.Scan()
	str := strings.Split(scanner.Text(), "")

	rsp := checkString(&str, &graph)

	if rsp {
		fmt.Println("Your string was accepted!")
	} else {
		fmt.Println("Your string was rejected!")
	}
}

func checkString(str *[]string, graph *map[string][]Pair) bool {
	pos := "S"

	for _, c := range *str {

		arr := (*graph)[pos]
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

func printGraph(graph *map[string][]Pair, vn *[]string) {
	for s, arr := range *graph {
		fmt.Print(convert(s, vn), ": [")
		for _, p := range arr {
			fmt.Printf("[%s: %s]", p.t, convert(p.n, vn))
		}
		fmt.Print("]\n")
	}
}

func convert(s string, vn *[]string) string {
	index := indexOf(s, vn)

	if index == -1 {
		return "q" + strconv.Itoa(len(*vn))
	}

	return "q" + strconv.Itoa(index)
}

func indexOf(element string, data *[]string) int {
	for i, v := range *data {
		if element == v {
			return i
		}
	}
	return -1
}
