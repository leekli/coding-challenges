package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Print short-hand
var ptLn = fmt.Println
var ptF = fmt.Printf

func main() {
	// Retrieve any args given by the user not incl. file name
	userArgsGiven := os.Args[1:]

	if len(userArgsGiven) == 0 {
		ptLn(BuildNoArgsGivenMsg())

		return
	}

	// Get the file path/name given as the final argument and check it exists
	var filePath string

	if len(userArgsGiven) > 1 {
		filePath = userArgsGiven[len(userArgsGiven) - 1]
	}

	fileExists := checkFileExists(filePath)

	if !fileExists {
		ptF("File '%v' does not exist.", filePath)
		return
	}

	// Just deal with -f flag for now
	if strings.Contains(userArgsGiven[0], "-f") {
		reNumsOnly := regexp.MustCompile("[0-9]+")
		fNumsList := reNumsOnly.FindAllString(userArgsGiven[0], -1)
		fieldToPrint, err := strconv.ParseInt(fNumsList[0], 10, 32)

		if err == nil {	
			PrintBySpecifiedField(filePath, int(fieldToPrint))
		}
	}
	
}

func BuildNoArgsGivenMsg() string {
	msgToReturn := "usage: "

	msgToReturn += "	cut -f list"

	return msgToReturn
}

func checkFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    }

	return false
}

func PrintBySpecifiedField(filePath string, fieldNum int) {
	file, err := os.Open(filePath)

    if err != nil {
        ptF("Error opening file: %v\n", err)
        return
    }

    defer file.Close()

	scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

		// Currently defaults to tabs
        fields := strings.Split(line, "\t")

        if fieldNum > len(fields) {
            ptF("Line does not have %d fields: %s\n", fieldNum, line)
            continue
        }

        fmt.Println(fields[fieldNum - 1])
    }

    if err := scanner.Err(); err != nil {
        ptF("Error reading file: %v\n", err)
    }
}