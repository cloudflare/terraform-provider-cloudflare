// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &SpectrumApplicationDataSource{}

func (d *SpectrumApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Description: "Application identifier.",
				Optional:    true,
			},
			"zone": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Sets the direction by which results are ordered.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Application field by which results are ordered.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("protocol", "app_id", "created_on", "modified_on", "dns"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results. This parameter is required in order to use other pagination parameters. If included in the query, `result_info` will be present in the response.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.AtLeast(1),
						},
					},
					"per_page": schema.Float64Attribute{
						Description: "Sets the maximum number of results per page.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(1, 100),
						},
					},
				},
			},
		},
	}
}

func (d *SpectrumApplicationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
