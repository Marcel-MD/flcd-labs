package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type FiniteAutomataRow map[string]string
type FiniteAutomata map[string]FiniteAutomataRow

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter transitions separated by spaces (0-a-0 0-b-1 ...): ")
	// 0-a-0 1-b-1 1-b-2 0-b-1 1-a-0 2-b-1
	scanner.Scan()
	transitions := strings.Split(scanner.Text(), " ")

	nfa := buildNfa(transitions)
	printFa(nfa)

	dfa := convertNfaToDfa(nfa)
	printFa(dfa)
}

func buildNfa(transitions []string) FiniteAutomata {
	nfa := make(FiniteAutomata)
	for _, t := range transitions {
		arr := strings.Split(t, "-")

		if nfa[arr[0]] == nil {
			nfa[arr[0]] = FiniteAutomataRow{arr[1]: arr[2]}
		} else {
			nfa[arr[0]][arr[1]] = prepareStateStr(nfa[arr[0]][arr[1]] + arr[2])
		}
	}
	return nfa
}

func convertNfaToDfa(nfa FiniteAutomata) FiniteAutomata {
	dfa := make(FiniteAutomata)

	// copy first row from nfa
	dfa["0"] = make(FiniteAutomataRow)
	for k, v := range nfa["0"] {
		dfa["0"][k] = v
	}
	printFaRow(dfa, "0")

	recursiveBuildDfa(&dfa, nfa)

	return dfa
}

func recursiveBuildDfa(dfa *FiniteAutomata, nfa FiniteAutomata) {
	for _, v := range *dfa {
		for _, vv := range v {
			if (*dfa)[vv] == nil {
				(*dfa)[vv] = buildStateRow(nfa, vv)
				printFaRow(*dfa, vv)
				recursiveBuildDfa(dfa, nfa)
			}
		}
	}
}

func buildStateRow(nfa FiniteAutomata, state string) FiniteAutomataRow {
	row := make(FiniteAutomataRow)
	states := strings.Split(state, "")

	for _, s := range states {
		for k, v := range nfa[s] {

			if row[k] == "" {
				row[k] = v
			} else if !strings.Contains(row[k], v) {
				v = v + row[k]
				row[k] = prepareStateStr(v)
			}
		}
	}

	return row
}

func prepareStateStr(str string) string {
	s := strings.Split(str, "")
	s = removeDuplicateStr(s)
	sort.Strings(s)
	return strings.Join(s, "")
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func printFa(fa FiniteAutomata) {
	fmt.Print("\n")

	for k, v := range fa {
		for _, c := range k {
			fmt.Printf("q%c", c)
		}
		fmt.Print(": [")

		for kk, vv := range v {
			fmt.Printf("[%s: ", kk)
			for _, c := range vv {
				fmt.Printf("q%c", c)
			}
			fmt.Print("]")
		}

		fmt.Print("]\n")
	}

	fmt.Print("\n")
}

func printFaRow(fa FiniteAutomata, row string) {
	for _, c := range row {
		fmt.Printf("q%c", c)
	}
	fmt.Print(":")

	for k, v := range fa[row] {
		fmt.Printf(" (%s -> ", k)
		for _, c := range v {
			fmt.Printf("q%c", c)
		}
		fmt.Print(")")
	}

	fmt.Print("\n")
}
