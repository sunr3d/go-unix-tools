package myXargs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type options struct {
	command     string
	commandArgs []string
}

func newOptions(args []string) (*options, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf(errNoCommand)
	}

	return &options{
		command:     args[1],
		commandArgs: args[2:],
	}, nil
}

func (o *options) run() error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputLine := scanner.Text()
		fullArgs := append(o.commandArgs, inputLine)

		cmd := exec.Command(o.command, fullArgs...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			return fmt.Errorf(errCmdExec, o.command, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(errStdinRead, err)
	}

	return nil
}
