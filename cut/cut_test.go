package main

import (
	"strings"
	"testing"
)

// Tab test file
var correctTestFile string = "test-data/sample.tsv"
var fakeTestFile string = "fake-folder/fake-file.tsv"

// Comma test file
var correctTestCommeFile string = "test-data/fourchords.csv"

func TestCutHelpMessage_ReturnsFullCutHelpMessage(test *testing.T) {
	result := CutHelpMessage()

	if !strings.Contains(result, "usage") {
		test.Errorf("BuildNoArgsGivenMsg(): Does not contain 'usage' when it should")
	}

	if !strings.Contains(result, "cut") {
		test.Errorf("BuildNoArgsGivenMsg(): Does not contain 'cut' when it should")
	}
}

func TestCheckFileExists_ReturnsTrueForValidFile(test *testing.T) {
	result := CheckFileExists(correctTestFile)

	if result != true {
		test.Errorf("CheckFileExists(): File does not exist when it should exist")
	}
}

func TestCheckFileExists_ReturnsFalseForValidFile(test *testing.T) {
	result := CheckFileExists(fakeTestFile)

	if result != false {
		test.Errorf("CheckFileExists(): File exists when it should not exist")
	}
}

func TestPrintBySpecifiedField_ReturnsErrorIfFieldNumIsZero_WithDefaultDelimiter(test *testing.T) {
	filePath := correctTestFile
	fieldNum := 0

	code, output := PrintBySpecifiedField(filePath, fieldNum, "\t")

	if code != -1 {
		test.Errorf("Expected: -1, Received: %d", code)
	}

	if !strings.Contains(output, "cut: values may not include zero") {
		test.Errorf("Expected output to contain: 'cut: values may not include zero', Received: %s", output)
	}

	filePath = fakeTestFile
	fieldNum = 0

	code, output = PrintBySpecifiedField(filePath, fieldNum, "")

	if code != -1 {
		test.Errorf("Expected: -1, Received: %d", code)
	}

	if !strings.Contains(output, "cut: values may not include zero") {
		test.Errorf("Expected output to contain: 'cut: values may not include zero', Received: %s", output)
	}
}

func TestPrintBySpecifiedField_ReturnsErrorIfFieldPathIsInvalid_WithDefaultDelimiter(test *testing.T) {
	filePath := fakeTestFile
	fieldNum := 2

	code, output := PrintBySpecifiedField(filePath, fieldNum, "\t")

	if code != -1 {
		test.Errorf("Expected: -1, Received: %d", code)
	}

	if !strings.Contains(output, "Error opening file") {
		test.Errorf("Expected output to contain: 'Error opening file', Received: %s", output)
	}
}

func TestPrintBySpecifiedField_ReturnsCorrectCodeAndOutputForValidFile_WithDefaultDelimiter(test *testing.T) {
	filePath := correctTestFile
	fieldNum := 2

	code, output := PrintBySpecifiedField(filePath, fieldNum, "\t")

	if code != 0 {
		test.Errorf("Expected: 0, Received: %d", code)
	}

	if !strings.Contains(output, "f1") {
		test.Errorf("Expected output to contain: 'f1'")
	}

	if !strings.Contains(output, "1") {
		test.Errorf("Expected output to contain: '1'")
	}

	if !strings.Contains(output, "16") {
		test.Errorf("Expected output to contain: '16'")
	}

	fieldNum = 3

	code, output = PrintBySpecifiedField(filePath, fieldNum, "\t")

	if code != 0 {
		test.Errorf("Expected: 0, Received: %d", code)
	}

	if !strings.Contains(output, "f2") {
		test.Errorf("Expected output to contain: 'f2'")
	}

	if !strings.Contains(output, "2") {
		test.Errorf("Expected output to contain: '2'")
	}

	if !strings.Contains(output, "22") {
		test.Errorf("Expected output to contain: '22'")
	}
}

func TestPrintBySpecifiedField_ReturnsCorrectCodeAndOutputForValidFile_WithCommaDelimiter(test *testing.T) {
	filePath := correctTestCommeFile
	fieldNum := 1
	delimiterChar := ","

	code, output := PrintBySpecifiedField(filePath, fieldNum, delimiterChar)

	if code != 0 {
		test.Errorf("Expected: 0, Received: %d", code)
	}

	if !strings.Contains(output, "Song title") {
		test.Errorf("Expected output to contain: 'Song title'")
	}

	if !strings.Contains(output, "All You Wanted") {
		test.Errorf("Expected output to contain: 'All You Wanted'")
	}

	if !strings.Contains(output, "Hair") {
		test.Errorf("Expected output to contain: 'Hair'")
	}

	if !strings.Contains(output, "Zombie") {
		test.Errorf("Expected output to contain: 'Zombie'")
	}
}