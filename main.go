package main

import (
	"github.com/JacobGsell/terraform-provider-shell/shell"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shell.Provider,
	})
}
