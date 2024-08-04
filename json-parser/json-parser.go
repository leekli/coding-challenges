package main

import (
	"errors"
	"fmt"
	"json-parser/utils"
)

func main() {
	// Change this line to be able take args eventually
	testFile := "test-data/task1-valid.json"

	// Read the provided JSON File (A check will be conducted to check the file exists)
	jsonData := utils.ReadFile(testFile)

	// Retrieve pre-existing list of valid tokens
	validTokensList := utils.GetValidTokensList()

	// Conduct 'Lexical Analysis' to ensure all characters are valid tokens & ensure no invalid tokens
	result, err := lexicalChecker(jsonData, validTokensList)
	utils.CheckError(err)

	if result {
		fmt.Println("Valid JSON!")
	}
}

func lexicalChecker(inputData []byte, validTokens []string) (bool, error) {
	for i := 0; i < len(inputData); i++ {
		if(!utils.CharIsValidToken(inputData[i], validTokens)) {
			return false, errors.New("invalid JSON")
		}
	}

	return true, nil
}