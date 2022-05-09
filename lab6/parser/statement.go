package parser

import (
	"lab/ast"
	"lab/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseDeclarationStatement() ast.Statement {
	if p.peekPeekTokenIs(token.ASSIGN) {
		return p.parseVariableDeclaration()
	}

	return p.parseFunctionDeclaration()
}

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclarationStatement {
	stmt := &ast.VariableDeclarationStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Identifier = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Value}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseFunctionDeclaration() *ast.FunctionDeclarationStatement {
	stmt := &ast.FunctionDeclarationStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Identifier = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Value}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	stmt.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlock()

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.Parameter {
	parameters := []*ast.Parameter{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return parameters
	}

	p.nextToken()
	param := &ast.Parameter{Token: p.currentToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	param.Identifier = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Value}

	parameters = append(parameters, param)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		param := &ast.Parameter{Token: p.currentToken}

		if !p.expectPeek(token.IDENTIFIER) {
			return nil
		}

		param.Identifier = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Value}

		parameters = append(parameters, param)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return parameters
}
