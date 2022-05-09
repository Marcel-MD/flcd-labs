package ast

import (
	"bytes"
	"lab/token"
)

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (s *ExpressionStatement) statement() {}
func (s *ExpressionStatement) String() string {
	return s.Value.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (s *ReturnStatement) statement() {}
func (s *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(s.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + s.Value.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type VariableDeclarationStatement struct {
	Token      token.Token
	Identifier *Identifier
	Value      Expression
}

func (s *VariableDeclarationStatement) statement() {}
func (s *VariableDeclarationStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: DECLARATION " + string(s.Token.Type) + "\n")
	out.WriteString(printTab() + "identifier: " + s.Identifier.String())
	out.WriteString(printTab() + "value: " + s.Value.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type FunctionDeclarationStatement struct {
	Token      token.Token
	Identifier *Identifier
	Parameters []*Parameter
	Body       *Block
}

func (s *FunctionDeclarationStatement) statement() {}
func (s *FunctionDeclarationStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: FUNCTION " + string(s.Token.Type) + "\n")
	out.WriteString(printTab() + "identifier: " + s.Identifier.String())
	out.WriteString(printTab() + "parameters: [\n")
	tab++
	for _, p := range s.Parameters {
		out.WriteString(printTab() + p.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	out.WriteString(printTab() + "body: " + s.Body.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}
