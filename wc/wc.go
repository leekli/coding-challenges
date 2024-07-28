package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"wc/utils"
)

// Global variables
var isStdIn bool = false
var filePathGiven string
var fileContents []byte

func main() {
	// Retrieve any args given by the user
	argsGiven := os.Args

	// First check if a standard input (stdin) pipe has been used
	isStdIn = utils.CheckIsStdIn()

	// If there is something in stdin, then read it and save to file contents
	if (isStdIn) {
		stdInput, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Println("Failed to read from stdin")
		}

		fileContents = stdInput
	}

	// If no stdin, then check and then read the file name/file path given by user and save to file contents
	if (!isStdIn) {
		filePathGiven = argsGiven[len(argsGiven)-1]

		if !utils.CheckFileExists(filePathGiven) {
			panic("File does not exist, or file is not valid.")
		}

		fileContents = utils.ReadFile(filePathGiven)
	}

	// Check which flag the user has provided and return the result
	if len(argsGiven) >= 2 {
		if argsGiven[1] == "-c" {
			// Count Total Bytes
			totalByteCount := countBytes(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[1] == "-l" {
			// Count Total Lines
			totalLineCount := countLines(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalLineCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[1] == "-w" {
			// Count Total Words
			totalWordCount := countWords(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalWordCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[1] == "-m" {
			// Count Total Characters
			totalCharCount := countChars(fileContents)

			formattedResult := fmt.Sprintf("    %d %s", totalCharCount, filePathGiven)

			fmt.Println(formattedResult)
		} else {
			fmt.Println("Invalid argument flag given.")
		}
	}
	
	// If no flag has been provided, then return the total bytes, lines and words count to user
	if (len(argsGiven) == 1 && isStdIn) || (len(argsGiven) == 2 && !isStdIn) {
		formattedResult := countWithNoFlags()

		fmt.Println(formattedResult)
	}

}

func countBytes(fileContents []byte) int {
	totalBytes := len(fileContents)

	return totalBytes
}

func countLines(fileContents []byte) int {
	lineCount := 0

	for _, byte := range fileContents {
		if byte == '\n' {
			lineCount++
		}
	}

	return lineCount
}

func countWords(fileContents []byte) int {
	fileContentsToString := string(fileContents)

	fileContentsSplitByWords := strings.Fields(fileContentsToString)

	wordCount := len(fileContentsSplitByWords)

	return wordCount
}

func countChars(fileContents []byte) int {
	charCount := len(fileContents)

	return charCount
}

func countWithNoFlags() string {
	totalByteCount := countBytes(fileContents)
	totalLineCount := countLines(fileContents)
	totalWordCount := countWords(fileContents)

	totalCountString := fmt.Sprintf("    %d   %d   %d %s", totalLineCount, totalWordCount, totalByteCount, filePathGiven)

	return totalCountString
}