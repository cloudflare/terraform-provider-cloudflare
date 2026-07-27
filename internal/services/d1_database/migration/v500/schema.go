package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareD1DatabaseSchema returns the legacy cloudflare_d1_database schema (schema_version=0).
// This is used by UpgradeFromLegacyV0 to parse state from the legacy framework provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/framework/service/d1/schema.go
func SourceCloudflareD1DatabaseSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"version": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
