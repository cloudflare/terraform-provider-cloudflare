package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FlexibleV0Schema returns a minimal schema for version 0 that doesn't validate actions.
// This allows the handler to work with raw state and detect whether actions is
// an array (v4) or object (early v5) without Terraform's schema validation failing.
//
// By not including actions in the schema, Terraform won't try to parse/validate it,
// allowing us to handle both formats in the UpgradeFromV0 handler.
func FlexibleV0Schema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"target": schema.StringAttribute{
				Required: true,
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			// NOTE: actions field intentionally omitted from schema
			// The handler will access it via raw state to detect format
		},
	}
}

// SourceCloudflarePageRuleSchema returns the source schema for legacy cloudflare_page_rule resource.
// Schema version: 3 (standard for SDKv2 resources without explicit version)
// Resource type: cloudflare_page_rule
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// IMPORTANT: This schema represents v4 state structure (SDKv2):
// - TypeList MaxItems:1 blocks are stored as arrays in state
// - TypeSet fields are stored as arrays in state
// - All fields must match v4 attribute names, even if renamed/deprecated in v5
func SourceCloudflarePageRuleSchema() schema.Schema {
	return schema.Schema{
		Version: 3, // SDKv2 standard version
		Attributes: map[string]schema.Attribute{
			// Top-level fields
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"target": schema.StringAttribute{
				Required: true,
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},

			// actions: TypeList MaxItems:1 in v4 (stored as array in state)
			"actions": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// Boolean fields (with default: false in v4)
						"always_use_https": schema.BoolAttribute{
							Optional: true,
						},
						"disable_apps": schema.BoolAttribute{
							Optional: true,
						},
						"disable_performance": schema.BoolAttribute{
							Optional: true,
						},
						"disable_railgun": schema.BoolAttribute{
							Optional: true,
						},
						"disable_security": schema.BoolAttribute{
							Optional: true,
						},
						"disable_zaraz": schema.BoolAttribute{
							Optional: true,
						},

						// String fields (on/off values)
						"automatic_https_rewrites": schema.StringAttribute{
							Optional: true,
						},
						"browser_check": schema.StringAttribute{
							Optional: true,
						},
						"cache_by_device_type": schema.StringAttribute{
							Optional: true,
						},
						"cache_deception_armor": schema.StringAttribute{
							Optional: true,
						},
						"email_obfuscation": schema.StringAttribute{
							Optional: true,
						},
						"explicit_cache_control": schema.StringAttribute{
							Optional: true,
						},
						"ip_geolocation": schema.StringAttribute{
							Optional: true,
						},
						"mirage": schema.StringAttribute{
							Optional: true,
						},
						"opportunistic_encryption": schema.StringAttribute{
							Optional: true,
						},
						"origin_error_page_pass_thru": schema.StringAttribute{
							Optional: true,
						},
						"respect_strong_etag": schema.StringAttribute{
							Optional: true,
						},
						"response_buffering": schema.StringAttribute{
							Optional: true,
						},
						"rocket_loader": schema.StringAttribute{
							Optional: true,
						},
						"server_side_exclude": schema.StringAttribute{
							Optional: true,
						},
						"sort_query_string_for_cache": schema.StringAttribute{
							Optional: true,
						},
						"true_client_ip_header": schema.StringAttribute{
							Optional: true,
						},
						"waf": schema.StringAttribute{
							Optional: true,
						},

						// String fields (other)
						"bypass_cache_on_cookie": schema.StringAttribute{
							Optional: true,
						},
						"cache_level": schema.StringAttribute{
							Optional: true,
						},
						"cache_on_cookie": schema.StringAttribute{
							Optional: true,
						},
						"host_header_override": schema.StringAttribute{
							Optional: true,
						},
						"polish": schema.StringAttribute{
							Optional: true,
						},
						"resolve_override": schema.StringAttribute{
							Optional: true,
						},
						"security_level": schema.StringAttribute{
							Optional: true,
						},
						"ssl": schema.StringAttribute{
							Optional: true,
						},

						// browser_cache_ttl: STRING in v4 (converts to Int64 in v5)
						"browser_cache_ttl": schema.StringAttribute{
							Optional: true,
						},

						// Numeric fields
						"edge_cache_ttl": schema.Int64Attribute{
							Optional: true,
						},

						// forwarding_url: TypeList MaxItems:1 in v4
						"forwarding_url": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"url": schema.StringAttribute{
										Required: true,
									},
									"status_code": schema.Int64Attribute{
										Required: true,
									},
								},
							},
						},

						// minify: TypeList MaxItems:1 in v4 (DEPRECATED - will be dropped)
						"minify": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"js": schema.StringAttribute{
										Required: true,
									},
									"css": schema.StringAttribute{
										Required: true,
									},
									"html": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},

						// cache_key_fields: TypeList MaxItems:1 (5 levels deep!)
						"cache_key_fields": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// cookie: TypeList MaxItems:1
									"cookie": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"check_presence": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
												"include": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},

									// header: TypeList MaxItems:1
									"header": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"check_presence": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
												"include": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
												"exclude": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},

									// host: TypeList MaxItems:1
									"host": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"resolved": schema.BoolAttribute{
													Optional: true,
													Computed: true,
												},
											},
										},
									},

									// query_string: TypeList MaxItems:1
									"query_string": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"include": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
												"exclude": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
													Computed:    true,
												},
												"ignore": schema.BoolAttribute{
													Optional: true,
													Computed: true,
												},
											},
										},
									},

									// user: TypeList MaxItems:1
									"user": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"device_type": schema.BoolAttribute{
													Optional: true,
													Computed: true,
												},
												"geo": schema.BoolAttribute{
													Optional: true,
													Computed: true,
												},
												// lang: May not exist in v4 state!
												"lang": schema.BoolAttribute{
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						// cache_ttl_by_status: TypeSet in v4 (array of objects in state)
						"cache_ttl_by_status": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"codes": schema.StringAttribute{
										Required: true,
									},
									"ttl": schema.Int64Attribute{
										Required: true,
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
