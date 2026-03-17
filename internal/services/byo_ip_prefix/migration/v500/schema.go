package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareByoIPPrefixSchema returns the legacy cloudflare_byo_ip_prefix schema (schema_version=0).
// This is used by UpgradeFromV0 to parse state from the legacy SDKv2 provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_byo_ip_prefix.go
func SourceCloudflareByoIPPrefixSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"prefix_id": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"advertisement": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}
