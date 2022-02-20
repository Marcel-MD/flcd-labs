package main

import (
	"fmt"
	"io/ioutil"
	"lab/lexer"
	"lab/token"
	"log"
	"os"
)

func main() {

	var tokens []token.Token

	if len(os.Args) > 1 {
		path := os.Args[1]
		tokens = scanFile(path)
	} else {
		tokens = scanFile("test.txt")
	}

	fmt.Println(tokens)
}

func scanFile(path string) []token.Token {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	code := string(content)
	l := lexer.New(code)
	return l.GetTokens()
}
