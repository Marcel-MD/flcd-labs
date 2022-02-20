package lexer

import (
	"lab/token"
)

type Lexer struct {
	input           string
	currentPosition int
	peekPosition    int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.movePosition()
	return l
}

func (l *Lexer) GetTokens() []token.Token {
	tokens := make([]token.Token, 0)
	currentToken := l.NextToken()
	tokens = append(tokens, currentToken)

	for currentToken.Type != token.EOF {
		currentToken = l.NextToken()
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()

	ch := l.currentChar()
	switch ch {
	case 0:
		t = token.New(token.EOF, "")
	case '=':
		if l.peekChar() == '=' {
			t = token.New(token.EQ, "==")
			l.movePosition()
		} else {
			t = token.New(token.ASSIGN, "=")
		}
	case '!':
		if l.peekChar() == '=' {
			t = token.New(token.NOT_EQ, "!=")
			l.movePosition()
		} else {
			t = token.New(token.NOT, "!")
		}
	case '&':
		if l.peekChar() == '&' {
			t = token.New(token.AND, "&&")
			l.movePosition()
		} else {
			t = token.New(token.ILLEGAL, "&")
		}
	case '|':
		if l.peekChar() == '|' {
			t = token.New(token.OR, "||")
			l.movePosition()
		} else {
			t = token.New(token.ILLEGAL, "|")
		}
	case '-':
		if l.peekChar() == '-' {
			t = token.New(token.DECREMENT, "--")
			l.movePosition()
		} else {
			t = token.New(token.MINUS, "-")
		}
	case '+':
		if l.peekChar() == '+' {
			t = token.New(token.INCREMENT, "++")
			l.movePosition()
		} else {
			t = token.New(token.PLUS, "+")
		}
	case '>':
		if l.peekChar() == '=' {
			t = token.New(token.GREAT_EQ, ">=")
			l.movePosition()
		} else {
			t = token.New(token.GREAT, ">")
		}
	case '<':
		if l.peekChar() == '=' {
			t = token.New(token.LESS_EQ, "<=")
			l.movePosition()
		} else {
			t = token.New(token.LESS, "<")
		}
	case '*':
		t = token.New(token.ASTERISK, "*")
	case '/':
		t = token.New(token.SLASH, "/")
	case ',':
		t = token.New(token.COMMA, ",")
	case ';':
		t = token.New(token.SEMICOLON, ";")
	case '(':
		t = token.New(token.LPAREN, "(")
	case ')':
		t = token.New(token.RPAREN, ")")
	case '{':
		t = token.New(token.LBRACE, "{")
	case '}':
		t = token.New(token.RBRACE, "}")
	case '"':
		l.movePosition()
		t = l.readString()
	default:
		if isLetter(ch) {
			t = l.readWord()
		} else if isDigit(ch) {
			t = l.readNumber()
		} else {
			t = token.New(token.ILLEGAL, string(ch))
		}
	}

	l.movePosition()
	return t
}

func (l *Lexer) movePosition() {
	l.currentPosition = l.peekPosition
	if l.peekPosition >= len(l.input) {
		return
	}
	l.peekPosition++
}

func (l *Lexer) backPosition() {
	l.peekPosition = l.currentPosition
	if l.currentPosition <= 0 {
		return
	}
	l.currentPosition--
}

func (l *Lexer) currentChar() byte {
	if l.currentPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.currentPosition]
	}
}

func (l *Lexer) peekChar() byte {
	if l.peekPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.peekPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	ch := l.currentChar()
	for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		l.movePosition()
		ch = l.currentChar()
	}
}

func (l *Lexer) readString() token.Token {
	start := l.currentPosition
	ch := l.currentChar()

	for ch != '"' && ch != 0 {
		l.movePosition()
		ch = l.currentChar()
	}

	return token.New(token.STR_LITERAL, l.input[start:l.currentPosition])
}

func (l *Lexer) readWord() token.Token {
	start := l.currentPosition
	ch := l.currentChar()

	for isLetter(ch) {
		l.movePosition()
		ch = l.currentChar()
	}

	l.backPosition()
	return token.GetKeywordToken(l.input[start : l.currentPosition+1])
}

func (l *Lexer) readNumber() token.Token {
	start := l.currentPosition
	ch := l.currentChar()

	for isDigit(ch) {
		l.movePosition()
		ch = l.currentChar()
	}

	if ch == '.' {
		l.movePosition()
		ch = l.currentChar()

		for isDigit(ch) {
			l.movePosition()
			ch = l.currentChar()
		}

		l.backPosition()
		return token.New(token.FLOAT_LITERAL, l.input[start:l.currentPosition+1])
	}

	l.backPosition()
	return token.New(token.INT_LITERAL, l.input[start:l.currentPosition+1])
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
