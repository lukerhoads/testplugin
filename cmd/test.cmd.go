package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
)

type TestCommand string

func (TestCommand) ParentCommand() []string {
	return []string{"starport", "chain", "serve"}
}

func (TestCommand) Name() string {
	return "TestCommand"
}

func (TestCommand) NumArgs() int {
	return 0
}

func (TestCommand) Usage() string {
	return "test"
}

func (TestCommand) ShortDesc() string {
	return "Short description"
}

func (TestCommand) LongDesc() string {
	return "Long description"
}

func (TestCommand) Exec(cmd *cobra.Command, flags []string) error {
	fmt.Println("Executing...")
	return nil
}

type TestCommands struct{}

func (TestCommands) Init(ctx context.Context) error {
	fmt.Println("test cmd module loaded")
	return nil
}

func (TestCommands) Registry() map[string]plugintypes.Command {
	return map[string]plugintypes.Command{
		"command": TestCommand("command"),
	}
}

func main() {
	commands := &TestCommands{}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"command": &plugintypes.CommandPlugin{Impl: commands},
		},
	})
}
