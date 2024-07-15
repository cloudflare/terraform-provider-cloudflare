// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &D1DatabasesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &D1DatabasesDataSource{}

func (r D1DatabasesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "a database name to search for.",
				Optional:    true,
			},
			"page": schema.Float64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"per_page": schema.Float64Attribute{
				Description: "Number of items per page.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(10, 10000),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Description: "Specifies the timestamp the resource was created as an ISO8601 string.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"uuid": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"version": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func (r *D1DatabasesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *D1DatabasesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
