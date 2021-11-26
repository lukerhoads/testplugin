package main

import (
	"context"
	"encoding/gob"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
)

type TestHook struct{}

func (TestHook) ParentCommand() []string {
	return []string{"starport", "chain", "serve"}
}

func (TestHook) Name() string {
	return "TestCommand"
}

func (TestHook) Type() string {
	return "PreRun"
}

func (TestHook) ShortDesc() string {
	return "Short description"
}

func (TestHook) LongDesc() string {
	return "Long description"
}

func (TestHook) PreRun(cmd *cobra.Command, flags []string) error {
	fmt.Println("Executing pre run...")
	return nil
}

func (TestHook) PostRun(cmd *cobra.Command, flags []string) error {
	return nil
}

type TestHooks struct{}

func (TestHooks) Init(ctx context.Context) error {
	fmt.Println("test hook module loaded")
	return nil
}

func (TestHooks) Registry() map[string]plugintypes.Hook {
	return map[string]plugintypes.Hook{
		"hook": TestHook{},
	}
}

func main() {
	gob.Register(TestHook{})
	hooks := &TestHooks{}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"hook": &plugintypes.HookPlugin{Impl: hooks},
		},
	})
}
