package file

import (
	"bufio"
	"def/log"
	"fmt"
	"os"
	"parser"
)

type File struct {
	parser              *parser.Parser
	hostSuspiciousCount map[string]int
}

// Parse - Parse file to command
func Parse(fileName string) error {
	file := new(File)
	file.parser = parser.Init()
	file.hostSuspiciousCount = make(map[string]int)

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
		if file.isPostBeforePut(log, file.parser.Logs) {
			fmt.Printf("POST before PUT, %s\n", line)
		}
		if file.isSuspiciousActivity(log) {
			fmt.Printf("Is suspicious, %s\n", line)
		}
	}
	return nil
}

func (f *File) isResponseBig(log *log.ApacheLog) bool {
	if log.SizeByte > 100000 {
		return true
	}
	return false
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

func (f *File) isPostBeforePut(log *log.ApacheLog, allLogs []*log.ApacheLog) bool {
	if log.ReqMethod != "POST" {
		return false
	}
	for i := len(allLogs) - 2; i >= 0; i-- {
		prevLog := allLogs[i]
		if prevLog.RemoteHost != log.RemoteHost {
			continue
		}
		if prevLog.ReqResource != log.ReqResource {
			continue
		}
		if prevLog.ReqMethod == "POST" {
			return true
		}
		if prevLog.ReqMethod == "PUT" {
			return false
		}
	}
	// Could also means no PUT Found before POST
	return true
}

func (f *File) isSuspiciousActivity(log *log.ApacheLog) bool {
	if log.StatusCode != "401" {
		return false
	}
	f.hostSuspiciousCount[log.RemoteHost]++
	if f.hostSuspiciousCount[log.RemoteHost] >= 5 {
		return true
	}
	return false
}
