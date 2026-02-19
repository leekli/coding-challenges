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

// Additional comprehensive tests for lexer (numbers and strings)
func TestLexer_NumberVariants_SingleToken(test *testing.T) {
	tests := []struct {
 		input string
 	 	expected string
 	}{
 		{"0", "0"},
 		{"10", "10"},
 		{"-10", "-10"},
 		{"0.5", "0.5"},
 		{"123.456", "123.456"},
 		{"1e10", "1e10"},
 		{"1E-10", "1E-10"},
 		{"1e+10", "1e+10"},
 		{"-1.23e-10", "-1.23e-10"},
 	}

 	for _, tt := range tests {
 		output := Lexer(tt.input)
 		assert.Equal(test, len(output), 1)
 		assert.Equal(test, output[0].Type, TokenNumber)
 		assert.Equal(test, output[0].Value, tt.expected)
 	}
}

func TestLexer_ArrayOfNumbers(test *testing.T) {
 	testJson := "[1, 2, 999]"

 	output := Lexer(testJson)

 	assert.Equal(test, len(output), 7)
 	assert.Equal(test, output[0].Type, TokenLeftBracket)
 	assert.Equal(test, output[1].Type, TokenNumber)
 	assert.Equal(test, output[1].Value, "1")
 	assert.Equal(test, output[2].Type, TokenComma)
 	assert.Equal(test, output[3].Type, TokenNumber)
 	assert.Equal(test, output[3].Value, "2")
 	assert.Equal(test, output[4].Type, TokenComma)
 	assert.Equal(test, output[5].Type, TokenNumber)
 	assert.Equal(test, output[5].Value, "999")
 	assert.Equal(test, output[6].Type, TokenRightBracket)
}

func TestLexer_ObjectWithNumberValue(test *testing.T) {
 	testJson := "{\"a\": 42}"

 	output := Lexer(testJson)

 	assert.Equal(test, len(output), 5)
 	assert.Equal(test, output[0].Type, TokenLeftBrace)
 	assert.Equal(test, output[1].Type, TokenString)
 	assert.Equal(test, output[1].Value, "\"a\"")
 	assert.Equal(test, output[2].Type, TokenColon)
 	assert.Equal(test, output[3].Type, TokenNumber)
 	assert.Equal(test, output[3].Value, "42")
 	assert.Equal(test, output[4].Type, TokenRightBrace)
}

func TestLexer_StringVariants(test *testing.T) {
 	tests := []struct {
 		input string
 		expected string
 	}{
 		{"\"\"", "\"\""},
 		{"\"hello\"", "\"hello\""},
 		{"\"a b c\"", "\"a b c\""},
 		{"\"123\"", "\"123\""},
 	}

 	for _, tt := range tests {
 		output := Lexer(tt.input)
 		assert.Equal(test, len(output), 1)
 		assert.Equal(test, output[0].Type, TokenString)
 		assert.Equal(test, output[0].Value, tt.expected)
 	}
}

func TestLexer_MixedArray(test *testing.T) {
 	testJson := "[\"a\", 1, true, null, {\"k\":\"v\"}]"

 	output := Lexer(testJson)

 	// Validate presence and sequence of token types (not every value)
 	expectedTypes := []TokenType{TokenLeftBracket, TokenString, TokenComma, TokenNumber, TokenComma, TokenBoolean, TokenComma, TokenNull, TokenComma, TokenLeftBrace, TokenString, TokenColon, TokenString, TokenRightBrace, TokenRightBracket}

 	assert.Equal(test, len(output), len(expectedTypes))

 	for i, tt := range expectedTypes {
 		assert.Equal(test, output[i].Type, tt)
 	}
}

func TestLexer_NumberTokenLength(test *testing.T) {
	input := "12345"
	out := Lexer(input)
	if assert.Equal(test, 1, len(out)) {
		assert.Equal(test, TokenNumber, out[0].Type)
		assert.Equal(test, "12345", out[0].Value)
		assert.Equal(test, len(out[0].Value), out[0].Length)
	}
}

func TestLexer_AdjacentPunctuationNoSpaces(test *testing.T) {
	input := "[1,2,3]"
	out := Lexer(input)
	expectedTypes := []TokenType{TokenLeftBracket, TokenNumber, TokenComma, TokenNumber, TokenComma, TokenNumber, TokenRightBracket}
	assert.Equal(test, len(expectedTypes), len(out))
	for i := range expectedTypes {
		assert.Equal(test, expectedTypes[i], out[i].Type)
	}
	// lengths for numbers
	assert.Equal(test, 1, out[1].Length)
	assert.Equal(test, 1, out[3].Length)
	assert.Equal(test, 1, out[5].Length)
}

func TestLexer_NegativeNumberLength(test *testing.T) {
	input := "-12345"
	out := Lexer(input)
	assert.Equal(test, 1, len(out))
	assert.Equal(test, TokenNumber, out[0].Type)
	assert.Equal(test, "-12345", out[0].Value)
	assert.Equal(test, len(out[0].Value), out[0].Length)
}

