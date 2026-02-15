package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Type tests (enum)
func TestTokenTypesConst_ChecksExistenceOfTokenTypes(test *testing.T) {
	assert.Equal(test, TokenType(0), TokenLeftBrace)
	assert.Equal(test, TokenType(1), TokenRightBrace)
	assert.Equal(test, TokenType(2), TokenLeftBracket)
	assert.Equal(test, TokenType(3), TokenRightBracket)
	assert.Equal(test, TokenType(4), TokenColon)
	assert.Equal(test, TokenType(5), TokenComma)
	assert.Equal(test, TokenType(6), TokenString)
	assert.Equal(test, TokenType(7), TokenNumber)
	assert.Equal(test, TokenType(8), TokenBoolean)
	assert.Equal(test, TokenType(9), TokenNull)
}