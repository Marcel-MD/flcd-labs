package parser

import (
	"fmt"
	"lab/ast"
	"lab/token"
)

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Value,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseAssignExpression(name ast.Expression) ast.Expression {
	expression := &ast.AssignExpression{Token: p.currentToken}

	if n, ok := name.(*ast.Identifier); ok {
		expression.Identifier = n
	} else {
		msg := fmt.Sprintf("expected assign token to be IDENT, got %s instead", name.String())
		p.errors = append(p.errors, msg)
	}

	p.nextToken()

	expression.Operator = "="
	expression.Value = p.parseExpression(LOWEST)
	return expression
}

func (p *Parser) parseFunctionCall(function ast.Expression) ast.Expression {
	exp := &ast.FunctionCall{Token: p.currentToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}
