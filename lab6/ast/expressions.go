package ast

import (
	"bytes"
	"fmt"
	"lab/token"
	"strings"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (e *PrefixExpression) expression() {}
func (e *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Operator)
	out.WriteString(e.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (e *InfixExpression) expression() {}
func (e *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString(" " + e.Operator + " ")
	out.WriteString(e.Right.String())
	out.WriteString(")")
	return out.String()
}

type AssignExpression struct {
	Token      token.Token
	Identifier *Identifier
	Operator   string
	Value      Expression
}

func (as *AssignExpression) expression() {}
func (as *AssignExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(as.Identifier.String())
	out.WriteString(as.Operator)
	out.WriteString(as.Value.String())
	out.WriteString(")")
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (ie *IfExpression) expression() {}
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("IF ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {

		out.WriteString("ELSE ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type WhileExpression struct {
	Token     token.Token
	Condition Expression
	Block     *Block
}

func (we *WhileExpression) expression() {}
func (we *WhileExpression) String() string {
	var out bytes.Buffer
	out.WriteString("WHILE ")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Block.String())
	return out.String()
}

type FunctionCall struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (fc *FunctionCall) expression() {}
func (fc *FunctionCall) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range fc.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(fc.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

type Parameter struct {
	Token      token.Token
	Identifier *Identifier
}

func (p *Parameter) expression() {}
func (p *Parameter) String() string {
	return fmt.Sprintf("[%s %s]", p.Token.Type, p.Identifier.String())
}
