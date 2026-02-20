package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceListSchema returns a combined schema that can deserialize both v4 (SDKv2)
// and v5 (Framework) cloudflare_list state.
//
// The published v5 provider (5.x) writes schema_version=0 (no Version field set),
// which is the same as v4. This means both v4 and v5 state arrive at the version 0
// upgrader. The combined schema includes:
//   - v4 "item" blocks with nested "value" sub-blocks (ListNestedBlock)
//   - v5 "items" set attribute with flat structure (SetNestedAttribute)
//   - v5 computed fields: created_on, modified_on, num_referencing_filters
//
// The handler detects which format is present and handles accordingly.
func SourceListSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"kind": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"num_items": schema.Float64Attribute{
				Computed: true,
			},
			// v5 computed fields
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"num_referencing_filters": schema.Float64Attribute{
				Computed: true,
			},
			// v5 items attribute (flat structure, no CustomType)
			"items": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"asn": schema.Int64Attribute{
							Optional: true,
						},
						"comment": schema.StringAttribute{
							Optional: true,
						},
						"ip": schema.StringAttribute{
							Optional: true,
						},
						"hostname": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"url_hostname": schema.StringAttribute{
									Required: true,
								},
								"exclude_exact_hostname": schema.BoolAttribute{
									Optional: true,
								},
							},
						},
						"redirect": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"source_url": schema.StringAttribute{
									Required: true,
								},
								"target_url": schema.StringAttribute{
									Required: true,
								},
								"status_code": schema.Int64Attribute{
									Optional: true,
								},
								"include_subdomains": schema.BoolAttribute{
									Optional: true,
								},
								"subpath_matching": schema.BoolAttribute{
									Optional: true,
								},
								"preserve_query_string": schema.BoolAttribute{
									Optional: true,
								},
								"preserve_path_suffix": schema.BoolAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		Blocks: map[string]schema.Block{
			// v4 item blocks with nested value sub-blocks
			"item": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							Optional: true,
						},
					},
					Blocks: map[string]schema.Block{
						"value": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"ip": schema.StringAttribute{
										Optional: true,
									},
									"asn": schema.Int64Attribute{
										Optional: true,
									},
								},
								Blocks: map[string]schema.Block{
									"hostname": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"url_hostname": schema.StringAttribute{
													Required: true,
												},
											},
										},
									},
									"redirect": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"source_url": schema.StringAttribute{
													Required: true,
												},
												"target_url": schema.StringAttribute{
													Required: true,
												},
												"status_code": schema.Int64Attribute{
													Optional: true,
												},
												"include_subdomains": schema.StringAttribute{
													Optional: true,
												},
												"subpath_matching": schema.StringAttribute{
													Optional: true,
												},
												"preserve_query_string": schema.StringAttribute{
													Optional: true,
												},
												"preserve_path_suffix": schema.StringAttribute{
													Optional: true,
												},
											},
										},
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
