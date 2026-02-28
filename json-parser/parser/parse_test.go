package parser

import (
	"json-parser/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Primitives_HappyPath(t *testing.T) {
    cases := []struct{
        name string
        tokens []lexer.Token
        expected any
    }{
        {"true", []lexer.Token{lexer.NewToken(lexer.TokenTrueBoolean, "true", 4)}, true},
        {"false", []lexer.Token{lexer.NewToken(lexer.TokenFalseBoolean, "false", 5)}, false},
        {"null", []lexer.Token{lexer.NewToken(lexer.TokenNull, "null", 4)}, nil},
        {"string", []lexer.Token{lexer.NewToken(lexer.TokenString, "hello", 5)}, "hello"},
        {"number", []lexer.Token{lexer.NewToken(lexer.TokenNumber, "42", 2)}, 42},
        {"negative number", []lexer.Token{lexer.NewToken(lexer.TokenNumber, "-7", 2)}, -7},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            v, ok := Parse(c.tokens)
            assert.True(t, ok)
            assert.Equal(t, c.expected, v)
        })
    }
}

func TestParse_Object_Array_HappyPath(t *testing.T) {
    t.Run("empty object", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
            lexer.NewToken(lexer.TokenRightBrace, "}", 1),
        }
        v, ok := Parse(tokens)
        assert.True(t, ok)
        m, _ := v.(map[string]any)
        assert.Equal(t, 0, len(m))
    })

    t.Run("object with values", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
            lexer.NewToken(lexer.TokenString, "a", 1),
            lexer.NewToken(lexer.TokenColon, ":", 1),
            lexer.NewToken(lexer.TokenNumber, "1", 1),
            lexer.NewToken(lexer.TokenComma, ",", 1),
            lexer.NewToken(lexer.TokenString, "b", 1),
            lexer.NewToken(lexer.TokenColon, ":", 1),
            lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
            lexer.NewToken(lexer.TokenNumber, "2", 1),
            lexer.NewToken(lexer.TokenRightBracket, "]", 1),
            lexer.NewToken(lexer.TokenRightBrace, "}", 1),
        }
        v, ok := Parse(tokens)
        assert.True(t, ok)
        m := v.(map[string]any)
        assert.Equal(t, 1, m["a"])
        arr := m["b"].([]any)
        assert.Equal(t, []any{2}, arr)
    })

    t.Run("empty array", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
            lexer.NewToken(lexer.TokenRightBracket, "]", 1),
        }
        v, ok := Parse(tokens)
        assert.True(t, ok)
        a := v.([]any)
        assert.Equal(t, 0, len(a))
    })

    t.Run("array of objects and arrays", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
            lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
            lexer.NewToken(lexer.TokenString, "x", 1),
            lexer.NewToken(lexer.TokenColon, ":", 1),
            lexer.NewToken(lexer.TokenNumber, "5", 1),
            lexer.NewToken(lexer.TokenRightBrace, "}", 1),
            lexer.NewToken(lexer.TokenComma, ",", 1),
            lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
            lexer.NewToken(lexer.TokenNumber, "6", 1),
            lexer.NewToken(lexer.TokenComma, ",", 1),
            lexer.NewToken(lexer.TokenNumber, "7", 1),
            lexer.NewToken(lexer.TokenRightBracket, "]", 1),
            lexer.NewToken(lexer.TokenRightBracket, "]", 1),
        }
        v, ok := Parse(tokens)
        assert.True(t, ok)
        top := v.([]any)
        obj := top[0].(map[string]any)
        assert.Equal(t, 5, obj["x"])
        arr := top[1].([]any)
        assert.Equal(t, []any{6, 7}, arr)
    })
}

func TestParse_TrailingEOFToken_IsAccepted(t *testing.T) {
    tokens := []lexer.Token{
        lexer.NewToken(lexer.TokenNumber, "1", 1),
        lexer.NewToken(lexer.TokenEOF, "", 0),
    }
    v, ok := Parse(tokens)
    assert.True(t, ok)
    assert.Equal(t, 1, v)
}

// Sad path tests - Parse should not panic but should return (nil, false)
func TestParse_InvalidInputs_ReturnsFalse(t *testing.T) {
    t.Run("unexpected EOF token only", func(t *testing.T) {
        tokens := []lexer.Token{lexer.NewToken(lexer.TokenEOF, "", 0)}
        v, ok := Parse(tokens)
        assert.False(t, ok)
        assert.Nil(t, v)
    })

    t.Run("incomplete object (missing colon/value)", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
            lexer.NewToken(lexer.TokenString, "k", 1),
            lexer.NewToken(lexer.TokenRightBrace, "}", 1),
        }
        v, ok := Parse(tokens)
        assert.False(t, ok)
        assert.Nil(t, v)
    })

    t.Run("incomplete array (comma without value)", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
            lexer.NewToken(lexer.TokenComma, ",", 1),
            lexer.NewToken(lexer.TokenRightBracket, "]", 1),
        }
        v, ok := Parse(tokens)
        assert.False(t, ok)
        assert.Nil(t, v)
    })

    t.Run("extra tokens after value", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenNumber, "1", 1),
            lexer.NewToken(lexer.TokenComma, ",", 1),
            lexer.NewToken(lexer.TokenNumber, "2", 1),
        }
        v, ok := Parse(tokens)
        assert.False(t, ok)
        assert.Nil(t, v)
    })

    t.Run("trailing comma in array", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
            lexer.NewToken(lexer.TokenNumber, "1", 1),
            lexer.NewToken(lexer.TokenComma, ",", 1),
            lexer.NewToken(lexer.TokenRightBracket, "]", 1),
        }
        v, ok := Parse(tokens)
        assert.False(t, ok)
        assert.Nil(t, v)
    })

    t.Run("missing closing brace", func(t *testing.T) {
        tokens := []lexer.Token{
            lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
            lexer.NewToken(lexer.TokenString, "a", 1),
            lexer.NewToken(lexer.TokenColon, ":", 1),
            lexer.NewToken(lexer.TokenNumber, "1", 1),
            // no right brace
        }
        v, ok := Parse(tokens)
        assert.False(t, ok)
        assert.Nil(t, v)
    })
}
