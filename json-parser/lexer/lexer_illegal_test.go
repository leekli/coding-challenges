package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleQuoteCreatesIllegalToken(t *testing.T) {
    input := `'single'`
    tokens := Lexer(input)
    assert.Equal(t, 1, len(tokens))
    tok := tokens[0]
    assert.Equal(t, TokenIllegal, tok.Type)
    assert.Equal(t, "", tok.Value)
    assert.Equal(t, 0, tok.Length)
}

func TestDodgyIdentifierCreatesIllegalTokenWithValue(t *testing.T) {
    input := `False`
    tokens := Lexer(input)
    assert.Equal(t, 1, len(tokens))
    tok := tokens[0]
    assert.Equal(t, TokenIllegal, tok.Type)
    assert.Equal(t, "False", tok.Value)
    assert.Equal(t, len("False"), tok.Length)
}

func TestUnterminatedStringCreatesIllegalToken(t *testing.T) {
    input := `"unterminated`
    tokens := Lexer(input)
    assert.Equal(t, 1, len(tokens))
    assert.Equal(t, TokenIllegal, tokens[0].Type)
}

func TestNegativeNumberWithLeadingZeroCreatesIllegalToken(t *testing.T) {
    input := `-01`
    tokens := Lexer(input)
    assert.Equal(t, 1, len(tokens))
    assert.Equal(t, TokenIllegal, tokens[0].Type)
}

func TestIllegalAfterValidTokenStopsLexing(t *testing.T) {
    input := `true False`
    tokens := Lexer(input)
    assert.GreaterOrEqual(t, len(tokens), 2)
    assert.Equal(t, TokenTrueBoolean, tokens[0].Type)
    assert.Equal(t, TokenIllegal, tokens[1].Type)
}

func TestStringWithControlCharacterIsIllegal(t *testing.T) {
    // includes U+0001 control character between quotes
 
    // build string with a control character
    s := "\"a" + string(rune(1)) + "b\""
    tokens := Lexer(s)
    assert.Equal(t, 1, len(tokens))
    assert.Equal(t, TokenIllegal, tokens[0].Type)
}
