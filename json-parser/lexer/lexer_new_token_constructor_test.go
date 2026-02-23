package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// New Token constructor function tests
func TestNewTokenConstructor_ReturnsEachTokenType(test *testing.T) {
	var leftBraceToken = NewToken(TokenLeftBrace, "}", 1)

	assert.Equal(test, leftBraceToken.Type, TokenLeftBrace)
	assert.Equal(test, leftBraceToken.Value, "}")
	assert.Equal(test, leftBraceToken.Length, 1)

	var rightBraceToken = NewToken(TokenRightBrace, "{", 1)

	assert.Equal(test, rightBraceToken.Type, TokenRightBrace)
	assert.Equal(test, rightBraceToken.Value, "{")
	assert.Equal(test, rightBraceToken.Length, 1)
	
	var leftBracketToken = NewToken(TokenLeftBracket, "[", 1)

	assert.Equal(test, leftBracketToken.Type, TokenLeftBracket)
	assert.Equal(test, leftBracketToken.Value, "[")
	assert.Equal(test, leftBracketToken.Length, 1)

	var rightBracketToken = NewToken(TokenRightBracket, "]", 1)

	assert.Equal(test, rightBracketToken.Type, TokenRightBracket)
	assert.Equal(test, rightBracketToken.Value, "]")
	assert.Equal(test, rightBracketToken.Length, 1)

	var colonToken = NewToken(TokenColon, ":", 1)

	assert.Equal(test, colonToken.Type, TokenColon)
	assert.Equal(test, colonToken.Value, ":")
	assert.Equal(test, colonToken.Length, 1)

	var commaToken = NewToken(TokenComma, ",", 1)

	assert.Equal(test, commaToken.Type, TokenComma)
	assert.Equal(test, commaToken.Value, ",")
	assert.Equal(test, commaToken.Length, 1)

	var stringToken = NewToken(TokenString, "\"Hello\"", 7)

	assert.Equal(test, stringToken.Type, TokenString)
	assert.Equal(test, stringToken.Value, "\"Hello\"")
	assert.Equal(test, stringToken.Length, 7)

	var wholeNumberToken = NewToken(TokenNumber, "123", 3)

	assert.Equal(test, wholeNumberToken.Type, TokenNumber)
	assert.Equal(test, wholeNumberToken.Value, "123")
	assert.Equal(test, wholeNumberToken.Length, 3)

	var fracNumberToken = NewToken(TokenNumber, "1.23", 4)

	assert.Equal(test, fracNumberToken.Type, TokenNumber)
	assert.Equal(test, fracNumberToken.Value, "1.23")
	assert.Equal(test, fracNumberToken.Length, 4)

	var expNumberToken = NewToken(TokenNumber, "1e23", 4)

	assert.Equal(test, expNumberToken.Type, TokenNumber)
	assert.Equal(test, expNumberToken.Value, "1e23")
	assert.Equal(test, expNumberToken.Length, 4)

	var trueBooleanToken = NewToken(TokenTrueBoolean, "true", 4)

	assert.Equal(test, trueBooleanToken.Type, TokenTrueBoolean)
	assert.Equal(test, trueBooleanToken.Value, "true")
	assert.Equal(test, trueBooleanToken.Length, 4)

	var falseBooleanToken = NewToken(TokenFalseBoolean, "false", 5)

	assert.Equal(test, falseBooleanToken.Type, TokenFalseBoolean)
	assert.Equal(test, falseBooleanToken.Value, "false")
	assert.Equal(test, falseBooleanToken.Length, 5)

	var nullToken = NewToken(TokenNull, "null", 4)

	assert.Equal(test, nullToken.Type, TokenNull)
	assert.Equal(test, nullToken.Value, "null")
	assert.Equal(test, nullToken.Length, 4)
}