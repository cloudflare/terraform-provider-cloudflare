// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &ZoneHoldDataSource{}
var _ datasource.DataSourceWithValidateConfig = &ZoneHoldDataSource{}

func (r ZoneHoldDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"hold": schema.BoolAttribute{
				Optional: true,
			},
			"hold_after": schema.StringAttribute{
				Optional: true,
			},
			"include_subdomains": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *ZoneHoldDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *ZoneHoldDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
