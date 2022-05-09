package main

import (
	"fmt"
	"io/ioutil"
	"lab/lexer"
	"lab/parser"
	"log"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		path := os.Args[1]
		scanFile(path)
	} else {
		scanFile("test.txt")
	}
}

func scanFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	code := string(content)
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		return
	}

	fmt.Print(program)
}

func printParserErrors(errors []string) {
	fmt.Print(" parser errors:\n")
	for _, msg := range errors {
		fmt.Print("\t" + msg + "\n")
	}
}
