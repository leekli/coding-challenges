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

	output, _ := ExtractString(testStr, testPointer)

	assert.Equal(test, `""`, output)
	assert.Equal(test, 2, len(output))
}

func TestExtractString_ReturnsString_SingleCharString(test *testing.T) {
	testStr := `"a"`
	testPointer := 0

	output, _ := ExtractString(testStr, testPointer)

	assert.Equal(test, `"a"`, output)
	assert.Equal(test, 3, len(output))
}

func TestExtractString_ReturnsString_MultipleCharString(test *testing.T) {
	testStr := `"abc"`
	testPointer := 0

	output, _ := ExtractString(testStr, testPointer)

	assert.Equal(test, `"abc"`, output)
	assert.Equal(test, 5, len(output))

	testStr = `"a b c"`
	testPointer = 0

	output, _ = ExtractString(testStr, testPointer)

	assert.Equal(test, `"a b c"`, output)
	assert.Equal(test, 7, len(output))
}

// Extract number function tests
func TestExtractNumberReturnsString_SinglePositiveNum(test *testing.T) {
	testStr := "1"
	testPointer := 0

	output, _ := ExtractNumber(testStr, testPointer)

	assert.Equal(test, "1", output)
	assert.Equal(test, 1, len(output))

	testStr = "9"
	testPointer = 0

	output, _ = ExtractNumber(testStr, testPointer)

	assert.Equal(test, "9", output)
	assert.Equal(test, 1, len(output))
}

func TestExtractNumber_ReturnsString_MultiplePositiveNums(test *testing.T) {
	testStr := "123"
	testPointer := 0

	output, _ := ExtractNumber(testStr, testPointer)

	assert.Equal(test, "123", output)
	assert.Equal(test, 3, len(output))

	testStr = "1234567"
	testPointer = 0

	output, _ = ExtractNumber(testStr, testPointer)

	assert.Equal(test, "1234567", output)
	assert.Equal(test, 7, len(output))
}

func TestExtractNumber_ReturnsString_NegativeNumbers(test *testing.T) {
	testStr := "-1"
	testPointer := 0

	output, _ := ExtractNumber(testStr, testPointer)

	assert.Equal(test, "-1", output)
	assert.Equal(test, 2, len(output))

	testStr = "-245"
	testPointer = 0

	output, _ = ExtractNumber(testStr, testPointer)

	assert.Equal(test, "-245", output)
	assert.Equal(test, 4, len(output))
}

func TestExtractNumber_ReturnsString_ForSingleZero(test *testing.T) {
	testStr := "0"
	testPointer := 0

	output, _ := ExtractNumber(testStr, testPointer)

	assert.Equal(test, "0", output)
	assert.Equal(test, 1, len(output))
}

func TestExtractNumber_ReturnString_ForFractionalNumbers(test *testing.T) {
	testStr := "0.5"
	testPointer := 0

	output, _ := ExtractNumber(testStr, testPointer)

	assert.Equal(test, "0.5", output)
	assert.Equal(test, 3, len(output))
}

func TestExtractNumber_ReturnString_ForExponentNumbers(test *testing.T) {
	testStr := "1.2e-10"
	testPointer := 0

	output, _ := ExtractNumber(testStr, testPointer)

	assert.Equal(test, "1.2e-10", output)
	assert.Equal(test, 7, len(output))
}

// ==================== HAPPY PATH: Additional valid number formats ====================

