package ast

import (
	"fmt"
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

func (p *Program) String() string {
	str := "Program {\n"
	for _, s := range p.Statements {
		str = str + s.String()
	}
	str = str + "}\n"
	return str
}

type Block struct {
	Token      token.Token
	Statements []Statement
}

func (b *Block) String() string {
	str := "{\n"
	for _, s := range b.Statements {
		str = str + s.String()
	}
	str = str + "}\n"
	return str
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expression() {}
func (i *Identifier) String() string {
	str := fmt.Sprintf("[%s]", i.Value)
	return str
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (l *IntegerLiteral) expression() {}
func (l *IntegerLiteral) String() string {
	str := fmt.Sprintf("[%s: %d]", l.Token.Type, l.Value)
	return str
}

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (l *FloatLiteral) expression() {}
func (l *FloatLiteral) String() string {
	str := fmt.Sprintf("[%s: %f]", l.Token.Type, l.Value)
	return str
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (l *StringLiteral) expression() {}
func (l *StringLiteral) String() string {
	str := fmt.Sprintf("[%s: '%s']", l.Token.Type, l.Value)
	return str
}

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (l *BooleanLiteral) expression() {}
func (l *BooleanLiteral) String() string {
	str := fmt.Sprintf("[%s: %t]", l.Token.Type, l.Value)
	return str
}
