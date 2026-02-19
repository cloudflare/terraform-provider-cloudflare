package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareTeamsRuleSchema returns the source schema for legacy cloudflare_teams_rule resource.
// Schema version: 0 (implicit in v4 SDKv2 - no explicit Version field specified)
// Resource type: cloudflare_teams_rule
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_teams_rules.go
func SourceCloudflareTeamsRuleSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 SDKv2 schema has no explicit Version field (defaults to 0)
		Attributes: map[string]schema.Attribute{
			// Identity fields
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},

			// Core fields
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Required: true, // Required in v4, Optional in v5
			},
			"precedence": schema.Int64Attribute{
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
			},
			"action": schema.StringAttribute{
				Required: true,
			},
			"filters": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},

			// Expression fields
			"traffic": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"identity": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"device_posture": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},

			// Computed fields
			"version": schema.Int64Attribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			// rule_settings - TypeList MaxItems:1 in v4 SDKv2
			"rule_settings": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						// Scalar fields
						"block_page_enabled": schema.BoolAttribute{
							Optional: true,
						},
						"block_page_reason": schema.StringAttribute{
							Optional: true,
						},
						"override_ips": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"override_host": schema.StringAttribute{
							Optional: true,
						},
						"ip_categories": schema.BoolAttribute{
							Optional: true,
						},
						"ignore_cname_category_matches": schema.BoolAttribute{
							Optional: true,
						},
						"allow_child_bypass": schema.BoolAttribute{
							Optional: true,
						},
						"bypass_parent_rule": schema.BoolAttribute{
							Optional: true,
						},
						"insecure_disable_dnssec_validation": schema.BoolAttribute{
							Optional: true,
						},
						"resolve_dns_through_cloudflare": schema.BoolAttribute{
							Optional: true,
						},
						"add_headers": schema.MapAttribute{
							ElementType: types.StringType, // v4: Map[string]string
							Optional:    true,
						},
					},
					Blocks: map[string]schema.Block{
						// All nested structures - TypeList MaxItems:1 in v4 SDKv2
						"audit_ssh": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"command_logging": schema.BoolAttribute{
										Required: true,
									},
								},
							},
						},
						"l4override": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"ip": schema.StringAttribute{
										Required: true,
									},
									"port": schema.Int64Attribute{
										Required: true,
									},
								},
							},
						},
						"biso_admin_controls": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Optional: true,
									},
									// v1 fields (deprecated - but present in v4 state)
									"disable_printing": schema.BoolAttribute{
										Optional: true,
									},
									"disable_copy_paste": schema.BoolAttribute{
										Optional: true,
									},
									"disable_download": schema.BoolAttribute{
										Optional: true,
									},
									"disable_keyboard": schema.BoolAttribute{
										Optional: true,
									},
									"disable_upload": schema.BoolAttribute{
										Optional: true,
									},
									"disable_clipboard_redirection": schema.BoolAttribute{
										Optional: true,
									},
									// v2 fields
									"copy": schema.StringAttribute{
										Optional: true,
									},
									"download": schema.StringAttribute{
										Optional: true,
									},
									"keyboard": schema.StringAttribute{
										Optional: true,
									},
									"paste": schema.StringAttribute{
										Optional: true,
									},
									"printing": schema.StringAttribute{
										Optional: true,
									},
									"upload": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
						"check_session": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"enforce": schema.BoolAttribute{
										Required: true,
									},
									"duration": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
						"egress": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"ipv6": schema.StringAttribute{
										Required: true,
									},
									"ipv4": schema.StringAttribute{
										Required: true,
									},
									"ipv4_fallback": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
						"untrusted_cert": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
						"payload_log": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Required: true,
									},
								},
							},
						},
						"notification_settings": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Optional: true,
									},
									"message": schema.StringAttribute{
										Optional: true,
									},
									"support_url": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
						"dns_resolvers": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Blocks: map[string]schema.Block{
									"ipv4": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"ip": schema.StringAttribute{
													Required: true,
												},
												"port": schema.Int64Attribute{
													Optional: true,
												},
												"vnet_id": schema.StringAttribute{
													Optional: true,
												},
												"route_through_private_network": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
									"ipv6": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"ip": schema.StringAttribute{
													Required: true,
												},
												"port": schema.Int64Attribute{
													Optional: true,
												},
												"vnet_id": schema.StringAttribute{
													Optional: true,
												},
												"route_through_private_network": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"resolve_dns_internally": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"view_id": schema.StringAttribute{
										Optional: true,
									},
									"fallback": schema.StringAttribute{
										Optional: true,
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
