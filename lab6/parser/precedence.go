package parser

import (
	"lab/token"
)

const (
	_ int = iota
	LOWEST
	LOGIC
	ASSIGN
	EQUALS
	COMPARE
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{
	token.OR:       LOGIC,
	token.AND:      LOGIC,
	token.ASSIGN:   ASSIGN,
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LESS:     COMPARE,
	token.GREAT:    COMPARE,
	token.GREAT_EQ: COMPARE,
	token.LESS_EQ:  COMPARE,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}
