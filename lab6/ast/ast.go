package ast

import (
	"bytes"
	"lab/token"
)

type Node interface {
	String() string
}

type Expression interface {
	Node
	expression()
}

type Statement interface {
	Node
	statement()
}

type Program struct {
	Statements []Statement
}

var tab int = 0

func printTab() string {
	var out bytes.Buffer
	for i := 0; i < tab; i++ {
		out.WriteString("  ")
	}
	return out.String()
}

func (p *Program) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: PROGRAM\n")
	out.WriteString(printTab() + "statements: [\n")
	tab++
	for _, s := range p.Statements {
		out.WriteString(printTab() + s.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString("}\n")
	return out.String()
}

type Block struct {
	Token      token.Token
	Statements []Statement
}

func (b *Block) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: BLOCK\n")
	out.WriteString(printTab() + "statements: [\n")
	tab++
	for _, s := range b.Statements {
		out.WriteString(printTab() + s.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expression() {}
func (i *Identifier) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(i.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + i.Token.Value + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (l *IntegerLiteral) expression() {}
func (l *IntegerLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + l.Token.Value + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (l *FloatLiteral) expression() {}
func (l *FloatLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + l.Token.Value + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (l *StringLiteral) expression() {}
func (l *StringLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + l.Token.Value + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (l *BooleanLiteral) expression() {}
func (l *BooleanLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + l.Token.Value + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}
