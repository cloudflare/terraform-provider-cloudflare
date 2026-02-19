package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceSpectrumApplicationSchema returns the source schema for legacy cloudflare_spectrum_application resource.
// Schema version: 1 (actual v4 provider schema version)
// Resource type: cloudflare_spectrum_application
//
// This minimal schema is used only for reading v4 state during migration.
// In v4 (SDKv2), dns, origin_dns, edge_ips, and origin_port_range were TypeList blocks
// with MaxItems:1, stored as arrays in state JSON. We define them as ListNestedBlock
// so the Plugin Framework can parse the v4 state correctly.
func SourceSpectrumApplicationSchema() schema.Schema {
	return schema.Schema{
		Version: 1, // Must match actual v4 schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"protocol": schema.StringAttribute{
				Required: true,
			},
			"origin_direct": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"origin_port": schema.Int64Attribute{
				Optional: true,
			},
			"argo_smart_routing": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"ip_firewall": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"proxy_protocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"tls": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"traffic_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			"dns": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"origin_dns": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
						"ttl": schema.Int64Attribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"edge_ips": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional: true,
						},
						"connectivity": schema.StringAttribute{
							Optional: true,
						},
						"ips": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"origin_port_range": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"start": schema.Int64Attribute{
							Optional: true,
						},
						"end": schema.Int64Attribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
