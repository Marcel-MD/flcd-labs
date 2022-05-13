package token

type Token struct {
	Type  TokenType
	Value string
}

func New(tokenType TokenType, value string) Token {
	return Token{Type: tokenType, Value: value}
}

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	NOT      = "NOT"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"

	LESS     = "LESS"
	GREAT    = "GREAT"
	LESS_EQ  = "LESS_EQ"
	GREAT_EQ = "GREAT_EQ"
	EQ       = "EQ"
	NOT_EQ   = "NOT_EQ"
	AND      = "AND"
	OR       = "OR"

	INCREMENT = "INCREMENT"
	DECREMENT = "DECREMENT"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	QUOTE     = "QUOTE"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	FLOAT      = "FLOAT"
	STR        = "STR"
	BOOL       = "BOOL"

	INT_LITERAL   = "INT_LITERAL"
	FLOAT_LITERAL = "FLOAT_LITERAL"
	STR_LITERAL   = "STR_LITERAL"

	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	FOR    = "FOR"
	WHILE  = "WHILE"
)

var keywords = map[string]TokenType{
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
	"while":  WHILE,
	"int":    INT,
	"float":  FLOAT,
	"str":    STR,
	"bool":   BOOL,
}

func GetKeywordToken(word string) Token {
	if tok, ok := keywords[word]; ok {
		return New(tok, word)
	}
	return New(IDENTIFIER, word)
}
