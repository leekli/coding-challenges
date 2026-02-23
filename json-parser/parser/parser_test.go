package parser

import (
	"json-parser/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValue_ShouldReturnGoTrueValue_ForTrueToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenTrueBoolean, "true", 4),
	}

	index := 0

	value := ParseValue(tokens, index)

	assert.Equal(t, value, true)
}

func TestParseValue_ShouldReturnGoFalseValue_ForFalseToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenFalseBoolean, "false", 5),
	}

	index := 0

	value := ParseValue(tokens, index)

	assert.Equal(t, value, false)
}

func TestParseValue_ShouldReturnGoNilValue_ForNullToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenNull, "null", 4),
	}

	index := 0

	value := ParseValue(tokens, index)

	assert.Equal(t, value, nil)
}

func TestParseValue_ShouldReturnGoStringValue_ForStringToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenString, "Hello", 5),
	}

	index := 0

	value := ParseValue(tokens, index)

	assert.Equal(t, value, "Hello")
}

func TestParseValue_ShouldReturnGoStringValue_ForStringTokenWithSpace(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenString, "Hello World", 11),
	}

	index := 0

	value := ParseValue(tokens, index)

	assert.Equal(t, value, "Hello World")
}

func TestParseValue_ShouldReturnGoIntValue_ForNumberToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenNumber, "1", 1),
	}

	index := 0

	value := ParseValue(tokens, index)

	assert.Equal(t, value, 1)
}

// Additional string tests

func TestParseValue_ShouldReturnGoStringValueg_ForStringTokenIfEmpty(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenString, "", 0),
	}

	v := ParseValue(tokens, 0)

	assert.Equal(t, "", v)
}

// Integer number tests (happy path)
func TestParseValue_ShouldReturnGoIntValue_ForPositiveIntegers(t *testing.T) {
	cases := []struct{
		in string
		expected int
	}{
		{"0", 0},
		{"1", 1},
		{"9", 9},
		{"42", 42},
		{"123456", 123456},
	}

	for _, c := range cases {
		tokens := []lexer.Token{lexer.NewToken(lexer.TokenNumber, c.in, len(c.in))}
		v := ParseValue(tokens, 0)
		assert.Equal(t, c.expected, v)
	}
}

func TestParseValue_ShouldReturnGoIntValue_ForNegativeIntegers(t *testing.T) {
	cases := []struct{
		in string
		expected int
	}{
		{"-1", -1},
		{"-42", -42},
		{"-100000", -100000},
	}

	for _, c := range cases {
		tokens := []lexer.Token{lexer.NewToken(lexer.TokenNumber, c.in, len(c.in))}
		v := ParseValue(tokens, 0)
		assert.Equal(t, c.expected, v)
	}
}

// TODO: Sort tests and logic for number decimal and exponents