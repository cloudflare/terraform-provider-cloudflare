// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ListItemDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"item_id": schema.StringAttribute{
				Description: "The unique ID of the item in the List.",
				Optional:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Optional:    true,
			},
			"include_subdomains": schema.BoolAttribute{
				Computed: true,
			},
			"preserve_path_suffix": schema.BoolAttribute{
				Computed: true,
			},
			"preserve_query_string": schema.BoolAttribute{
				Computed: true,
			},
			"source_url": schema.StringAttribute{
				Computed: true,
			},
			"status_code": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.OneOf(
						301,
						302,
						307,
						308,
					),
				},
			},
			"subpath_matching": schema.BoolAttribute{
				Computed: true,
			},
			"target_url": schema.StringAttribute{
				Computed: true,
			},
			"url_hostname": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"list_id": schema.StringAttribute{
						Description: "The unique ID of the list.",
						Required:    true,
					},
					"search": schema.StringAttribute{
						Description: "A search query to filter returned items. Its meaning depends on the list type: IP addresses must start with the provided string, hostnames and bulk redirects must contain the string, and ASNs must match the string exactly.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ListItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ListItemDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(
			path.MatchRoot("account_identifier"),
			path.MatchRoot("item_id"),
			path.MatchRoot("list_id"),
		),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_identifier")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("item_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("list_id")),
	}
}
