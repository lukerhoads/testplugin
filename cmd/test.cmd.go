package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
)

type testCmd struct{}

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
	log.Println("Executing...")
	return nil
}

type testCmds struct{}

func (t testCmds) Init(ctx context.Context) error {
	log.Println("test cmd module loaded")
	return nil
}

func (t testCmds) Registry() map[string]plugintypes.Command {
	return map[string]plugintypes.Command{
		"command": testCmd{},
	}
}

var Commands testCmds

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"command": &plugintypes.CommandPlugin{Impl: testCmds{}},
		},
	})
}
