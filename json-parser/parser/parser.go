package parser

import (
	"json-parser/lexer"
	"strconv"
)

func Parse() bool {
	// STILL TO IMPLEMENT ONCE PARSE VALUE SETUP AND TESTED
	return false
}

func ParseValue(tokenList []lexer.Token, currentIndex int) (any, int) {
	token := Peek(tokenList, currentIndex)

	switch token.Type {
	case lexer.TokenTrueBoolean:
		_, idx := Consume(tokenList, currentIndex, lexer.TokenTrueBoolean)
		
		return true, idx
	case lexer.TokenFalseBoolean:
		_, idx := Consume(tokenList, currentIndex, lexer.TokenFalseBoolean)
		
		return false, idx
	case lexer.TokenNull:
		_, idx := Consume(tokenList, currentIndex, lexer.TokenNull)
		
		return nil, idx
	case lexer.TokenString:
		token, idx := Consume(tokenList, currentIndex, lexer.TokenString)
		
		return token.Value, idx
	case lexer.TokenNumber:
		token, idx := Consume(tokenList, currentIndex, lexer.TokenNumber)

		num, err := strconv.Atoi(token.Value)

		if err != nil {
			panic("❌ JSON Parser: Could not convert to number")
		}
		
		return num, idx
	case lexer.TokenLeftBrace:
		parsedObj, idx := ParseObject(tokenList, currentIndex)

		return parsedObj, idx
	case lexer.TokenLeftBracket:
		parsedArr, idx := ParseArray(tokenList, currentIndex)

		return parsedArr, idx
	default:
		panic("❌ JSON Parser: Unexpected token")
	}
}

func ParseObject(tokenList []lexer.Token, currentIndex int) (map[string]any, int) {
    var parsedObj = map[string]any{}

    _, currentIndex = Consume(tokenList, currentIndex, lexer.TokenLeftBrace)

    // Handle empty object
    nextToken := Peek(tokenList, currentIndex)

    if nextToken.Type == lexer.TokenRightBrace {
        _, currentIndex = Consume(tokenList, currentIndex, lexer.TokenRightBrace)

        return parsedObj, currentIndex
    }

    for {
        // Parse key (should be a string)
        keyToken, idx := Consume(tokenList, currentIndex, lexer.TokenString)
        currentIndex = idx

        // Parse colon
        _, idx = Consume(tokenList, currentIndex, lexer.TokenColon)
        currentIndex = idx

        // Parse value (could be any type & recursive)
        value, idx := ParseValue(tokenList, currentIndex)
        currentIndex = idx

        parsedObj[keyToken.Value] = value

        // Check for comma or right brace (end of object)
        next := Peek(tokenList, currentIndex)

        if next.Type == lexer.TokenComma {
            _, idx = Consume(tokenList, currentIndex, lexer.TokenComma)

            currentIndex = idx

            continue
        } else if next.Type == lexer.TokenRightBrace {
            _, currentIndex = Consume(tokenList, currentIndex, lexer.TokenRightBrace)

            break
        } else {
            panic("❌ JSON Parser: Object - expected , or }")
        }
    }

    return parsedObj, currentIndex
}

func ParseArray(tokenList []lexer.Token, currentIndex int) ([]any, int) {
    var parsedArray = []any{}

    _, currentIndex = Consume(tokenList, currentIndex, lexer.TokenLeftBracket)

    // Handle empty array
    nextToken := Peek(tokenList, currentIndex)

    if nextToken.Type == lexer.TokenRightBracket {
        _, currentIndex = Consume(tokenList, currentIndex, lexer.TokenRightBracket)

        return parsedArray, currentIndex
    }

    for {
        // Parse value (could be any type & recursive)
        value, idx := ParseValue(tokenList, currentIndex)
        currentIndex = idx

        parsedArray = append(parsedArray, value)

        // Check for comma or right bracket (end of array)
        next := Peek(tokenList, currentIndex)
		
        if next.Type == lexer.TokenComma {
            _, idx = Consume(tokenList, currentIndex, lexer.TokenComma)

            currentIndex = idx

            continue
        } else if next.Type == lexer.TokenRightBracket {
            _, currentIndex = Consume(tokenList, currentIndex, lexer.TokenRightBracket)

            break
        } else {
            panic("❌ JSON Parser: Object - expected , or ]")
        }
    }

    return parsedArray, currentIndex
}