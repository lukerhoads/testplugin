package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
)

type TestHookMap struct{}

func (t *TestHookMap) Hooks() []string {
	return []string{"TestHk"}
}

type TestHook plugintypes.Hook

func (t TestHook) GetParentCommand() []string {
	return t.ParentCommand
}

func (t TestHook) GetName() string {
	return t.Name
}

func (t TestHook) GetType() string {
	return t.HookType
}

func (TestHook) PreRun(cmd *cobra.Command, flags []string) error {
	log.Println("Executing pre run...")
	return nil
}

func (TestHook) PostRun(cmd *cobra.Command, flags []string) error {
	return nil
}

func init() {
	// gob.Register(TestHook{})
}

var TestHk TestHook = TestHook{
	ParentCommand: []string{"starport", "chain", "serve", "test"},
	Name:          "TestHook",
	HookType:      "pre",
}

func main() {
	hooks := &TestHookMap{}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Exiting...")
		os.Exit(0)
	}()

	// time this out maybe?
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"hook_map": &plugintypes.HookMapperPlugin{Impl: hooks},
			"TestHk":   &plugintypes.HookModulePlugin{Impl: TestHk},
		},
	})
}
