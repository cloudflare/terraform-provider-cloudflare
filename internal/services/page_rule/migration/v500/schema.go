package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceV4PageRuleSchema returns the schema for v4 page_rule resources (SDKv2 format).
// This schema represents the state structure from the v4 provider (schema_version=0).
func SourceV4PageRuleSchema() schema.Schema {
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
			"created_on": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
			},
			"modified_on": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
			},
			// In v4 SDKv2, TypeList with MaxItems:1 is stored as an array
			"actions": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// Boolean fields
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
						// String fields (on/off)
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
						// browser_cache_ttl is STRING in v4
						"browser_cache_ttl": schema.StringAttribute{
							Optional: true,
						},
						// Numeric fields
						"edge_cache_ttl": schema.Int64Attribute{
							Optional: true,
						},
						// Nested structures (TypeList MaxItems:1 in v4)
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
						"cache_key_fields": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cookie": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"check_presence": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"include": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
											},
										},
									},
									"header": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"check_presence": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"include": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"exclude": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
											},
										},
									},
									"host": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"resolved": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
									"query_string": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"include": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"exclude": schema.SetAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"ignore": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
									"user": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"device_type": schema.BoolAttribute{
													Optional: true,
												},
												"geo": schema.BoolAttribute{
													Optional: true,
												},
												"lang": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
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
