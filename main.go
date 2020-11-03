package main

import (
	"github.com/davidji99/terraform-provider-asana/asana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: asana.Provider})
}
