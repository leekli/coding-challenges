package parser

import (
	"json-parser/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseReturnsFalseForIllegalStart(t *testing.T) {
    input := `'oops'`
    tokens := lexer.Lexer(input)
    _, ok := Parse(tokens)
    assert.False(t, ok, "parse should be invalid for illegal start")
}

func TestParseReturnsFalseWhenIllegalInObject(t *testing.T) {
    input := `{"a": 'b'}`
    tokens := lexer.Lexer(input)
    _, ok := Parse(tokens)
    assert.False(t, ok, "parse should be invalid for illegal token inside object")
}

func TestParseReturnsFalseWhenIllegalIdentifierAfterValue(t *testing.T) {
    input := `true False`
    tokens := lexer.Lexer(input)
    _, ok := Parse(tokens)
    assert.False(t, ok, "parse should be invalid when illegal token follows a valid value")
}

func TestParseReturnsFalseForBadIdentifierAlone(t *testing.T) {
    input := `False`
    tokens := lexer.Lexer(input)
    _, ok := Parse(tokens)
    assert.False(t, ok, "parse should be invalid for bad identifier alone")
}
