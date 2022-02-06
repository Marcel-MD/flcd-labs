package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter transitions separated by spaces (0-a-0 0-b-1 ...): ")
	// 0-a-0 1-b-1 1-b-2 0-b-1 1-a-0 2-b-1
	scanner.Scan()
	transitions := strings.Split(scanner.Text(), " ")

	// building nondeterministinc finite automata matrix graph
	nfa := make(map[string]map[string]string)
	for _, t := range transitions {
		arr := strings.Split(t, "-")

		if nfa[arr[0]] == nil {
			nfa[arr[0]] = map[string]string{arr[1]: arr[2]}
		} else {
			nfa[arr[0]][arr[1]] = sortString(nfa[arr[0]][arr[1]] + arr[2])
		}
	}

	printFa(nfa)
}

func printFa(fa map[string]map[string]string) {
	fmt.Print("\n")

	for k, v := range fa {
		fmt.Printf("q%s: [", k)

		for kk, vv := range v {
			fmt.Printf("[%s: ", kk)
			for _, c := range vv {
				fmt.Printf("q%c", c)
			}
			fmt.Print("]")
		}

		fmt.Print("]\n")
	}
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
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
