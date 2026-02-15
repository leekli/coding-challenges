package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Lexer function tests
func TestLexer_ReturnsEmptyTokenList_ForEmptySpace(test *testing.T) {
	testJson := ""

	output := Lexer(testJson)

	assert.Equal(test, len(output), 0)
}

func TestLexer_ReturnsEmptyTokenList_ForSingleWhitespace(test *testing.T) {
	testJson := " "

	output := Lexer(testJson)

	assert.Equal(test, len(output), 0)
}

func TestLexer_ReturnsEmptyTokenList_ForMultipleWhitespace(test *testing.T) {
	testJson := "    "

	output := Lexer(testJson)

	assert.Equal(test, len(output), 0)

	testJson = "                        "

	output = Lexer(testJson)

	assert.Equal(test, len(output), 0)
}

func TestLexer_ReturnsTokenList_ForSingleLeftBrace(test *testing.T) {
	testJson := "{"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenLeftBrace)
	assert.Equal(test, output[0].Value, "{")
	assert.Equal(test, output[0].Length, 1)
}

func TestLexer_ReturnsTokenList_ForSingleRightBrace(test *testing.T) {
	testJson := "}"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenRightBrace)
	assert.Equal(test, output[0].Value, "}")
	assert.Equal(test, output[0].Length, 1)
}

func TestLexer_ReturnsTokenList_ForSingleLeftBracket(test *testing.T) {
	testJson := "["

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[0].Length, 1)
}

func TestLexer_ReturnsTokenList_ForSingleRightBracket(test *testing.T) {
	testJson := "]"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenRightBracket)
	assert.Equal(test, output[0].Value, "]")
	assert.Equal(test, output[0].Length, 1)
}

func TestLexer_ReturnsTokenList_ForSingleColon(test *testing.T) {
	testJson := ":"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenColon)
	assert.Equal(test, output[0].Value, ":")
	assert.Equal(test, output[0].Length, 1)
}

func TestLexer_ReturnsTokenList_ForSingleComma(test *testing.T) {
	testJson := ","

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenComma)
	assert.Equal(test, output[0].Value, ",")
	assert.Equal(test, output[0].Length, 1)
}

func TestLexer_ReturnsTokenList_ForTrueBoolean(test *testing.T) {
	testJson := "true"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenBoolean)
	assert.Equal(test, output[0].Value, "true")
	assert.Equal(test, output[0].Length, 4)
}

func TestLexer_ReturnsTokenList_ForFalseBoolean(test *testing.T) {
	testJson := "false"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenBoolean)
	assert.Equal(test, output[0].Value, "false")
	assert.Equal(test, output[0].Length, 5)
}

func TestLexer_ReturnsTokenList_ForNull(test *testing.T) {
	testJson := "null"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNull)
	assert.Equal(test, output[0].Value, "null")
	assert.Equal(test, output[0].Length, 4)
}

func TestLexer_ReturnsTokenList_ForOpenAndClosedBrackets(test *testing.T) {
	testJson := "[]"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 2)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenRightBracket)
	assert.Equal(test, output[1].Value, "]")
}

func TestLexer_ReturnsTokenList_ForOpenAndClosedBraces(test *testing.T) {
	testJson := "{}"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 2)
	assert.Equal(test, output[0].Type, TokenLeftBrace)
	assert.Equal(test, output[0].Value, "{")
	assert.Equal(test, output[1].Type, TokenRightBrace)
	assert.Equal(test, output[1].Value, "}")
}

func TestLexer_ReturnsTokenList_ForString(test *testing.T) {
	testJson := `""`

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenString)
	assert.Equal(test, output[0].Value, `""`)
	assert.Equal(test, output[0].Length, 2)

	testJson = `"a"`

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenString)
	assert.Equal(test, output[0].Value, `"a"`)
	assert.Equal(test, output[0].Length, 3)

	testJson = `"abcdef"`

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenString)
	assert.Equal(test, output[0].Value, `"abcdef"`)
	assert.Equal(test, output[0].Length, 8)

	testJson = `"123"`

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenString)
	assert.Equal(test, output[0].Value, `"123"`)
	assert.Equal(test, output[0].Length, 5)
}

