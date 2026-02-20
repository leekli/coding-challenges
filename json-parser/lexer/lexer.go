package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
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

	for !IsAtEnd(jsonInput, stringPointer) {
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
				stringValue, _ := ExtractString(jsonInput, stringPointer)
				stringLen := len(stringValue)

				token := NewToken(TokenString, stringValue, stringLen)
				tokenList = append(tokenList, token)		
				
				stringPointer += stringLen
			// Negative number value
			case "-":
				numValue, _ := ExtractNumber(jsonInput, stringPointer)
				numLen := len(numValue)

				token := NewToken(TokenNumber, numValue, numLen)
				tokenList = append(tokenList, token)		
				
				stringPointer += numLen
			}

			// Positive number value
			if charIsDigit {
				numValue, _ := ExtractNumber(jsonInput, stringPointer)
				numLen := len(numValue)

				token := NewToken(TokenNumber, numValue, numLen)
				tokenList = append(tokenList, token)		
				
				stringPointer += numLen
			}
		}		
	}

	return tokenList
}

func IsAtEnd(jsonInput string, stringPointer int) bool {
	return stringPointer >= len(jsonInput)
}

func Peek(jsonInput string, stringPointer int) string {
    if stringPointer < 0 || stringPointer >= len(jsonInput) {
        return ""
    }

    char := jsonInput[stringPointer]
	
    return string(char)
}

func isHexDigit(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func ExtractString(jsonInput string, stringPointer int) (string, error) {
	if Peek(jsonInput, stringPointer) != `"` {
		return "", fmt.Errorf("String must start with \"")
	}

	var foundString strings.Builder

	// Add opening quote
	foundString.WriteByte('"')
	stringPointer++

	for {
		if stringPointer >= len(jsonInput) {
			return "", fmt.Errorf("Unterminated string")
		}

		// Read next byte/rune
		b := jsonInput[stringPointer]

		// If next one is a ", then this is an empty string, add it as such
		if b == '"' {
			foundString.WriteByte('"')
			stringPointer++
			break
		}

		// Deal with various escape sequences
		if b == '\\' {
			// Need at least one more byte
			if stringPointer+1 >= len(jsonInput) {
				return "", fmt.Errorf("Invalid escape at end of input")
			}

			next := jsonInput[stringPointer + 1]

			switch next {
			case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
				// Preserve original escape sequence in token value
				foundString.WriteByte('\\')
				foundString.WriteByte(next)
				stringPointer += 2

				continue
			case 'u':
				// Expect exactly four hex digits after \u
				if stringPointer + 6 > len(jsonInput) {
					return "", fmt.Errorf("Incomplete unicode escape")
				}

				hex := jsonInput[stringPointer + 2 : stringPointer + 6]

				for i := 0; i < 4; i++ {
					if !isHexDigit(hex[i]) {
						return "", fmt.Errorf("Invalid unicode escape")
					}
				}

				// Append \uXXXX
				foundString.WriteString("\\u")
				foundString.WriteString(hex)

				// Move string pointer past \uXXXX
				stringPointer += 6

				continue
			default:
				return "", fmt.Errorf("Invalid escape sequence: \\%c", next)
			}
		}

		// Unescaped characters: must not be control characters (U+0000 through U+001F)
		r, size := utf8.DecodeRuneInString(jsonInput[stringPointer:])
		if r == utf8.RuneError && size == 1 {
			// Invalid UTF-8 byte sequence
			return "", fmt.Errorf("Invalid UTF-8 in string")
		}
		if r <= 0x1F {
			return "", fmt.Errorf("Unescaped control character in string")
		}

		// Append the raw bytes for this 
		foundString.WriteString(jsonInput[stringPointer : stringPointer + size])

		stringPointer += size
	}

	return foundString.String(), nil
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

func ExtractNumber(jsonInput string, stringPointer int) (string, error) {
	var foundNumber strings.Builder

	// Deal with optional single negative sign
	if Peek(jsonInput, stringPointer) == "-" {
		foundNumber.WriteString("-")

		stringPointer++		
	}

	// Deal with integer part, no leading 0s unless the integer is 0 on its own or 0.x
	if Peek(jsonInput, stringPointer) == "0" {
		foundNumber.WriteString(string(jsonInput[stringPointer]))

		stringPointer++	

		// Disallow leading zeros: if '0' is followed by another digit, it's invalid (e.g. "01")
        if isDigit(jsonInput, stringPointer) {
            return "", fmt.Errorf("Leading zeros are not allowed")
        }
	} else if isDigit(jsonInput, stringPointer) {
		for isDigit(jsonInput, stringPointer) {
			foundNumber.WriteString(string(jsonInput[stringPointer]))

			stringPointer++
		}
	} else {
		return "", fmt.Errorf("Number must start with a digit or minus sign")
	}

	// Deal with fractional/decimal part
	if Peek(jsonInput, stringPointer) == "." {
		foundNumber.WriteString(string(jsonInput[stringPointer]))

		stringPointer++	
		
		// Must have at least one digit after decimal point
		if !isDigit(jsonInput, stringPointer) {
			return "", fmt.Errorf("A digit must follow a decimal point")
		}

		for isDigit(jsonInput, stringPointer) {
			foundNumber.WriteString(string(jsonInput[stringPointer]))

			stringPointer++		
		}
	}

	// Deal with exponent part
	if Peek(jsonInput, stringPointer) == "e" || Peek(jsonInput, stringPointer) == "E" {
		foundNumber.WriteString(string(jsonInput[stringPointer]))

		stringPointer++	
		
		// Handle optional - or + signs on exponent
		if Peek(jsonInput, stringPointer) == "+" || Peek(jsonInput, stringPointer) == "-" {
			foundNumber.WriteString(string(jsonInput[stringPointer]))

			stringPointer++	
		}

		// Must have at least one digit in the exponent 
		if !isDigit(jsonInput, stringPointer) {
			return "", fmt.Errorf("Exponent must contain at least one digit")
		}

		for isDigit(jsonInput, stringPointer) {
			foundNumber.WriteString(string(jsonInput[stringPointer]))

			stringPointer++		
		}
	}

	return foundNumber.String(), nil
}