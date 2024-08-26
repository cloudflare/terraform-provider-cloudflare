// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ListDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the list was created.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the list.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Optional:    true,
			},
			"kind": schema.StringAttribute{
				Description: "The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ip",
						"redirect",
						"hostname",
						"asn",
					),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the list was last modified.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "An informative name for the list. Use this name in filter and rule expressions.",
				Optional:    true,
			},
			"num_items": schema.Float64Attribute{
				Description: "The number of items in the list.",
				Optional:    true,
			},
			"num_referencing_filters": schema.Float64Attribute{
				Description: "The number of [filters](/operations/filters-list-filters) referencing the list.",
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *ListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("list_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("list_id")),
	}
}
