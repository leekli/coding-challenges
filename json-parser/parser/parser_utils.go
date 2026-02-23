package parser

import (
	"json-parser/lexer"
	"log"
)

// Helper functions

func Peek(tokenList []lexer.Token, currentIndex int) lexer.Token {
    if currentIndex >= len(tokenList) {
		eofToken := lexer.NewToken(lexer.TokenEOF, "", 0)

        return eofToken
    }

    token := tokenList[currentIndex]
	
    return token
}

func Consume(tokenList []lexer.Token, currentIndex int, expectedTokenType lexer.TokenType) (lexer.Token, int) {
	newIndex := currentIndex

	token := Peek(tokenList, currentIndex)

	if token.Type == expectedTokenType {
		newIndex++
		return token, newIndex
	} else {
		log.Fatalf("❌ JSON Parser: Unexpected token type")
	}

	return token, newIndex
}