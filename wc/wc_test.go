package main

import (
	"strconv"
	"strings"
	"testing"
	"wc/utils"
)

/*
	- TESTS TO DO:
		1. CountWords: Try with different delimiters: spaces, tabs, newlines
		2. CountWords: Test with special characters and punctuation
		3. CountChars: Test with files with only whitespace
		4. Test main() with different args
		5. Test main() with some stdin
*/

func readTestFile(filePath string) []byte {
	fileContents := utils.ReadFile(filePath)

	return fileContents
}

func TestCountBytes_ReturnsCorrectBytesCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountBytes(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountBytes_ReturnsCorrectBytesCountForUTF8File(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountBytes(input)

	if output != 342190 {
		test.Errorf("Expected: 342190, Received: %d", output)
	}
}

func TestCountBytes_ReturnsCorrectBytesCountForASCIIFile(test *testing.T) {
	input := readTestFile("test_files/test_ascii.txt")

	output := CountBytes(input)

	if output != 15088 {
		test.Errorf("Expected: 15088, Received: %d", output)
	}
}

func TestCountLines_ReturnsCorrectLineCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountLines(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountLines_ReturnsCorrectLineCountForFileWithContents(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountLines(input)

	if output != 7145 {
		test.Errorf("Expected: 7145, Received: %d", output)
	}

	input = readTestFile("test_files/test_ascii.txt")

	output = CountLines(input)

	if output != 204 {
		test.Errorf("Expected: 204, Received: %d", output)
	}
}

func TestCountLines_ReturnsCorrectLineCountForVariousNewLineChars(test *testing.T) {
	input := readTestFile("test_files/test_newlines.txt")

	output := CountLines(input)

	if output != 8 {
		test.Errorf("Expected: 8, Received: %d", output)
	}
}

func TestCountWords_ReturnsCorrectWordCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountWords(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountWords_ReturnsCorrectWordCountForFileWithContents(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountWords(input)

	if output != 58164 {
		test.Errorf("Expected: 58164, Received: %d", output)
	}

	input = readTestFile("test_files/test_ascii.txt")

	output = CountWords(input)

	if output != 2000 {
		test.Errorf("Expected: 2000, Received: %d", output)
	}
}

func TestCountChars_ReturnsCorrectCharCountForEmptyFile(test *testing.T) {
	input := readTestFile("test_files/empty_test.txt")

	output := CountChars(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountChars_ReturnsCorrectCharCountForFileWithContents(test *testing.T) {
	input := readTestFile("test_files/test.txt")

	output := CountChars(input)

	if output != 339292 {
		test.Errorf("Expected: 339292, Received: %d", output)
	}

	input = readTestFile("test_files/test_ascii.txt")

	output = CountChars(input)

	if output != 15088 {
		test.Errorf("Expected: 15088, Received: %d", output)
	}
}

func TestCountWithNoFlags_ReturnsCorrectStringOfCountsForUTF8File(test *testing.T) {
	fileName := "test_files/test.txt"

	input := readTestFile(fileName)

	outputStr := CountWithNoFlags()
	byteCount := CountBytes(input)
	lineCount := CountLines(input)
	wordCount := CountWords(input)

	byteCountStr := strconv.Itoa(byteCount)
	lineCountStr := strconv.Itoa(lineCount)
	wordCountStr := strconv.Itoa(wordCount)

	if strings.Contains(outputStr, byteCountStr) {
		test.Errorf("Expected: %s to be present, but it was not.", byteCountStr)
	}

	if strings.Contains(outputStr, lineCountStr) {
		test.Errorf("Expected: %s to be present, but it was not.", lineCountStr)
	}

	if strings.Contains(outputStr, wordCountStr) {
		test.Errorf("Expected: %s to be present, but it was not.", wordCountStr)
	}

	if strings.Contains(outputStr, fileName) {
		test.Errorf("Expected: %s to be present, but it was not.", fileName)
	}
}

func TestCountWithNoFlags_ReturnsCorrectStringOfCountsForASCIIFile(test *testing.T) {
	fileName := "test_files/test_ascii.txt"

	input := readTestFile(fileName)

	outputStr := CountWithNoFlags()
	byteCount := CountBytes(input)
	lineCount := CountLines(input)
	wordCount := CountWords(input)

	byteCountStr := strconv.Itoa(byteCount)
	lineCountStr := strconv.Itoa(lineCount)
	wordCountStr := strconv.Itoa(wordCount)

	if strings.Contains(outputStr, byteCountStr) {
		test.Errorf("Expected: %s to be present, but it was not.", byteCountStr)
	}

	if strings.Contains(outputStr, lineCountStr) {
		test.Errorf("Expected: %s to be present, but it was not.", lineCountStr)
	}

	if strings.Contains(outputStr, wordCountStr) {
		test.Errorf("Expected: %s to be present, but it was not.", wordCountStr)
	}

	if strings.Contains(outputStr, fileName) {
		test.Errorf("Expected: %s to be present, but it was not.", fileName)
	}
}