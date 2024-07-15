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

var _ datasource.DataSourceWithConfigValidators = &SpectrumApplicationsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &SpectrumApplicationsDataSource{}

func (r SpectrumApplicationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{},
				},
			},
		},
	}
}

func (r *SpectrumApplicationsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *SpectrumApplicationsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
