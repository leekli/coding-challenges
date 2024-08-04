package utils

import (
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filePath string) []byte {
	checkFileExists(filePath)
	
	data, err := os.ReadFile(filePath)

	CheckError(err)

	return data
}

func checkFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    } else {
        return false
    }
}

func GetValidTokensList() []string {
	tokensList := []string{"{", "}"}

	return tokensList
}

func CharIsValidToken(charNum byte, tokenList []string) bool {
	charNumToStr := string(charNum)

	for i := 0; i < len(tokenList); i++ {
		if tokenList[i] == charNumToStr {
			return true
		}
	}

	return false
 }