package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Peek function tests
func TestPeek_ValidIndex(test *testing.T) {
	input := "abcdef"

	tests := []struct {
		index    int
		expected string
	}{
		{0, "a"},
		{1, "b"},
		{5, "f"},
	}

	for _, tt := range tests {
		got := Peek(input, tt.index)

		assert.Equal(test, tt.expected, got, "Peek(%q, %d)", input, tt.index)
	}
}

func TestPeek_NegativeIndex(test *testing.T) {
	input := "abcdef"

	got := Peek(input, -1)

	assert.Equal(test, "", got, "Peek(%q, -1)", input)
}

func TestPeek_IndexEqualToLength(test *testing.T) {
	input := "abcdef"

	got := Peek(input, len(input))

	assert.Equal(test, "", got, "Peek(%q, %d)", input, len(input))
}

func TestPeek_IndexGreaterThanLength(test *testing.T) {
	input := "abcdef"

	got := Peek(input, len(input) +1)

	assert.Equal(test, "", got, "Peek(%q, %d)", input, len(input)+1)
}

func TestPeek_EmptyString(test *testing.T) {
	input := ""

	got := Peek(input, 0)

	assert.Equal(test, "", got, "Peek(%q, 0)", input)
}

func TestPeek_UnicodeCharacter(test *testing.T) {
	input := "a😊b"
	// Note: Go strings are byte-indexed, so this will return partial runes for multi-byte characters

	got := Peek(input, 1) // Should return the first byte of the emoji

	expected := string(input[1])

	assert.Equal(test, expected, got, "Peek(%q, 1)", input)
}

// Extract string function tests
func TestExtractString_ReturnsString_EmptyString(test *testing.T) {
	testStr := `""`
	testPointer := 0

	output := ExtractString(testStr, testPointer)

	assert.Equal(test, `""`, output)
	assert.Equal(test, 2, len(output))
}

func TestExtractString_ReturnsString_SingleCharString(test *testing.T) {
	testStr := `"a"`
	testPointer := 0

	output := ExtractString(testStr, testPointer)

	assert.Equal(test, `"a"`, output)
	assert.Equal(test, 3, len(output))
}

func TestExtractString_ReturnsString_MultipleCharString(test *testing.T) {
	testStr := `"abc"`
	testPointer := 0

	output := ExtractString(testStr, testPointer)

	assert.Equal(test, `"abc"`, output)
	assert.Equal(test, 5, len(output))

	testStr = `"a b c"`
	testPointer = 0

	output = ExtractString(testStr, testPointer)

	assert.Equal(test, `"a b c"`, output)
	assert.Equal(test, 7, len(output))
}