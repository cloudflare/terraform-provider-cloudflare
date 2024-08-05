// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &WaitingRoomRulesDataSource{}

func (d *WaitingRoomRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"action": schema.StringAttribute{
				Description: "The action to take when the expression matches.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("bypass_waiting_room"),
				},
			},
			"expression": schema.StringAttribute{
				Description: "Criteria defining when there is a match for the current rule.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The ID of the rule.",
				Optional:    true,
			},
			"last_updated": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"version": schema.StringAttribute{
				Description: "The version of the rule.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the rule.",
				Computed:    true,
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "When set to true, the rule is enabled.",
				Computed:    true,
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"waiting_room_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *WaitingRoomRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
