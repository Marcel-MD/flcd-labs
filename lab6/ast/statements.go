package ast

import (
	"fmt"
	"lab/token"
)

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (s *ExpressionStatement) statement() {}
func (s *ExpressionStatement) String() string {
	return s.Value.String() + "\n"
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (s *ReturnStatement) statement() {}
func (s *ReturnStatement) String() string {
	str := fmt.Sprintf("{%s: %s}\n", s.Token.Type, s.Value.String())
	return str
}

type VariableDeclarationStatement struct {
	Token      token.Token
	Identifier *Identifier
	Value      Expression
}

func (s *VariableDeclarationStatement) statement() {}
func (s *VariableDeclarationStatement) String() string {
	str := fmt.Sprintf("{%s %s = %s}\n", s.Token.Type, s.Identifier.String(), s.Value.String())
	return str
}

type FunctionDeclarationStatement struct {
	Token      token.Token
	Identifier *Identifier
	Parameters []*Parameter
	Body       *Block
}

func (s *FunctionDeclarationStatement) statement() {}
func (s *FunctionDeclarationStatement) String() string {
	str := fmt.Sprintf("Function {%s %s (", s.Token.Type, s.Identifier.String())
	for _, p := range s.Parameters {
		str = str + p.String()
	}
	str = str + fmt.Sprintf(") %s}\n", s.Body.String())
	return str
}
