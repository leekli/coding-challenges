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

	value, _ := ParseValue(tokens, index)

	assert.Equal(t, value, true)
}

func TestParseValue_ShouldReturnGoFalseValue_ForFalseToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenFalseBoolean, "false", 5),
	}

	index := 0

	value, _ := ParseValue(tokens, index)

	assert.Equal(t, value, false)
}

func TestParseValue_ShouldReturnGoNilValue_ForNullToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenNull, "null", 4),
	}

	index := 0

	value, _ := ParseValue(tokens, index)

	assert.Equal(t, value, nil)
}

func TestParseValue_ShouldReturnGoStringValue_ForStringToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenString, "Hello", 5),
	}

	index := 0

	value, _ := ParseValue(tokens, index)

	assert.Equal(t, value, "Hello")
}

func TestParseValue_ShouldReturnGoStringValue_ForStringTokenWithSpace(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenString, "Hello World", 11),
	}

	index := 0

	value, _ := ParseValue(tokens, index)

	assert.Equal(t, value, "Hello World")
}

func TestParseValue_ShouldReturnGoIntValue_ForNumberToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenNumber, "1", 1),
	}

	index := 0

	value, _ := ParseValue(tokens, index)

	assert.Equal(t, value, 1)
}

// Additional string tests

func TestParseValue_ShouldReturnGoStringValueg_ForStringTokenIfEmpty(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenString, "", 0),
	}

	v, _ := ParseValue(tokens, 0)

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
		v, _ := ParseValue(tokens, 0)
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
		v, _ := ParseValue(tokens, 0)
		assert.Equal(t, c.expected, v)
	}
}

// TODO: Sort tests and logic for number decimal and exponents

// ParseObject tests
func TestParseObject_ShouldParseEmptyObject(t *testing.T) {

	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	obj, _ := ParseObject(tokens, 0)

	assert.Equal(t, 0, len(obj))
}

func TestParseObject_ShouldParseSingleStringPair(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "key", 3),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenString, "value", 5),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	obj, _ := ParseObject(tokens, 0)

	assert.Equal(t, 1, len(obj))
	assert.Equal(t, "value", obj["key"])
}

func TestParseObject_ShouldParseVariousValueTypes(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),

		lexer.NewToken(lexer.TokenString, "s", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenString, "str", 3),
		lexer.NewToken(lexer.TokenComma, ",", 1),

		lexer.NewToken(lexer.TokenString, "n", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "42", 2),
		lexer.NewToken(lexer.TokenComma, ",", 1),

		lexer.NewToken(lexer.TokenString, "t", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenTrueBoolean, "true", 4),
		lexer.NewToken(lexer.TokenComma, ",", 1),

		lexer.NewToken(lexer.TokenString, "f", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenFalseBoolean, "false", 5),
		lexer.NewToken(lexer.TokenComma, ",", 1),

		lexer.NewToken(lexer.TokenString, "u", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNull, "null", 4),

		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	obj, _ := ParseObject(tokens, 0)

	assert.Equal(t, "str", obj["s"])
	assert.Equal(t, 42, obj["n"])
	assert.Equal(t, true, obj["t"])
	assert.Equal(t, false, obj["f"])
	assert.Nil(t, obj["u"])
}

func TestParseObject_ShouldParseMultiplePairs(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "a", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenString, "b", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenString, "c", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "3", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	obj, _ := ParseObject(tokens, 0)

	assert.Equal(t, 3, len(obj))
	assert.Equal(t, 1, obj["a"])
	assert.Equal(t, 2, obj["b"])
	assert.Equal(t, 3, obj["c"])
}

func TestParseObject_ShouldParseNestedObject(t *testing.T) {
	tokens := []lexer.Token{
		// outer {
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "outer", 5),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		// inner {
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "inner", 5),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "7", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	obj, _ := ParseObject(tokens, 0)

	nested, ok := obj["outer"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, 7, nested["inner"])
}

// Extend ParseValue to include parsing objects
func TestParseValue_ShouldReturnObject_ForObjectTokens(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "a", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "10", 2),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}

	v, _ := ParseValue(tokens, 0)

	obj, ok := v.(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, 10, obj["a"])
}

func TestParseValue_ShouldReturnArray_ForArrayTokens(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	v, _ := ParseValue(tokens, 0)
	arr, ok := v.([]any)
	assert.True(t, ok)
	assert.Equal(t, []any{1, 2}, arr)
}

func TestParseValue_ShouldReturnObject_ForNestedObjectTokens(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "obj", 3),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "x", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "5", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}
	v, _ := ParseValue(tokens, 0)
	obj, ok := v.(map[string]any)
	assert.True(t, ok)
	nested, ok2 := obj["obj"].(map[string]any)
	assert.True(t, ok2)
	assert.Equal(t, 5, nested["x"])
}

