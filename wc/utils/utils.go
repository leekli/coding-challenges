package utils

import (
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func HasArgs(args []string) bool {
	return len(args) > 0
}

func CheckFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    } else {
        return false
    }
}

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)

	CheckError(err)

	return data
}