package main

import (
	"os"
	"scanner/file"
	"strconv"
)

func main() {
	isExitOnFail := true
	logFile := "example.txt"
	if len(os.Args) >= 2 {
		logFile = os.Args[1]
	}
	if len(os.Args) >= 3 {
		isExitOnFail, _ = strconv.ParseBool(os.Args[2])
	}
	file.Parse(logFile, isExitOnFail)
	return
}
