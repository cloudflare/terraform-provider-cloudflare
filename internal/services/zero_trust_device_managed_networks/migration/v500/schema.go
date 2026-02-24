package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareDeviceManagedNetworksSchema returns the legacy cloudflare_device_managed_networks schema (schema_version=0).
// This is used by MoveState and UpgradeFromLegacyV0 to parse state from the legacy SDKv2 provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_device_managed_networks.go
func SourceCloudflareDeviceManagedNetworksSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
		Blocks: map[string]schema.Block{
			// Source config is a list block with MaxItems: 1 (TypeList in SDKv2)
			// Target config is a SingleNestedAttribute
			"config": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"tls_sockaddr": schema.StringAttribute{
							Required: true,
						},
						"sha256": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}
