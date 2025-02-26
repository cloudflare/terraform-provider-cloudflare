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

var _ datasource.DataSourceWithConfigValidators = (*WaitingRoomRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"waiting_room_id": schema.StringAttribute{
				Required: true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"action": schema.StringAttribute{
				Description: "The action to take when the expression matches.\navailable values: \"bypass_waiting_room\"",
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
			"id": schema.StringAttribute{
				Description: "The ID of the rule.",
				Computed:    true,
			},
			"last_updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"version": schema.StringAttribute{
				Description: "The version of the rule.",
				Computed:    true,
			},
		},
	}
}

func (d *WaitingRoomRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WaitingRoomRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