func TestParseValue_ShouldReturnArray_ForNestedArrayTokens(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "3", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "4", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	v, _ := ParseValue(tokens, 0)
	arr, ok := v.([]any)
	assert.True(t, ok)
	arr1, ok1 := arr[0].([]any)
	arr2, ok2 := arr[1].([]any)
	assert.True(t, ok1)
	assert.True(t, ok2)
	assert.Equal(t, []any{1, 2}, arr1)
	assert.Equal(t, []any{3, 4}, arr2)
}

func TestParseValue_ShouldReturnObjectWithArrayValue(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "arr", 3),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}
	v, _ := ParseValue(tokens, 0)
	obj, ok := v.(map[string]any)
	assert.True(t, ok)
	arr, ok2 := obj["arr"].([]any)
	assert.True(t, ok2)
	assert.Equal(t, []any{1, 2}, arr)
}

func TestParseValue_ShouldReturnArrayWithObjectValue(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "x", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "5", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	v, _ := ParseValue(tokens, 0)
	arr, ok := v.([]any)
	assert.True(t, ok)
	obj, ok2 := arr[0].(map[string]any)
	assert.True(t, ok2)
	assert.Equal(t, 5, obj["x"])
}

func TestParseValue_ShouldFailOnUnexpectedToken(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenEOF, "", 0),
	}
	
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unexpected token, got none")
		}
	}()
	ParseValue(tokens, 0)
}

func TestParseValue_ShouldFailOnInvalidObjectStructure(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "key", 3),
		// Missing colon and value
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid object structure, got none")
		}
	}()
	ParseValue(tokens, 0)
}

func TestParseValue_ShouldFailOnInvalidArrayStructure(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1), // Comma without value
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid array structure, got none")
		}
	}()
	ParseValue(tokens, 0)
}

// ParseArray tests
func TestParseArray_ShouldParseEmptyArray(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	assert.Equal(t, 0, len(arr))
}

func TestParseArray_ShouldParseSingleString(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenString, "hello", 5),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	assert.Equal(t, 1, len(arr))
	assert.Equal(t, "hello", arr[0])
}

func TestParseArray_ShouldParseSingleNumber(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "42", 2),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 42, arr[0])
}

func TestParseArray_ShouldParseMultipleValues(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "3", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	assert.Equal(t, []any{1, 2, 3}, arr)
}

func TestParseArray_ShouldParseMixedTypes(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenString, "foo", 3),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "99", 2),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenTrueBoolean, "true", 4),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNull, "null", 4),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	assert.Equal(t, "foo", arr[0])
	assert.Equal(t, 99, arr[1])
	assert.Equal(t, true, arr[2])
	assert.Nil(t, arr[3])
}

func TestParseArray_ShouldParseArrayOfObjects(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "a", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenLeftBrace, "{", 1),
		lexer.NewToken(lexer.TokenString, "b", 1),
		lexer.NewToken(lexer.TokenColon, ":", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenRightBrace, "}", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	obj1, ok1 := arr[0].(map[string]any)
	obj2, ok2 := arr[1].(map[string]any)
	assert.True(t, ok1)
	assert.True(t, ok2)
	assert.Equal(t, 1, obj1["a"])
	assert.Equal(t, 2, obj2["b"])
}

func TestParseArray_ShouldParseNestedArrays(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "1", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "2", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenNumber, "3", 1),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNumber, "4", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	arr1, ok1 := arr[0].([]any)
	arr2, ok2 := arr[1].([]any)
	assert.True(t, ok1)
	assert.True(t, ok2)
	assert.Equal(t, []any{1, 2}, arr1)
	assert.Equal(t, []any{3, 4}, arr2)
}

func TestParseArray_ShouldParseArrayWithObjectAndArray(t *testing.T) {
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
	arr, _ := ParseArray(tokens, 0)
	obj, ok1 := arr[0].(map[string]any)
	arr2, ok2 := arr[1].([]any)
	assert.True(t, ok1)
	assert.True(t, ok2)
	assert.Equal(t, 5, obj["x"])
	assert.Equal(t, []any{6, 7}, arr2)
}

func TestParseArray_ShouldParseArrayWithBooleansAndNulls(t *testing.T) {
	tokens := []lexer.Token{
		lexer.NewToken(lexer.TokenLeftBracket, "[", 1),
		lexer.NewToken(lexer.TokenTrueBoolean, "true", 4),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenFalseBoolean, "false", 5),
		lexer.NewToken(lexer.TokenComma, ",", 1),
		lexer.NewToken(lexer.TokenNull, "null", 4),
		lexer.NewToken(lexer.TokenRightBracket, "]", 1),
	}
	arr, _ := ParseArray(tokens, 0)
	assert.Equal(t, true, arr[0])
	assert.Equal(t, false, arr[1])
	assert.Nil(t, arr[2])
}