// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &ListItemDataSource{}
var _ datasource.DataSourceWithValidateConfig = &ListItemDataSource{}

func (r ListItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Optional:    true,
			},
			"item_id": schema.StringAttribute{
				Description: "The unique ID of the item in the List.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"list_id": schema.StringAttribute{
						Description: "The unique ID of the list.",
						Required:    true,
					},
					"cursor": schema.StringAttribute{
						Description: "The pagination cursor. An opaque string token indicating the position from which to continue when requesting the next/previous set of records. Cursor values are provided under `result_info.cursors` in the response. You should make no assumptions about a cursor's content or length.",
						Optional:    true,
					},
					"per_page": schema.Int64Attribute{
						Description: "Amount of results to include in each paginated response. A non-negative 32 bit integer.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 500),
						},
					},
					"search": schema.StringAttribute{
						Description: "A search query to filter returned items. Its meaning depends on the list type: IP addresses must start with the provided string, hostnames and bulk redirects must contain the string, and ASNs must match the string exactly.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *ListItemDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *ListItemDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
