// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &TeamsListDataSource{}
var _ datasource.DataSourceWithValidateConfig = &TeamsListDataSource{}

func (r TeamsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"list_id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Optional:    true,
			},
			"count": schema.Float64Attribute{
				Description: "The number of items in the list.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the list.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the list.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of list.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("SERIAL", "URL", "DOMAIN", "EMAIL", "IP"),
				},
			},
			"updated_at": schema.StringAttribute{
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
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
				},
			},
		},
	}
}

func (r *TeamsListDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *TeamsListDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
