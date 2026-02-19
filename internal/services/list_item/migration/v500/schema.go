package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceListItemSchema returns the v4 (SDKv2) cloudflare_list_item schema (schema_version=0).
// In v4, hostname and redirect were TypeList with MaxItems: 1, which maps to ListNestedBlock.
// Redirect boolean fields are stored as strings ("enabled"/"disabled").
func SourceListItemSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"list_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ip": schema.StringAttribute{
				Optional: true,
			},
			"asn": schema.Int64Attribute{
				Optional: true,
			},
			"comment": schema.StringAttribute{
				Optional: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
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
	}
}

// SourceListItemV1Schema returns the v4.52.5 framework cloudflare_list_item schema (schema_version=1).
// In v4.52.5, hostname and redirect are still ListNestedBlocks, but redirect's boolean fields
// are actual BoolAttributes (v4.52.5's 0→1 upgrader converted from "enabled"/"disabled" strings).
func SourceListItemV1Schema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"list_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ip": schema.StringAttribute{
				Optional: true,
			},
			"asn": schema.Int64Attribute{
				Optional: true,
			},
			"comment": schema.StringAttribute{
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
	}
}
