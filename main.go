package main

import (
	"fmt"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	tfplugin "github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Level:  hclog.Trace,
		Output: os.Stderr,
	})

	client := plugin.NewClient(&plugin.ClientConfig{
		Cmd:              exec.Command("sh", "-c", fmt.Sprintf("./plugin/%s", os.Getenv("PROVIDER_PLUGIN"))),
		HandshakeConfig:  tfplugin.Handshake,
		VersionedPlugins: tfplugin.VersionedPlugins,
		Managed:          true,
		Logger:           logger,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}

	raw, err := rpcClient.Dispense(tfplugin.ProviderPluginName)
	if err != nil {
		panic(err)
	}

	p := raw.(*tfplugin.GRPCProvider)
	p.PluginClient = client

	// Any actions needed for plugin
	resp := p.GetSchema()
	fmt.Println(resp)
}
