package shell

import (
	"bufio"
	"fmt"
	"os"
	"parser"
)

type Shell struct {
	parser *parser.Parser
}

// Parse - Parse shell to command
func Parse() error {
	shell := new(Shell)
	shell.parser = parser.Init()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input:")
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Println("\nOutput:")
		shell.parser.Map(line)
		fmt.Println("")
	}
}
