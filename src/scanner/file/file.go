package file

import (
	"bufio"
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
		file.parser.Map(line)
	}
	return nil
}
