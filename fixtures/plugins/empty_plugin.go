package main

import "github.com/liamawhite/cli-with-i18n/plugin"

type EmptyPlugin struct{}

func (c *EmptyPlugin) Run(cliConnection plugin.CliConnection, args []string) {}

func (c *EmptyPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:     "EmptyPlugin",
		Commands: []plugin.Command{},
	}
}

func main() {
	plugin.Start(new(EmptyPlugin))
}
