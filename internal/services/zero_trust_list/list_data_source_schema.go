// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &ZeroTrustListsDataSource{}

func (d *ZeroTrustListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Description: "The type of list.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("SERIAL", "URL", "DOMAIN", "EMAIL", "IP"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "API Resource UUID tag.",
							Computed:    true,
							Optional:    true,
						},
						"list_count": schema.Float64Attribute{
							Description: "The number of items in the list.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "The description of the list.",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the list.",
							Computed:    true,
							Optional:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of list.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("SERIAL", "URL", "DOMAIN", "EMAIL", "IP"),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustListsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
