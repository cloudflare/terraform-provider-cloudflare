package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareRecordSchema returns the legacy cloudflare_record schema (schema_version=3).
// This is used by MoveState and UpgradeFromLegacyV3 to parse state from the legacy SDKv2 provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_record.go
func SourceCloudflareRecordSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			// Source uses "value", target uses "content"
			"value": schema.StringAttribute{
				Optional: true,
			},
			// Source also has content (used by API response)
			"content": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"priority": schema.Int64Attribute{
				Optional: true,
			},
			"proxied": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional: true,
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			// Deprecated/removed in target
			"allow_overwrite": schema.BoolAttribute{
				Optional: true,
			},
			"hostname": schema.StringAttribute{
				Computed: true,
			},
			"proxiable": schema.BoolAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"metadata": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
		},
		Blocks: map[string]schema.Block{
			// Source data is a list block with MaxItems: 1
			// Target data is a SingleNestedAttribute
			"data": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						// CAA fields
						"flags": schema.StringAttribute{
							Optional: true,
						},
						"tag": schema.StringAttribute{
							Optional: true,
						},
						// Source CAA uses "content", target uses "value"
						"content": schema.StringAttribute{
							Optional: true,
						},

						// SRV fields
						"service": schema.StringAttribute{
							Optional: true,
						},
						"proto": schema.StringAttribute{
							Optional: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
						"priority": schema.Int64Attribute{
							Optional: true,
						},
						"weight": schema.Int64Attribute{
							Optional: true,
						},
						"port": schema.Int64Attribute{
							Optional: true,
						},
						"target": schema.StringAttribute{
							Optional: true,
						},

						// DNSKEY/DS/CERT fields
						"algorithm": schema.Int64Attribute{
							Optional: true,
						},
						"key_tag": schema.Int64Attribute{
							Optional: true,
						},
						"type": schema.Int64Attribute{
							Optional: true,
						},
						"protocol": schema.Int64Attribute{
							Optional: true,
						},
						"public_key": schema.StringAttribute{
							Optional: true,
						},
						"digest": schema.StringAttribute{
							Optional: true,
						},
						"digest_type": schema.Int64Attribute{
							Optional: true,
						},
						"certificate": schema.StringAttribute{
							Optional: true,
						},

						// TLSA fields
						"usage": schema.Int64Attribute{
							Optional: true,
						},
						"selector": schema.Int64Attribute{
							Optional: true,
						},
						"matching_type": schema.Int64Attribute{
							Optional: true,
						},

						// LOC fields
						"altitude": schema.Float64Attribute{
							Optional: true,
						},
						"lat_degrees": schema.Int64Attribute{
							Optional: true,
						},
						"lat_direction": schema.StringAttribute{
							Optional: true,
						},
						"lat_minutes": schema.Int64Attribute{
							Optional: true,
						},
						"lat_seconds": schema.Float64Attribute{
							Optional: true,
						},
						"long_degrees": schema.Int64Attribute{
							Optional: true,
						},
						"long_direction": schema.StringAttribute{
							Optional: true,
						},
						"long_minutes": schema.Int64Attribute{
							Optional: true,
						},
						"long_seconds": schema.Float64Attribute{
							Optional: true,
						},
						"precision_horz": schema.Float64Attribute{
							Optional: true,
						},
						"precision_vert": schema.Float64Attribute{
							Optional: true,
						},
						"size": schema.Float64Attribute{
							Optional: true,
						},

						// NAPTR fields
						"order": schema.Int64Attribute{
							Optional: true,
						},
						"preference": schema.Int64Attribute{
							Optional: true,
						},
						"regex": schema.StringAttribute{
							Optional: true,
						},
						"replacement": schema.StringAttribute{
							Optional: true,
						},

						// SSHFP fields
						"fingerprint": schema.StringAttribute{
							Optional: true,
						},

						// URI fields (uses content for value)
						"value": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
