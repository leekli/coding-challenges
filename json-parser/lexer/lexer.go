package lexer

import (
	"strings"
	"unicode"
)

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

func Lexer(jsonInput string) []Token {
	tokenList := []Token{}
	stringPointer := 0

	for stringPointer < len(jsonInput) {
		currentChar := string(jsonInput[stringPointer])

		charIsDigit := unicode.IsDigit(rune(jsonInput[stringPointer]))

		switch currentChar {
		// Whitespace (Skip)
		case " ":
			stringPointer++

			continue
		// Open [left] brace
		case "{":
			token := NewToken(TokenLeftBrace, "{", len(currentChar))
			tokenList = append(tokenList, token)

			stringPointer++
		// Closed [right] brace
		case "}":
			token := NewToken(TokenRightBrace, "}", len(currentChar))
			tokenList = append(tokenList, token)

			stringPointer++
		// Open [left] bracket
		case "[":
			token := NewToken(TokenLeftBracket, "[", len(currentChar))
			tokenList = append(tokenList, token)

			stringPointer++
		// Closed [right] bracket
		case "]":
			token := NewToken(TokenRightBracket, "]", len(currentChar))
			tokenList = append(tokenList, token)

			stringPointer++		
		// Colon
		case ":":
			token := NewToken(TokenColon, ":", len(currentChar))
			tokenList = append(tokenList, token)

			stringPointer++	
		// Comma
		case ",":
			token := NewToken(TokenComma, ",", len(currentChar))
			tokenList = append(tokenList, token)

			stringPointer++	
		// true boolean			
		case "t":
			len := 4

			token := NewToken(TokenBoolean, "true", len)
			tokenList = append(tokenList, token)		
			
			stringPointer += len
		// false boolean
		case "f":
			len := 5

			token := NewToken(TokenBoolean, "false", len)
			tokenList = append(tokenList, token)		
			
			stringPointer += len
		// null value
		case "n":
			len := 4

			token := NewToken(TokenNull, "null", len)
			tokenList = append(tokenList, token)		
			
			stringPointer += len
		// String values
		case `"`:
			stringValue := ExtractString(jsonInput, stringPointer)
			stringLen := len(stringValue)

			token := NewToken(TokenString, stringValue, stringLen)
			tokenList = append(tokenList, token)		
			
			stringPointer += stringLen
		// Negative number value
		case "-":
			numValue := ExtractNumber(jsonInput, stringPointer)
			numLen := len(numValue)

			token := NewToken(TokenNumber, numValue, numLen)
			tokenList = append(tokenList, token)		
			
			stringPointer += numLen
		}

		// Positive number value
		if charIsDigit {
			numValue := ExtractNumber(jsonInput, stringPointer)
			numLen := len(numValue)

			token := NewToken(TokenNumber, numValue, numLen)
			tokenList = append(tokenList, token)		
			
			stringPointer += numLen
		}
	}

	return tokenList
}

func Peek(jsonInput string, stringPointer int) string {
    if stringPointer < 0 || stringPointer >= len(jsonInput) {
        return ""
    }

    char := jsonInput[stringPointer]
	
    return string(char)
}

func ExtractString(jsonInput string, stringPointer int) string {
	// NO LOGIC RIGHT NOW FOR ESCAPE CHARS etc YET!
	// FRAGILE IF NO END " IS FOUND
	
	var foundString strings.Builder

	// First char is " as identified by lexer (move pointer forward by 1)
	foundString.WriteString(`"`)
	stringPointer++

	// Loop through each subsequent char and keep adding to found string until a " end string quote is found
	for Peek(jsonInput, stringPointer) != `"` {
		nextChar := Peek(jsonInput, stringPointer)

		foundString.WriteString(nextChar)

		stringPointer++
	}

	// Add end string quote
	if string(jsonInput[stringPointer]) == `"` {
		foundString.WriteString(`"`)
	}

	return foundString.String()
}

func isDigit(jsonInput string, stringPointer int) bool {
    if stringPointer < 0 || stringPointer >= len(jsonInput) {
        return false
    }

	isCharDigit := unicode.IsDigit(rune(jsonInput[stringPointer]))

	if isCharDigit {
		return true
	}

	return false
}

func ExtractNumber(jsonInput string, stringPointer int) string {
	// Need to extend to include: decimals, 0.x's, exponents
	// To add tests

	var foundNumber strings.Builder

	// Deal with optional single negative sign
	if jsonInput[stringPointer] == '-' {
		foundNumber.WriteString("-")

		stringPointer++		
	}

	// Get the first digit, check validity
	if isDigit(jsonInput, stringPointer) {
		foundNumber.WriteString(string(jsonInput[stringPointer]))

		stringPointer++
	}

	// Cycle through the rest of the digits, checking validity
	for isDigit(jsonInput, stringPointer) {
		foundNumber.WriteString(string(jsonInput[stringPointer]))

		stringPointer++
	}

	return foundNumber.String()
}