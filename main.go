package main

import (
	"flag"
	"log"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	version string = "dev"
	commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: provider.New(version),
		ProviderAddr: "registry.terraform.io/cloudflare/cloudflare",
		Debug:        debugMode,
	}

	logFlags := log.Flags()
	logFlags = logFlags &^ (log.Ldate | log.Ltime)
	log.SetFlags(logFlags)

	plugin.Serve(opts)
}
