// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_keys_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessKeysConfigurationDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessKeysConfigurationDataSource{}

func (r AccessKeysConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
		},
	}
}

func (r *AccessKeysConfigurationDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessKeysConfigurationDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
