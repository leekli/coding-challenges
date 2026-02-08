package main

import (
	"io"
	"json-parser/lexer"
	"log"
	"os"
)

func CheckFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    }
	
	return false
}

func ReadFileStream(filePath string) (string, error) {
	data, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer data.Close()

	bytes, err := io.ReadAll(data)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {
	// Retrieve any args given by the user
	argsGiven := os.Args

	if argsGiven == nil {
		log.Fatalf("❌ JSON Parser: No arguments given.")
	}

	if len(argsGiven) < 2 {
		log.Fatalf("❌ JSON Parser: You must supply a file path.")
	}

	// Get the given file path from the user args & check it exists
	filePathGiven := argsGiven[1]

	if !CheckFileExists(filePathGiven) {
		log.Fatalf("❌ JSON Parser: Given file does not exist, or file is not valid.")
	}

	// Read the given file, returned as string representation, check for error
	fileContents, err := ReadFileStream(filePathGiven)

	if err != nil {
		log.Fatalf("❌ JSON Parser: There was an error with reading the file.")
	}
	
	// Pass to the Lexer
	lexer.Lexer(fileContents)
}