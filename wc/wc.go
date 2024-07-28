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
		panic("No arguments given.")
	}

	filePathGiven := argsGiven[len(argsGiven)-1]

	if !utils.CheckFileExists(filePathGiven) {
		panic("File does not exist, or file is not valid.")
	}

	if utils.HasArgs(argsGiven) {
		if argsGiven[0] == "-c" {
			// Count Total Bytes
			totalByteCount := countBytes(filePathGiven)

			formattedResult := fmt.Sprintf("   %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[0] == "-l" {
			// Count Total Lines
			totalLineCount := countLines(filePathGiven)

			formattedResult := fmt.Sprintf("   %d %s", totalLineCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[0] == "-w" {
			// Count Total Words
			totalWordCount := countWords(filePathGiven)

			formattedResult := fmt.Sprintf("   %d %s", totalWordCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[0] == "-m" {
		// Count Total Characters
			totalCharCount := countChars(filePathGiven)

			formattedResult := fmt.Sprintf("    %d %s", totalCharCount, filePathGiven)

			fmt.Println(formattedResult)
		} else {
			filePathGiven = argsGiven[0]

			totalByteCount := countBytes(filePathGiven)
			totalLineCount := countLines(filePathGiven)
			totalWordCount := countWords(filePathGiven)

			formattedResult := fmt.Sprintf("    %d   %d   %d %s", totalLineCount, totalWordCount, totalByteCount, filePathGiven)

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