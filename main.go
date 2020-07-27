package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cloudflare"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudflare.Provider})
}
