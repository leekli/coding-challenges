package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Print short-hand globals
var ptLn = fmt.Println
var ptF = fmt.Printf

func main() {
	// Retrieve any args given by the user not incl. file name
	userArgsGiven := os.Args[1:]

	// If no args given, display the usage help screen
	if len(userArgsGiven) == 0 {
		ptLn(BuildNoArgsGivenMsg())

		return
	}

	// Build variables required to print what the user requested

	// Get the file path/name given as the final argument and check it exists
	var filePath string

	if len(userArgsGiven) > 1 {
		filePath = userArgsGiven[len(userArgsGiven) - 1]
	}

	fileExists := CheckFileExists(filePath)

	if !fileExists {
		ptF("File '%v' does not exist.", filePath)
		return
	}

	var fieldToPrint int
	var delimiterToUse string

	// The -f flag
	// -f list 		The list specifies fields, separated in the input by the field delimiter character (see the -d option).  Output fields are separated by a single occurrence of the field delimiter character.
	if strings.Contains(userArgsGiven[0], "-f") {
		reNumsOnly := regexp.MustCompile("[0-9]+")
		fNumsList := reNumsOnly.FindAllString(userArgsGiven[0], -1)
		fieldNum, err := strconv.ParseInt(fNumsList[0], 10, 32)

		if err != nil {	
			ptLn("There was an error with the field number provided with the -f flag")
			return
		}

		fieldToPrint = int(fieldNum)
	}

	// The -d flag
	// -d delim 	Use delim as the field delimiter character instead of the tab character.
	if strings.Contains(userArgsGiven[1], "-d") {
		reDelimChar := regexp.MustCompile("^-d(.)$")
		dDelimCharMatch := reDelimChar.FindStringSubmatch(userArgsGiven[1])

		if len(dDelimCharMatch) > 1 {
			delimiterToUse = dDelimCharMatch[1]
		} else {
			delimiterToUse = "\t"
		}
	}

	PrintBySpecifiedField(filePath, fieldToPrint, delimiterToUse)
}

func BuildNoArgsGivenMsg() string {
	msgToReturn := "usage: "

	msgToReturn += "	cut -f list [-d delim]"

	return msgToReturn
}

func CheckFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    }

	return false
}

func PrintBySpecifiedField(filePath string, fieldNum int, delimiterChar string) int {
	if fieldNum == 0 {
		ptF("cut: values may not include zero")

        return -1
	}

	file, err := os.Open(filePath)

    if err != nil {
        ptF("Error opening file: %v\n", err)

        return -1
    }

    defer file.Close()

	scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        fields := strings.Split(line, delimiterChar)

        if fieldNum > len(fields) {
            ptLn(" ")
            continue
        }

        fmt.Println(fields[fieldNum - 1])
    }

    if err := scanner.Err(); err != nil {
        ptF("Error reading file: %v\n", err)
    }

	return 0
}