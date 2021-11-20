package main

import (
	"context"
	"fmt"

	"github.com/lukerhoads/starport/starport/services/plugins"
	"github.com/spf13/cobra"
)

type testCmd string

func (testCmd) ParentCommand() []string {
	return []string{"starport", "chain", "serve"}
}

func (testCmd) Name() string {
	return "TestCommand"
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

func (exampleCmds) Registry() map[string]plugins.Command {
	return map[string]plugins.Command{
		"test": testCmd("test"),
	}
}

var Commands exampleCmds
