package main

import (
	"fmt"
	"io"
	"json-parser/lexer"
	"json-parser/parser"
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
		panic("❌ JSON Parser: No arguments given.")
	}

	if len(argsGiven) < 2 {
		panic("❌ JSON Parser: You must supply a file path.")
	}

	// Get the given file path from the user args & check it exists
	filePathGiven := argsGiven[1]

	if !CheckFileExists(filePathGiven) {
		panic("❌ JSON Parser: Given file does not exist, or file is not valid.")
	}

	// Read the given file, returned as string representation, check for error
	fileContents, err := ReadFileStream(filePathGiven)

	if err != nil {
		panic("❌ JSON Parser: There was an error with reading the file.")
	}
	
	// Pass to the Lexer
	tokenList := lexer.Lexer(fileContents)

	// Pass token list to the parser
	_, validJson := parser.Parse(tokenList)

	if validJson {
		fmt.Println("✅ The provided JSON file is VALID")
	} else {
		fmt.Println("❌ The provided JSON file is INVALID")
	}
}