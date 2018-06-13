package file

import (
	"bufio"
	"def/log"
	"fmt"
	"os"
	"parser"
)

type File struct {
	parser *parser.Parser
}

// Parse - Parse file to command
func Parse(fileName string) error {
	file := new(File)
	file.parser = parser.Init()

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		log, err := file.parser.Map(line)
		if err != nil {
			continue
		}
		if file.isResponseBig(log) {
			fmt.Printf("Response is too big, %s\n", line)
		}
		if file.isFailToFeedRover(log) {
			fmt.Printf("Failed to feed Rover, %s\n", line)
		}
	}
	return nil
}

func (f *File) isResponseBig(log *log.ApacheLog) bool {
	if log.SizeByte > 100000 {
		return false
	}
	return true
}

func (f *File) isFailToFeedRover(log *log.ApacheLog) bool {
	lenHost := len(log.RemoteHost)
	if lenHost < 12 {
		return false
	}
	if log.RemoteHost[lenHost-12:] != ".example.org" {
		return false
	}
	if log.StatusCode != "401" {
		return false
	}
	return true
}
