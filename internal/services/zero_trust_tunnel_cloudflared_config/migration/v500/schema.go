package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UnionV0Schema returns a permissive schema for version=0 that can parse both:
// - v4 SDKv2 state (config as list)
// - early v5 state (config as object)
//
// We model config as DynamicAttribute and detect/parse the exact format in the upgrader.
func UnionV0Schema(buildSchema func(context.Context) schema.Schema, ctx context.Context) *schema.Schema {
	s := buildSchema(ctx)
	s.Version = 0
	s.Attributes["config"] = schema.DynamicAttribute{
		Optional: true,
		Computed: true,
	}
	return &s
}

// SourceV4TunnelConfigSchema returns the schema for the v4 SDKv2 zero_trust_tunnel_cloudflared_config resource.
// Schema version: 0 (default SDKv2, no explicit SchemaVersion defined).
//
// Both cloudflare_tunnel_config (deprecated) and cloudflare_zero_trust_tunnel_cloudflared_config
// use the same underlying schema in v4.
//
// CRITICAL: TypeList MaxItems:1 blocks are represented as schema.ListNestedAttribute (not SingleNestedAttribute).
// This is because SDKv2 stores MaxItems:1 blocks as JSON arrays in state.
// TypeSet of objects (ip_rules) is also represented as schema.ListNestedAttribute to parse the raw state.
// TypeSet of strings (aud_tag) is schema.SetAttribute{ElementType: types.StringType}.
func SourceV4TunnelConfigSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"tunnel_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			// TypeList MaxItems:1 → schema.ListNestedAttribute
			"config": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// warp_routing: TypeList MaxItems:1 → removed in v5
						"warp_routing": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
						// origin_request: TypeList MaxItems:1
						"origin_request": schema.ListNestedAttribute{
							Optional:     true,
							NestedObject: originRequestNestedObject(),
						},
						// ingress_rule: TypeList (renamed to ingress in v5)
						"ingress_rule": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"hostname": schema.StringAttribute{
										Optional: true,
									},
									"path": schema.StringAttribute{
										Optional: true,
									},
									"service": schema.StringAttribute{
										Required: true,
									},
									// origin_request per ingress: TypeList MaxItems:1
									"origin_request": schema.ListNestedAttribute{
										Optional:     true,
										NestedObject: originRequestNestedObject(),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// originRequestNestedObject returns the NestedAttributeObject for origin_request blocks.
// Used at both config-level and ingress-level origin_request.
//
// Duration fields (connect_timeout, tls_timeout, tcp_keep_alive, keep_alive_timeout) are
// stored as strings in v4 state (e.g., "30s", "10s", "1m30s").
// ip_rules is TypeSet of objects → schema.ListNestedAttribute (must parse, will be dropped).
// access is TypeList MaxItems:1 → schema.ListNestedAttribute.
// aud_tag is TypeSet of TypeString → schema.SetAttribute{ElementType: types.StringType}.
func originRequestNestedObject() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			// Duration fields: stored as strings in v4
			"connect_timeout": schema.StringAttribute{
				Optional: true,
			},
			"tls_timeout": schema.StringAttribute{
				Optional: true,
			},
			"tcp_keep_alive": schema.StringAttribute{
				Optional: true,
			},
			"keep_alive_timeout": schema.StringAttribute{
				Optional: true,
			},
			// Int fields
			"keep_alive_connections": schema.Int64Attribute{
				Optional: true,
			},
			// Bool fields
			"no_happy_eyeballs": schema.BoolAttribute{
				Optional: true,
			},
			"no_tls_verify": schema.BoolAttribute{
				Optional: true,
			},
			"disable_chunked_encoding": schema.BoolAttribute{
				Optional: true,
			},
			"http2_origin": schema.BoolAttribute{
				Optional: true,
			},
			// String fields
			"http_host_header": schema.StringAttribute{
				Optional: true,
			},
			"origin_server_name": schema.StringAttribute{
				Optional: true,
			},
			"ca_pool": schema.StringAttribute{
				Optional: true,
			},
			"proxy_type": schema.StringAttribute{
				Optional: true,
			},
			// Removed in v5 (must parse but will be dropped)
			"bastion_mode": schema.BoolAttribute{
				Optional: true,
			},
			"proxy_address": schema.StringAttribute{
				Optional: true,
			},
			"proxy_port": schema.Int64Attribute{
				Optional: true,
			},
			// ip_rules: TypeSet of objects → ListNestedAttribute (must parse, will be dropped)
			"ip_rules": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							Optional: true,
						},
						"ports": schema.ListAttribute{
							ElementType: types.Int64Type,
							Optional:    true,
						},
						"allow": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			// access: TypeList MaxItems:1 → schema.ListNestedAttribute
			"access": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"required": schema.BoolAttribute{
							Optional: true,
						},
						"team_name": schema.StringAttribute{
							Optional: true,
						},
						// aud_tag: TypeSet of TypeString → schema.SetAttribute
						"aud_tag": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}
