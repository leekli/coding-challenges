package parser

import (
	"json-parser/lexer"
	"log"
	"strconv"
)

func Parse() bool {
	// STILL TO IMPLEMENT ONCE PARSE VALUE SETUP AND TESTED
	return false
}

func ParseValue(tokenList []lexer.Token, currentIndex int) any {
	token := Peek(tokenList, currentIndex)

	switch token.Type {
	case lexer.TokenTrueBoolean:
		Consume(tokenList, currentIndex, lexer.TokenTrueBoolean)
		
		return true
	case lexer.TokenFalseBoolean:
		Consume(tokenList, currentIndex, lexer.TokenFalseBoolean)
		
		return false
	case lexer.TokenNull:
		Consume(tokenList, currentIndex, lexer.TokenNull)
		
		return nil
	case lexer.TokenString:
		token, _ := Consume(tokenList, currentIndex, lexer.TokenString)
		
		return token.Value
	case lexer.TokenNumber:
		token, _ := Consume(tokenList, currentIndex, lexer.TokenNumber)

		num, err := strconv.Atoi(token.Value)

		if err != nil {
			log.Fatalf("❌ JSON Parser: Could not convert to number")
		}
		
		return num
	default:
		log.Fatalf("❌ JSON Parser: Unexpected token")

		return -99
	}
}