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
		if argsGiven[0] == "-c" {
			totalByteCount := countBytes(filePathGiven)

			formattedResult := fmt.Sprintf("  %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		}
	}
}

func countBytes(filePath string) int {
	fileContents := utils.ReadFile(filePath)

	totalBytes := len(fileContents)

	return totalBytes
}