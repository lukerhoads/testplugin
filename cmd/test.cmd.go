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

type TestCommandMap struct{}

func (t *TestCommandMap) Commands() []string {
	return []string{"TestCmd"}
}

type TestCommand plugintypes.Command

func (t TestCommand) GetParentCommand() []string {
	return t.ParentCommand
}

func (t TestCommand) GetName() string {
	return t.Name
}

func (t TestCommand) GetUsage() string {
	return t.Usage
}

func (t TestCommand) GetShortDesc() string {
	return t.ShortDesc
}

func (t TestCommand) GetLongDesc() string {
	return t.LongDesc
}

func (t TestCommand) GetNumArgs() int {
	return t.NumArgs
}

func (TestCommand) Exec(cmd *cobra.Command, flags []string) error {
	log.Println("Executing...")
	return nil
}

func init() {
	// gob.Register(TestCommandMap{})
}

var TestCmd TestCommand = TestCommand{
	ParentCommand: []string{"starport", "chain", "serve"},
	Name:          "TestCommand",
	NumArgs:       0,
	Usage:         "test",
	ShortDesc:     "Short description",
	LongDesc:      "Long description",
}

func main() {
	commandMap := &TestCommandMap{}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Exiting...")
		os.Exit(0)
	}()

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugintypes.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"command_map": &plugintypes.CommandMapperPlugin{Impl: commandMap},
			"TestCmd":     &plugintypes.CommandModulePlugin{Impl: TestCmd},
		},
	})
}
