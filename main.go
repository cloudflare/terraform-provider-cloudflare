// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
  "context"
  "flag"
  "log"

  "github.com/cloudflare/terraform-provider-cloudflare/internal"
  "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
  version string = "dev"
)

func main() {
  var debug bool

  flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
  flag.Parse()

  opts := providerserver.ServeOpts{
    Address: "registry.terraform.io/cloudflare/cloudflare",
    Debug:   debug,
  }

  err := providerserver.Serve(context.Background(), internal.NewProvider(version), opts)

  if err != nil {
    log.Fatal(err.Error())
  }
}
