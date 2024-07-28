package main

import (
	"fmt"
	"os"
	"wc/utils"
)

func main() {
	argsGiven := os.Args[1:]

	if !utils.HasArgs(argsGiven) {
		panic("No arguments given")
	}

	filePathGiven := argsGiven[len(argsGiven)-1]

	if utils.HasArgs(argsGiven) {
		// Count Bytes
		if argsGiven[0] == "-c" {
			totalByteCount := countBytes(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		}

		// Count Lines
		if argsGiven[0] == "-l" {
			totalLineCount := countLines(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalLineCount, filePathGiven)

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