// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &WaitingRoomRulesListDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WaitingRoomRulesListDataSource{}

func (r WaitingRoomRulesListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"waiting_room_id": schema.StringAttribute{
				Required: true,
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
						"id": schema.StringAttribute{
							Description: "The ID of the rule.",
							Computed:    true,
						},
						"action": schema.StringAttribute{
							Description: "The action to take when the expression matches.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("bypass_waiting_room"),
							},
						},
						"description": schema.StringAttribute{
							Description: "The description of the rule.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "When set to true, the rule is enabled.",
							Computed:    true,
						},
						"expression": schema.StringAttribute{
							Description: "Criteria defining when there is a match for the current rule.",
							Computed:    true,
						},
						"last_updated": schema.StringAttribute{
							Computed: true,
						},
						"version": schema.StringAttribute{
							Description: "The version of the rule.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *WaitingRoomRulesListDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WaitingRoomRulesListDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
