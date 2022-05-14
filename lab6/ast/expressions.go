package ast

import (
	"bytes"
	"lab/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (e *PrefixExpression) expression() {}
func (e *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: PREFIX " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "operator: " + e.Operator + "\n")
	out.WriteString(printTab() + "right: " + e.Right.String())
	tab--
	out.WriteString(printTab() + "}\n")
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
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: INFIX " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "left: " + e.Left.String())
	out.WriteString(printTab() + "operator: " + e.Operator + "\n")
	out.WriteString(printTab() + "right: " + e.Right.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type AssignExpression struct {
	Token      token.Token
	Identifier *Identifier
	Operator   string
	Value      Expression
}

func (e *AssignExpression) expression() {}
func (e *AssignExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "identifier: " + e.Identifier.String())
	out.WriteString(printTab() + "operator: " + e.Operator + "\n")
	out.WriteString(printTab() + "value: " + e.Value.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (e *IfExpression) expression() {}
func (e *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "condition: " + e.Condition.String())
	out.WriteString(printTab() + "consequence: " + e.Consequence.String())
	if e.Alternative != nil {
		out.WriteString(printTab() + "alternative: " + e.Alternative.String())
	}
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type WhileExpression struct {
	Token     token.Token
	Condition Expression
	Body      *Block
}

func (e *WhileExpression) expression() {}
func (e *WhileExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "condition: " + e.Condition.String())
	out.WriteString(printTab() + "body: " + e.Body.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type ForExpression struct {
	Token       token.Token
	Declaration *VariableDeclarationStatement
	Condition   Expression
	Increment   Expression
	Body        *Block
}

func (e *ForExpression) expression() {}
func (e *ForExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "declaration: " + e.Declaration.String())
	out.WriteString(printTab() + "condition: " + e.Condition.String())
	out.WriteString(printTab() + "increment: " + e.Increment.String())
	out.WriteString(printTab() + "body: " + e.Body.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type FunctionCall struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (e *FunctionCall) expression() {}
func (e *FunctionCall) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: CALL\n")
	out.WriteString(printTab() + "function: " + e.Function.String())
	out.WriteString(printTab() + "arguments: [\n")
	tab++
	for _, a := range e.Arguments {
		out.WriteString(printTab() + a.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}
