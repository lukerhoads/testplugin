package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
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

func (testHooks) Registry() map[string]plugintypes.Hook {
	return map[string]plugintypes.Hook{
		"hook": testHook("test"),
	}
}

var Commands testHooks

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"hook": &plugintypes.HookPlugin{Impl: testHooks{}},
		},
	})
}
