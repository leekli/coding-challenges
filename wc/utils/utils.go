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

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)

	CheckError(err)

	return data
}