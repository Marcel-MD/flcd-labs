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

	fmt.Print("Enter the number of final state: ")
	// 2
	scanner.Scan()
	end := scanner.Text()

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

	newStates := make([]string, 0)
	stateRows := make([]string, 0)
	// getting new states and the rows where they appeared
	newStates, stateRows = getNewStates(nfa, newStates, stateRows)

	// building deterministinc finite automata matrix graph
	dfa := make(map[string]map[string]string)

	// add q0 row as first row
	dfa["0"] = make(map[string]string)
	for k, v := range nfa["0"] {
		dfa["0"][k] = v
	}

	// add rows where new states appeared
	for _, row := range stateRows {
		if row != "0" {
			for k, v := range nfa[row] {
				if dfa[row] == nil {
					dfa[row] = make(map[string]string)
				}
				dfa[row][k] = v
			}
		}
	}

	// build new states rows
	for _, s := range newStates {
		dfa[s] = buildStateRow(nfa, s)
	}

	// build other rows if needed
	for _, v := range dfa {
		for _, vv := range v {
			if dfa[vv] == nil {
				dfa[vv] = buildStateRow(nfa, vv)
			}
		}
	}

	fmt.Print("\n")
	printFa(dfa)

	fmt.Println("\nGive me a string to check if it's accepted by DFA: ")
	// ababbb
	scanner.Scan()
	str := strings.Split(scanner.Text(), "")
	rsp := checkString(dfa, str, end)

	if rsp {
		fmt.Println("Your string was accepted!")
	} else {
		fmt.Println("Your string was rejected!")
	}
}

func checkString(dfa map[string]map[string]string, str []string, end string) bool {

	pos := "0"

	for _, c := range str {
		check := false

		for k, v := range dfa[pos] {
			if c == k {
				check = true
				pos = v
			}
		}

		if !check {
			return false
		}
	}

	if strings.Contains(pos, end) {
		return true
	}

	return false
}

func buildStateRow(nfa map[string]map[string]string, state string) map[string]string {
	row := make(map[string]string)
	states := strings.Split(state, "")

	for _, s := range states {
		for k, v := range nfa[s] {

			if row[k] == "" {
				row[k] = v
			} else if !strings.Contains(row[k], v) {
				v = v + row[k]
				row[k] = sortString(v)
			}
		}
	}

	return row
}

func getNewStates(fa map[string]map[string]string, newStates []string, stateRows []string) ([]string, []string) {
	for k, v := range fa {
		for _, vv := range v {
			if len(vv) > 1 {
				if indexOf(vv, newStates) == -1 {
					newStates = append(newStates, vv)
					stateRows = append(stateRows, k)
				}
			}
		}
	}

	return newStates, stateRows
}

func printFa(fa map[string]map[string]string) {
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
}

func sortString(str string) string {
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

// Because there is no indexOf in Go
func indexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1
}