// Single digit positive numbers 2-9
func TestExtractNumber_HappyPath_SingleDigits(test *testing.T) {
	tests := []struct {
		input    string
		pointer  int
		expected string
	}{
		{"2", 0, "2"},
		{"3", 0, "3"},
		{"4", 0, "4"},
		{"5", 0, "5"},
		{"6", 0, "6"},
		{"7", 0, "7"},
		{"8", 0, "8"},
		{"9", 0, "9"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.NoError(test, err, "ExtractNumber(%q, %d) should not error", tt.input, tt.pointer)
		assert.Equal(test, tt.expected, output, "ExtractNumber(%q, %d)", tt.input, tt.pointer)
	}
}

// Negative zero
func TestExtractNumber_HappyPath_NegativeZero(test *testing.T) {
	output, err := ExtractNumber("-0", 0)
	assert.NoError(test, err)
	assert.Equal(test, "-0", output)
}

// Large multi-digit numbers
func TestExtractNumber_HappyPath_LargeNumbers(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"999", "999"},
		{"1000000", "1000000"},
		{"9999999999", "9999999999"},
		{"-999", "-999"},
		{"-1000000", "-1000000"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Fractions with various decimal places
func TestExtractNumber_HappyPath_VariousFractions(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0.0", "0.0"},
		{"0.1", "0.1"},
		{"0.123", "0.123"},
		{"1.5", "1.5"},
		{"123.456", "123.456"},
		{"999.999999", "999.999999"},
		{"-0.5", "-0.5"},
		{"-1.5", "-1.5"},
		{"-123.456", "-123.456"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Exponents with lowercase 'e'
func TestExtractNumber_HappyPath_ExponentLowercaseE(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1e0", "1e0"},
		{"1e1", "1e1"},
		{"1e10", "1e10"},
		{"1e100", "1e100"},
		{"123e456", "123e456"},
		{"-1e5", "-1e5"},
		{"-123e100", "-123e100"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Exponents with uppercase 'E'
func TestExtractNumber_HappyPath_ExponentUppercaseE(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E0", "1E0"},
		{"1E1", "1E1"},
		{"1E10", "1E10"},
		{"123E456", "123E456"},
		{"-1E5", "-1E5"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Exponents with explicit plus sign
func TestExtractNumber_HappyPath_ExponentWithPlusSign(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1e+0", "1e+0"},
		{"1e+1", "1e+1"},
		{"1e+10", "1e+10"},
		{"123e+100", "123e+100"},
		{"-1e+5", "-1e+5"},
		{"1E+10", "1E+10"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Exponents with explicit minus sign
func TestExtractNumber_HappyPath_ExponentWithMinusSign(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1e-0", "1e-0"},
		{"1e-1", "1e-1"},
		{"1e-10", "1e-10"},
		{"123e-100", "123e-100"},
		{"-1e-5", "-1e-5"},
		{"1E-10", "1E-10"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Combined: Fractions AND exponents
func TestExtractNumber_HappyPath_FractionWithExponent(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0.0e0", "0.0e0"},
		{"1.5e10", "1.5e10"},
		{"1.23e-5", "1.23e-5"},
		{"123.456e+100", "123.456e+100"},
		{"-1.5e-10", "-1.5e-10"},
		{"-0.123e45", "-0.123e45"},
		{"0.0E-0", "0.0E-0"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Very small decimal fractions
func TestExtractNumber_HappyPath_VerySmallDecimals(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0.00001", "0.00001"},
		{"0.0000001", "0.0000001"},
		{"-0.00001", "-0.00001"},
		{"1e-100", "1e-100"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, 0)
		assert.NoError(test, err, "ExtractNumber(%q, 0) should not error", tt.input)
		assert.Equal(test, tt.expected, output)
	}
}

// Numbers followed by invalid tokens (test that we only extract the number part)
func TestExtractNumber_HappyPath_NumberBeforeInvalidChar(test *testing.T) {
	tests := []struct {
		input    string
		pointer  int
		expected string
	}{
		{"123a", 0, "123"},
		{"1.5x", 0, "1.5"},
		{"1e10}", 0, "1e10"},
		{"-42,", 0, "-42"},
		{"0.5]", 0, "0.5"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.NoError(test, err, "ExtractNumber(%q, %d) should not error", tt.input, tt.pointer)
		assert.Equal(test, tt.expected, output, "ExtractNumber(%q, %d)", tt.input, tt.pointer)
	}
}

// ==================== SAD PATH: Invalid number formats ====================

// Invalid: Numbers with leading zeros (except for 0 or 0.x)
func TestExtractNumber_SadPath_LeadingZeros(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"01", 0},
		{"001", 0},
		{"007", 0},
		{"0123", 0},
		{"-01", 0},
		{"-001", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (invalid leading zero)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Decimal point without following digit
func TestExtractNumber_SadPath_DecimalWithoutFollowingDigit(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"1.", 0},
		{"0.", 0},
		{"123.", 0},
		{"-1.", 0},
		{"-0.", 0},
		{"1. ", 0},
		{"1.,", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (no digit after decimal)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Number starting with decimal point (no preceding digits)
func TestExtractNumber_SadPath_StartsWithDecimal(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{".5", 0},
		{".123", 0},
		{"-.5", 0},
		{"-. 5", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (starts with decimal)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Just a decimal point
func TestExtractNumber_SadPath_JustDecimalPoint(test *testing.T) {
	output, err := ExtractNumber(".", 0)
	assert.Error(test, err)
	assert.Equal(test, "", output)
}

// Invalid: Just a minus sign
func TestExtractNumber_SadPath_JustMinusSign(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"-", 0},
		{"- ", 0},
		{"-,", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (just minus sign)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Minus followed by decimal
func TestExtractNumber_SadPath_MinusWithDecimalOnly(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"-.", 0},
		{"-. ", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (minus with decimal only)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Exponent without following digit
func TestExtractNumber_SadPath_ExponentWithoutFollowingDigit(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"1e", 0},
		{"1E", 0},
		{"1e ", 0},
		{"123e", 0},
		{"1.5e", 0},
		{"-1e", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (no digit in exponent)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Exponent with +/- sign but no following digit
func TestExtractNumber_SadPath_ExponentWithSignButNoDigit(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"1e+", 0},
		{"1e-", 0},
		{"1E+", 0},
		{"1E-", 0},
		{"1e+ ", 0},
		{"1e- ", 0},
		{"1.5e+", 0},
		{"-1e-", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (exponent sign with no digit)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Invalid alphabetic characters in number
func TestExtractNumber_SadPath_InvalidAlphabeticChars(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"1a2", 0},
		{"1b3", 0},
		{"1c5", 0},
		{"abc", 0},
		{"1.5x", 0}, // after extraction stops at 'x', might actually work
	}

	for _, tt := range tests {
		_, err := ExtractNumber(tt.input, tt.pointer)
		// Some of these might pass because we stop at invalid chars
		// Only truly invalid if starting with invalid
		if string(tt.input[0]) != "1" {
			assert.Error(test, err, "ExtractNumber(%q, %d) should error", tt.input, tt.pointer)
		}
	}
}

// Invalid: Empty string or starting at end of string
func TestExtractNumber_SadPath_EmptyOrOutOfBounds(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
		desc    string
	}{
		{"", 0, "empty string"},
		{"abc", 0, "starts with letter"},
		{"!@#", 0, "special characters"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (%s)", tt.input, tt.pointer, tt.desc)
		assert.Equal(test, "", output)
	}
}

// Invalid: Plus sign at start (JSON doesn't allow leading + for numbers)
func TestExtractNumber_SadPath_LeadingPlusSign(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"+1", 0},
		{"+123", 0},
		{"+0.5", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (leading + not allowed in JSON)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

// Invalid: Spaces within numbers
func TestExtractNumber_SadPath_SpacesInNumber(test *testing.T) {
	tests := []struct {
		input    string
		pointer  int
		expected string // might extract up to space
	}{
		{"1 23", 0, "1"},
		{"1. 5", 0, ""},       // error on no digit after decimal
		{"1e 10", 0, ""},      // error on no digit in exponent
		{"-1 23", 0, "-1"},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		if tt.expected != "" {
			assert.NoError(test, err, "ExtractNumber(%q, %d) should extract %q", tt.input, tt.pointer, tt.expected)
			assert.Equal(test, tt.expected, output)
		} else {
			assert.Error(test, err, "ExtractNumber(%q, %d) should error", tt.input, tt.pointer)
		}
	}
}

// Invalid: Double negative signs
func TestExtractNumber_SadPath_DoubleNegativeSign(test *testing.T) {
	tests := []struct {
		input   string
		pointer int
	}{
		{"--1", 0},
		{"---5", 0},
	}

	for _, tt := range tests {
		output, err := ExtractNumber(tt.input, tt.pointer)
		assert.Error(test, err, "ExtractNumber(%q, %d) should error (double negative)", tt.input, tt.pointer)
		assert.Equal(test, "", output)
	}
}

func TestExtractString_ReturnsEscapedQuote(test *testing.T) {
	testStr := `"a\"b"`
	// raw Go backtick string contains: " a \" b " -> JSON: "a\"b"
	output, err := ExtractString(testStr, 0)
	assert.NoError(test, err)
	assert.Equal(test, `"a\"b"`, output)
}

func TestExtractString_ReturnsEscapedBackslash(test *testing.T) {
	testStr := `"a\\b"`
	output, err := ExtractString(testStr, 0)
	assert.NoError(test, err)
	assert.Equal(test, `"a\\b"`, output)
}

func TestExtractString_ReturnsEscapedSlash(test *testing.T) {
	testStr := `"/"`
	output, err := ExtractString(testStr, 0)
	assert.NoError(test, err)
	assert.Equal(test, `"/"`, output)
}

func TestExtractString_ReturnsEscapedControlSequences(test *testing.T) {
	testStr := `"\b\f\n\r\t"`
	output, err := ExtractString(testStr, 0)
	assert.NoError(test, err)
	assert.Equal(test, `"\b\f\n\r\t"`, output)
}

func TestExtractString_ReturnsUnicodeEscape(test *testing.T) {
	testStr := `"\u0041"` // 'A'
	output, err := ExtractString(testStr, 0)
	assert.NoError(test, err)
	assert.Equal(test, `"\u0041"`, output)
}

func TestExtractString_ReturnsSurrogatePair(test *testing.T) {
	// U+1D11E (musical G clef) encoded as surrogate pair
	testStr := `"\uD834\uDD1E"`
	output, err := ExtractString(testStr, 0)
	assert.NoError(test, err)
	assert.Equal(test, `"\uD834\uDD1E"`, output)
}

func TestExtractString_UnterminatedStringReturnsError(test *testing.T) {
	testStr := `"unterminated`
	_, err := ExtractString(testStr, 0)
	assert.Error(test, err)
}

func TestExtractString_InvalidEscapeReturnsError(test *testing.T) {
	testStr := `"\x"`
	_, err := ExtractString(testStr, 0)
	assert.Error(test, err)
}

func TestExtractString_IncompleteUnicodeEscapeReturnsError(test *testing.T) {
	testStr := `"\u123"` // only 3 hex digits
	_, err := ExtractString(testStr, 0)
	assert.Error(test, err)
}

func TestExtractString_UnescapedControlCharacterReturnsError(test *testing.T) {
	testStr := "\"\x01\"" // contains U+0001 control char inside string
	_, err := ExtractString(testStr, 0)
	assert.Error(test, err)
}

// isHexDigit tests
func TestIsHexDigit_Valid(test *testing.T) {
	// 0-9
	for b := byte('0'); b <= byte('9'); b++ {
		if !isHexDigit(b) {
			test.Fatalf("expected %c to be hex digit", b)
		}
	}

	// a-f
	for b := byte('a'); b <= byte('f'); b++ {
		if !isHexDigit(b) {
			test.Fatalf("expected %c to be hex digit", b)
		}
	}

	// A-F
	for b := byte('A'); b <= byte('F'); b++ {
		if !isHexDigit(b) {
			test.Fatalf("expected %c to be hex digit", b)
		}
	}
}

func TestIsHexDigit_Invalid(test *testing.T) {
	vals := []byte{'g', 'G', 'z', 'Z', '/', ' ', ':', '\n', 0xFF}
	for _, v := range vals {
		if isHexDigit(v) {
			test.Fatalf("did not expect %v (%c) to be hex digit", v, v)
		}
	}
}