func TestLexer_LargeIntegerAndExponent(test *testing.T) {
	a := "12345678901234567890"
	out := Lexer(a)
	assert.Equal(test, 1, len(out))
	assert.Equal(test, a, out[0].Value)
	assert.Equal(test, len(a), out[0].Length)

	b := "1.2e+308"
	out2 := Lexer(b)
	assert.Equal(test, 1, len(out2))
	assert.Equal(test, b, out2[0].Value)
	assert.Equal(test, len(b), out2[0].Length)
}

func TestLexer_NestedStructures_NumbersAndStrings(test *testing.T) {
	input := "{\"a\":[{\"b\":0.5},-1E10],\"c\":\"x\"}"
	out := Lexer(input)
	// find numeric tokens and check values
	nums := []string{}
	strs := []string{}
	for _, t := range out {
		if t.Type == TokenNumber {
			nums = append(nums, t.Value)
		}
		if t.Type == TokenString {
			strs = append(strs, t.Value)
		}
	}
	assert.Contains(test, nums, "0.5")
	assert.Contains(test, nums, "-1E10")
	assert.Contains(test, strs, "\"a\"")
	assert.Contains(test, strs, "\"b\"")
	assert.Contains(test, strs, "\"c\"")
}

func TestLexer_MultipleTokensSequence(test *testing.T) {
	input := "true,false,null,0"
	out := Lexer(input)
	expected := []TokenType{TokenBoolean, TokenComma, TokenBoolean, TokenComma, TokenNull, TokenComma, TokenNumber}
	assert.Equal(test, len(expected), len(out))
	for i := range expected {
		assert.Equal(test, expected[i], out[i].Type)
	}
}

func TestLexer_EmptyNestedStructures(test *testing.T) {
	input := "[{},[]]"
	out := Lexer(input)
	expected := []TokenType{TokenLeftBracket, TokenLeftBrace, TokenRightBrace, TokenComma, TokenLeftBracket, TokenRightBracket, TokenRightBracket}
	assert.Equal(test, len(expected), len(out))
	for i := range expected {
		assert.Equal(test, expected[i], out[i].Type)
	}
}

func TestLexer_NumberFollowedByPunctuation(test *testing.T) {
	input := "0.5]"
	out := Lexer(input)
	assert.Equal(test, 2, len(out))
	assert.Equal(test, TokenNumber, out[0].Type)
	assert.Equal(test, "0.5", out[0].Value)
	assert.Equal(test, TokenRightBracket, out[1].Type)

	input2 := "-42,"
	out2 := Lexer(input2)
	assert.Equal(test, 2, len(out2))
	assert.Equal(test, TokenNumber, out2[0].Type)
	assert.Equal(test, "-42", out2[0].Value)
	assert.Equal(test, TokenComma, out2[1].Type)
}

func TestLexer_LeadingTrailingSpaces(test *testing.T) {
	input := "  1  "
	out := Lexer(input)
	assert.Equal(test, 1, len(out))
	assert.Equal(test, TokenNumber, out[0].Type)
	assert.Equal(test, "1", out[0].Value)
}

func TestIsDigit_Bounds(test *testing.T) {
	assert.True(test, isDigit("123", 0))
	assert.True(test, isDigit("123", 2))
	assert.False(test, isDigit("123", -1))
	assert.False(test, isDigit("123", 3))
	// multi-byte rune: emoji bytes should not be digits
	assert.False(test, isDigit("a😊b", 1))
}

func TestPeek_MultiByteAndBounds(test *testing.T) {
	s := "a😊b"
	// first byte
	assert.Equal(test, string(s[0]), Peek(s, 0))
	// second byte (part of emoji)
	assert.Equal(test, string(s[1]), Peek(s, 1))
	// out of bounds
	assert.Equal(test, "", Peek(s, -1))
	assert.Equal(test, "", Peek(s, len(s)))
}

func TestLexer_StringsWithEscapes_ViaLexer(test *testing.T) {
	input := `"a\"b"` // JSON text: "a\"b"
	out := Lexer(input)
	assert.Equal(test, 1, len(out))
	assert.Equal(test, TokenString, out[0].Type)
	assert.Equal(test, `"a\"b"`, out[0].Value)
	assert.Equal(test, len(out[0].Value), out[0].Length)

	// in an array
	input2 := ` [ "\\" , "\u0041" ] `
	out2 := Lexer(input2)
	// expected tokens: [, string, comma, string, ] => 5
	assert.Equal(test, 5, len(out2))
	assert.Equal(test, TokenLeftBracket, out2[0].Type)
	assert.Equal(test, TokenString, out2[1].Type)
	assert.Equal(test, `"\\"`, out2[1].Value)
	assert.Equal(test, TokenComma, out2[2].Type)
	assert.Equal(test, TokenString, out2[3].Type)
	assert.Equal(test, `"\u0041"`, out2[3].Value)
	assert.Equal(test, TokenRightBracket, out2[4].Type)
}
