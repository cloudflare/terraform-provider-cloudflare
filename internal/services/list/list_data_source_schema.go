// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ListsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ListsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the list.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "The RFC 3339 timestamp of when the list was created.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the list.",
							Computed:    true,
						},
						"kind": schema.StringAttribute{
							Description: "The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).",
							Computed:    true,
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
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "An informative name for the list. Use this name in filter and rule expressions.",
							Computed:    true,
						},
						"num_items": schema.Float64Attribute{
							Description: "The number of items in the list.",
							Computed:    true,
						},
						"num_referencing_filters": schema.Float64Attribute{
							Description: "The number of [filters](/operations/filters-list-filters) referencing the list.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ListsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
