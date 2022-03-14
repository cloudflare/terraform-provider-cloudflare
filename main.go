package main

import (
	"flag"

	"github.com/cloudflare/terraform-provider-cloudflare/cloudflare"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: cloudflare.Provider,
		ProviderAddr: "registry.terraform.io/cloudflare/cloudflare",
		Debug:        debugMode,
	}

	plugin.Serve(opts)
}
