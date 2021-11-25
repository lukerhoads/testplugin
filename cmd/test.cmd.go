package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
)

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

func (testCmd) Exec(cmd *cobra.Command, flags []string) error {
	fmt.Println("Executing...")
	return nil
}

type testCmds struct{}

func (testCmds) Init(ctx context.Context) error {
	fmt.Println("test cmd module loaded")
	return nil
}

func (testCmds) Registry() map[string]plugintypes.Command {
	return map[string]plugintypes.Command{
		"command": testCmd("test"),
	}
}

var Commands testCmds

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"command": &plugintypes.CommandPlugin{Impl: testCmds{}},
		},
	})
}
