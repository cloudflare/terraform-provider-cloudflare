package main

import (
	"context"
	"flag"
	"log"

	"github.com/cloudflare/terraform-provider-cloudflare/cloudflare"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: cloudflare.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/cloudflare/cloudflare", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
