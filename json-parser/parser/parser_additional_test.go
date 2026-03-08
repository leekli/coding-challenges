package parser

import (
	"json-parser/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_EmptyTokenList_ReturnsFalse(t *testing.T) {
    var tokens []lexer.Token
    _, ok := Parse(tokens)
    assert.False(t, ok, "Parse should return false for empty token list")
}

func TestParseValue_PanicsOnNonIntegerNumber(t *testing.T) {
    tokens := []lexer.Token{
        lexer.NewToken(lexer.TokenNumber, "1.23", 4),
    }
    assert.Panics(t, func() {
        _, _ = ParseValue(tokens, 0)
    }, "ParseValue should panic when given non-integer number token")
}

func TestParse_ReturnsFalseWhenTrailingNonEOFNonIllegalToken(t *testing.T) {
    // Example: a valid value followed by an unexpected comma
    tokens := []lexer.Token{
        lexer.NewToken(lexer.TokenTrueBoolean, "true", 4),
        lexer.NewToken(lexer.TokenComma, ",", 1),
    }
    _, ok := Parse(tokens)
    assert.False(t, ok, "Parse should return false when trailing unexpected token present")
}

func TestParse_ReturnsTrueWhenEOFFollowsValue(t *testing.T) {
    tokens := []lexer.Token{
        lexer.NewToken(lexer.TokenNumber, "42", 2),
        lexer.NewToken(lexer.TokenEOF, "", 0),
    }
    v, ok := Parse(tokens)
    assert.True(t, ok, "Parse should return true when EOF follows value")
    assert.Equal(t, 42, v, "Expected parsed value 42")
}
