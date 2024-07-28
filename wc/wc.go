package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"wc/utils"
)

func main() {
	argsGiven := os.Args[1:]

	filePathGiven := argsGiven[len(argsGiven)-1]

	var fileContents []byte

	if len(argsGiven) > 1 {
		if !utils.CheckFileExists(filePathGiven) {
			panic("File does not exist, or file is not valid.")
		} else {
			fileContents = utils.ReadFile(filePathGiven)
		}
	} else {
		stdInput, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Println("Failed to read from stdin")
		}

		fileContents = stdInput
	}

	if utils.HasArgs(argsGiven) {
		if argsGiven[0] == "-c" {
			// Count Total Bytes
			totalByteCount := countBytes(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[0] == "-l" {
			// Count Total Lines
			totalLineCount := countLines(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalLineCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[0] == "-w" {
			// Count Total Words
			totalWordCount := countWords(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalWordCount, filePathGiven)

			fmt.Println(formattedResult)
		} else if argsGiven[0] == "-m" {
		// Count Total Characters
			totalCharCount := countChars(fileContents)

			formattedResult := fmt.Sprintf("    %d %s", totalCharCount, filePathGiven)

			fmt.Println(formattedResult)
		} else {
			filePathGiven = argsGiven[0]

			totalByteCount := countBytes(fileContents)
			totalLineCount := countLines(fileContents)
			totalWordCount := countWords(fileContents)

			formattedResult := fmt.Sprintf("    %d   %d   %d %s", totalLineCount, totalWordCount, totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		}
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