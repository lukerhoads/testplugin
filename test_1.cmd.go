package main

import (
	"context"
	"fmt"

	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
)

func main() {}

type testCmd string

func (testCmd) ParentCommand() []string {
	return []string{"starport", "chain", "serve"}
}

func (testCmd) Name() string {
	return "TestCommand"
}

func (testCmd) NumArgs() int {
	return 0
}

func (testCmd) Usage() string {
	return "test"
}

func (testCmd) ShortDesc() string {
	return "Short description"
}

func (testCmd) LongDesc() string {
	return "Long description"
}

func (testCmd) Exec(*cobra.Command, []string) error {
	fmt.Println("Executing...")
	return nil
}

type exampleCmds struct{}

func (exampleCmds) Init(ctx context.Context) error {
	fmt.Println("test module loaded")
	return nil
}

func (exampleCmds) Registry() map[string]plugintypes.Command {
	return map[string]plugintypes.Command{
		"test": testCmd("test"),
	}
}

var Commands exampleCmds
