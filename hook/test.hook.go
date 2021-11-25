package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/spf13/cobra"
	"github.com/tendermint/starport/starport/services/pluginsrpc"
)

type testHook string

func (testHook) ParentCommand() []string {
	return []string{"starport", "chain", "serve"}
}

func (testHook) Name() string {
	return "TestCommand"
}

func (testHook) Type() string {
	return "PreRun"
}

func (testHook) ShortDesc() string {
	return "Short description"
}

func (testHook) LongDesc() string {
	return "Long description"
}

func (testHook) PreRun(cmd *cobra.Command, flags []string) error {
	fmt.Println("Executing pre run...")
	return nil
}

func (testHook) PostRun(cmd *cobra.Command, flags []string) error {
	return nil
}

type testHooks struct{}

func (testHooks) Init(ctx context.Context) error {
	fmt.Println("test hook module loaded")
	return nil
}

func (testHooks) Registry() map[string]pluginsrpc.Hook {
	return map[string]pluginsrpc.Hook{
		"test": testHook("test"),
	}
}

var Commands testHooks

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginsrpc.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"test": &pluginsrpc.CommandPlugin{Impl: &testHooks{}},
		},
	})
}
