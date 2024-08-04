package utils

import (
	"testing"
)

func TestFileExistsReturnsFalse(t *testing.T) {
	testPath := "wrongdir/nofile.json"

	result := checkFileExists(testPath)

	expected := false

	if result != expected {
		t.Error()
	}
}

func TestFileExistsReturnsTrue(t *testing.T) {
	testPath := "../test-data/task1-valid.json"

	result := checkFileExists(testPath)

	expected := true

	if result != expected {
		t.Error()
	}
}

func TestCharIsValidTokenReturnsTrue(t *testing.T) {
	var testByte byte = 123

	testTokenList := []string{"{", "}"}

	expected := true

	result := CharIsValidToken(testByte, testTokenList)

	if result != expected {
		t.Error()
	}

	testByte = 125

	result = CharIsValidToken(testByte, testTokenList)

	if result != expected {
		t.Error()
	}
}

func TestCharIsValidTokenReturnsFalse(t *testing.T) {
	var testByte byte = 124

	testTokenList := []string{"{", "}"}

	expected := false

	result := CharIsValidToken(testByte, testTokenList)

	if result != expected {
		t.Error()
	}

	testByte = 176

	result = CharIsValidToken(testByte, testTokenList)

	if result != expected {
		t.Error()
	}
}