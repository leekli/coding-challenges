package main

import (
	"strings"
	"testing"
)

var correctTestFile string = "test-data/sample.tsv"
var fakeTestFile string = "fake-folder/fake-file.tsv"

func TestBuildNoArgsGivenMsgContainsKeyWords(test *testing.T) {
	result := BuildNoArgsGivenMsg()

	if !strings.Contains(result, "usage") {
		test.Errorf("BuildNoArgsGivenMsg(): Does not contain 'usage' when it should")
	}
}

func TestCheckFileExistsReturnsTrue(test *testing.T) {
	result := CheckFileExists(correctTestFile)

	if result != true {
		test.Errorf("CheckFileExists(): File does not exist when it should exist")
	}
}

func TestCheckFileExistsReturnsFalse(test *testing.T) {
	result := CheckFileExists(fakeTestFile)

	if result != false {
		test.Errorf("CheckFileExists(): File exists when it should not exist")
	}
}

func TestPrintSpecifiedFieldReturnsErrorWithFieldSetToZero(test *testing.T) {
	result := PrintBySpecifiedField(correctTestFile, 0)

	if result != -1 {
		test.Errorf("PrintBySpecifiedField(): Function should have returned a -1 code")
	}
}

func TestPrintSpecifiedFieldReturnsErrorWithInvalidFile(test *testing.T) {
	result := PrintBySpecifiedField(fakeTestFile, 2)

	if result != -1 {
		test.Errorf("PrintBySpecifiedField(): Function should have returned a -1 code")
	}
}