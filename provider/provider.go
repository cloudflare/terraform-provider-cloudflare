// Package provider exposes the Cloudflare Terraform provider implementation so
// that it can be embedded in other Go programs.
//
// The implementation itself lives under internal/, which Go's import rules make
// unreachable from outside this module. Tools that host the provider in-process
// rather than running it as a plugin binary -- code generators, test harnesses,
// and Crossplane providers built with upjet -- therefore have no way to obtain
// it. This package is a thin, dependency-free re-export that closes that gap
// without moving any of the implementation out of internal/.
package provider

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
)

// New returns a factory for the Cloudflare Terraform provider, in the shape
// expected by providerserver.Serve. The version is reported to Terraform as the
// provider version.
func New(version string) func() fwprovider.Provider {
	return internal.NewProvider(version)
}

// NewProvider returns the Cloudflare Terraform provider directly, for callers
// that need the provider itself rather than a factory.
func NewProvider(version string) fwprovider.Provider {
	return internal.NewProvider(version)()
}
