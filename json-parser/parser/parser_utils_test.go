package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"json-parser/lexer"
)

// Peek happy path: valid indices
func TestPeek_ValidIndex(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, `"a"`, 3),
	}

	got := Peek(tokens, 0)
	assert.Equal(t, tokens[0], got)

	got = Peek(tokens, 1)
	assert.Equal(t, tokens[1], got)
}

// Peek sad path: index at or beyond length should return EOF token
func TestPeek_IndexAtOrBeyondLength_ReturnsEOF(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
	}

	eof := Peek(tokens, 1)
	assert.Equal(t, lexer.NewToken(lexer.TokenEOF, "", 0), eof)

	eof = Peek(tokens, 10)
	assert.Equal(t, lexer.NewToken(lexer.TokenEOF, "", 0), eof)
}

// Peek sad path: negative index should panic (out of bounds)
func TestPeek_NegativeIndex_Panics(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
	}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		_ = Peek(tokens, -1)
	}()

	assert.True(t, didPanic, "Peek should panic for negative index")
}

// Consume happy path: matching token type increments index and returns token
func TestConsume_HappyPath_ReturnsTokenAndIndex(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	tok, newIndex := Consume(tokens, 0, lexer.TokenLeftBrace)

	assert.Equal(t, tokens[0], tok)
	assert.Equal(t, 1, newIndex)
}

// Consume happy path at EOF: requesting EOF at end returns EOF token and advances index
func TestConsume_HappyPath_EOFAtEnd(t *testing.T) {
	tokens := []lexer.Token{}

	tok, newIndex := Consume(tokens, 0, lexer.TokenEOF)

	assert.Equal(t, lexer.Token{Type: lexer.TokenEOF, Value: "", Length: 0}, tok)
	assert.Equal(t, 1, newIndex)
}