func TestLexer_ReturnsTokenList_ForPositiveNumber(test *testing.T) {
	testJson := "0"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "0")

	testJson = "1"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "1")

	testJson = "12"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "12")

	testJson = "1234567890"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "1234567890")
}

func TestLexer_ReturnsTokenList_ForNegativeNumber(test *testing.T) {
	testJson := "-1"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "-1")

	testJson = "-12"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "-12")

	testJson = "-1234567890"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 1)
	assert.Equal(test, output[0].Type, TokenNumber)
	assert.Equal(test, output[0].Value, "-1234567890")
}

func TestLexer_ReturnsTokenList_ForBracketsWithValueInside(test *testing.T) {
	testJson := "[true]"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 3)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenBoolean)
	assert.Equal(test, output[1].Value, "true")
	assert.Equal(test, output[2].Type, TokenRightBracket)
	assert.Equal(test, output[2].Value, "]")

	testJson = "[false]"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 3)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenBoolean)
	assert.Equal(test, output[1].Value, "false")
	assert.Equal(test, output[2].Type, TokenRightBracket)
	assert.Equal(test, output[2].Value, "]")

	testJson = "[null]"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 3)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenNull)
	assert.Equal(test, output[1].Value, "null")
	assert.Equal(test, output[2].Type, TokenRightBracket)
	assert.Equal(test, output[2].Value, "]")

	testJson = "[\"hello\"]"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 3)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenString)
	assert.Equal(test, output[1].Value, `"hello"`)
	assert.Equal(test, output[2].Type, TokenRightBracket)
	assert.Equal(test, output[2].Value, "]")

	testJson = "[1]"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 3)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenNumber)
	assert.Equal(test, output[1].Value, "1")
	assert.Equal(test, output[2].Type, TokenRightBracket)
	assert.Equal(test, output[2].Value, "]")
}

func TestLexer_ReturnsTokenList_ForBracketsWithMultipleValuesInside(test *testing.T) {
	testJson := "[true, false]"

	output := Lexer(testJson)

	assert.Equal(test, len(output), 5)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenBoolean)
	assert.Equal(test, output[1].Value, "true")
	assert.Equal(test, output[2].Type, TokenComma)
	assert.Equal(test, output[2].Value, ",")
	assert.Equal(test, output[3].Type, TokenBoolean)
	assert.Equal(test, output[3].Value, "false")
	assert.Equal(test, output[4].Type, TokenRightBracket)
	assert.Equal(test, output[4].Value, "]")

	testJson = "[false, null, true]"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 7)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenBoolean)
	assert.Equal(test, output[1].Value, "false")
	assert.Equal(test, output[2].Type, TokenComma)
	assert.Equal(test, output[2].Value, ",")
	assert.Equal(test, output[3].Type, TokenNull)
	assert.Equal(test, output[3].Value, "null")
	assert.Equal(test, output[4].Type, TokenComma)
	assert.Equal(test, output[4].Value, ",")
	assert.Equal(test, output[5].Type, TokenBoolean)
	assert.Equal(test, output[5].Value, "true")
	assert.Equal(test, output[6].Type, TokenRightBracket)
	assert.Equal(test, output[6].Value, "]")

	testJson = "[1, 2, 999]"

	output = Lexer(testJson)

	assert.Equal(test, len(output), 7)
	assert.Equal(test, output[0].Type, TokenLeftBracket)
	assert.Equal(test, output[0].Value, "[")
	assert.Equal(test, output[1].Type, TokenNumber)
	assert.Equal(test, output[1].Value, "1")
	assert.Equal(test, output[2].Type, TokenComma)
	assert.Equal(test, output[2].Value, ",")
	assert.Equal(test, output[3].Type, TokenNumber)
	assert.Equal(test, output[3].Value, "2")
	assert.Equal(test, output[4].Type, TokenComma)
	assert.Equal(test, output[4].Value, ",")
	assert.Equal(test, output[5].Type, TokenNumber)
	assert.Equal(test, output[5].Value, "999")
	assert.Equal(test, output[6].Type, TokenRightBracket)
	assert.Equal(test, output[6].Value, "]")
}