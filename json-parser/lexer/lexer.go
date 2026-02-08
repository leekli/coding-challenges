package lexer

import "fmt"

// Lexer JSON tokens types
type TokenType int

const (
	TokenLeftBrace TokenType = iota
	TokenRightBrace
	TokenLeftBracket
	TokenRightBracket
	TokenColon
	TokenComma
	TokenString
	TokenNumber
	TokenBoolean
	TokenNull
)

type Token struct {
	Type TokenType
	Value string
	Length int
}

func NewToken(tokenType TokenType, tokenValue string, tokenLength int) Token {
	token := Token{ tokenType, tokenValue, tokenLength }

	return token
}

func Lexer(jsonInput string) {
	fmt.Println(jsonInput)
}