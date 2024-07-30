// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &WorkersKVDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WorkersKVDataSource{}

func (r WorkersKVDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Required:    true,
			},
			"key_name": schema.StringAttribute{
				Description: "A key's name. The name may be at most 512 bytes. All printable, non-whitespace characters are valid. Use percent-encoding to define key names as part of a URL.",
				Required:    true,
			},
		},
	}
}

func (r *WorkersKVDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WorkersKVDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
