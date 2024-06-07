package utils

import (
	"reflect"
	"testing"

	"github.com/cloudflare/cloudflare-go"
)

func TestUserAgentBuilding(t *testing.T) {
	tests := []struct {
		input  UserAgentBuilderParams
		expect string
	}{
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0")}, expect: "terraform-provider-cloudflare/1.0"},
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0"), PluginType: cloudflare.StringPtr("terraform-plugin-foo")}, expect: "terraform-provider-cloudflare/1.0 terraform-plugin-foo"},
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0"), PluginType: cloudflare.StringPtr("terraform-plugin-foo"), PluginVersion: cloudflare.StringPtr("1.2.3")}, expect: "terraform-provider-cloudflare/1.0 terraform-plugin-foo/1.2.3"},
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0"), PluginType: cloudflare.StringPtr("terraform-plugin-foo"), PluginVersion: cloudflare.StringPtr("1.2.3"), TerraformVersion: cloudflare.StringPtr("9.9.9")}, expect: "terraform-provider-cloudflare/1.0 terraform-plugin-foo/1.2.3 terraform/9.9.9"},
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0"), OperatorSuffix: cloudflare.StringPtr("example/v88")}, expect: "terraform-provider-cloudflare/1.0 example/v88"},
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0"), OperatorSuffix: cloudflare.StringPtr("example/v88"), TerraformVersion: cloudflare.StringPtr("1.2.3")}, expect: "terraform-provider-cloudflare/1.0 example/v88"},
		{input: UserAgentBuilderParams{ProviderVersion: cloudflare.StringPtr("1.0"), TerraformVersion: cloudflare.StringPtr("1.2.3")}, expect: "terraform-provider-cloudflare/1.0 terraform/1.2.3"},
	}

	for _, tc := range tests {
		got := BuildUserAgent(tc.input)
		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
