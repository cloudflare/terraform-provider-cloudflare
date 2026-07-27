package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareHyperdriveConfigSchema returns the legacy cloudflare_hyperdrive_config schema (schema_version=0).
// This is used by UpgradeFromV0 to parse state from the legacy framework provider v4.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/framework/service/hyperdrive_config/schema.go
func SourceCloudflareHyperdriveConfigSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"origin": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Required: true,
					},
					"password": schema.StringAttribute{
						Required:  true,
						Sensitive: true,
					},
					"host": schema.StringAttribute{
						Required: true,
					},
					"port": schema.Int64Attribute{
						Optional: true,
					},
					"scheme": schema.StringAttribute{
						Required: true,
					},
					"user": schema.StringAttribute{
						Required: true,
					},
					"access_client_id": schema.StringAttribute{
						Optional: true,
					},
					"access_client_secret": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
					},
				},
			},
			"caching": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"disabled": schema.BoolAttribute{
						Optional: true,
						Computed: true,
					},
					"max_age": schema.Int64Attribute{
						Optional: true,
					},
					"stale_while_revalidate": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
		},
	}
}

