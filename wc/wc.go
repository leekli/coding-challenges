package main

import (
	"fmt"
	"os"
	"strings"
	"wc/utils"
)

func main() {
	argsGiven := os.Args[1:]

	if !utils.HasArgs(argsGiven) {
		panic("No arguments given")
	}

	filePathGiven := argsGiven[len(argsGiven)-1]

	if utils.HasArgs(argsGiven) {
		// Count Total Bytes
		if argsGiven[0] == "-c" {
			totalByteCount := countBytes(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		}

		// Count Total Lines
		if argsGiven[0] == "-l" {
			totalLineCount := countLines(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalLineCount, filePathGiven)

			fmt.Println(formattedResult)
		}

		// Count Total Words
		if argsGiven[0] == "-w" {
			totalWordCount := countWords(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalWordCount, filePathGiven)

			fmt.Println(formattedResult)
		}

		// Count Total Characters
		if argsGiven[0] == "-m" {
			totalCharCount := countChars(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalCharCount, filePathGiven)

			fmt.Println(formattedResult)
		}
	}
}

func countBytes(filePath string) int {
	fileContents := utils.ReadFile(filePath)

	totalBytes := len(fileContents)

	return totalBytes
}

func countLines(filePath string) int {
	fileContents := utils.ReadFile(filePath)

	lineCount := 0

	for _, byte := range fileContents {
		if byte == '\n' {
			lineCount++
		}
	}

	return lineCount
}

func countWords(filePath string) int {
	fileContents := utils.ReadFile(filePath)

	fileContentsToString := string(fileContents)

	fileContentsSplitByWords := strings.Fields(fileContentsToString)

	wordCount := len(fileContentsSplitByWords)

	return wordCount
}

func countChars(filePath string) int {
	fileContents := utils.ReadFile(filePath)

	charCount := len(fileContents)

	return charCount
}