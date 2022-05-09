package parser

import (
	"fmt"
	"lab/ast"
	"lab/token"
	"strconv"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}
	value, err := strconv.ParseInt(p.currentToken.Value, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.currentToken}
	value, err := strconv.ParseFloat(p.currentToken.Value, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.currentToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.currentToken, Value: p.currentTokenIs(token.TRUE)}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.currentToken, Value: p.currentToken.Value}
}
