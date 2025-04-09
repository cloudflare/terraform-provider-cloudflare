// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersForPlatformsScriptSecretDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A JavaScript variable name for the secret binding.",
				Computed:    true,
			},
			"secret_name": schema.StringAttribute{
				Description: "A JavaScript variable name for the secret binding.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"dispatch_namespace": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Required:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of this secret, this is what will be used to access it inside the Worker.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of secret.\nAvailable values: \"secret_text\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("secret_text"),
				},
			},
		},
	}
}

func (d *WorkersForPlatformsScriptSecretDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersForPlatformsScriptSecretDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
