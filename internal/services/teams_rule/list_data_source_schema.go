// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &TeamsRulesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &TeamsRulesDataSource{}

func (r TeamsRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
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
							Description: "The API resource UUID.",
							Computed:    true,
						},
						"action": schema.StringAttribute{
							Description: "The action to preform when the associated traffic, identity, and device posture expressions are either absent or evaluate to `true`.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("on", "off", "allow", "block", "scan", "noscan", "safesearch", "ytrestricted", "isolate", "noisolate", "override", "l4_override", "egress", "audit_ssh", "resolve"),
							},
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"deleted_at": schema.StringAttribute{
							Description: "Date of deletion, if any.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "The description of the rule.",
							Computed:    true,
						},
						"device_posture": schema.StringAttribute{
							Description: "The wirefilter expression used for device posture check matching.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "True if the rule is enabled.",
							Computed:    true,
						},
						"filters": schema.ListAttribute{
							Description: "The protocol or layer to evaluate the traffic, identity, and device posture expressions.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"identity": schema.StringAttribute{
							Description: "The wirefilter expression used for identity matching.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the rule.",
							Computed:    true,
						},
						"precedence": schema.Int64Attribute{
							Description: "Precedence sets the order of your rules. Lower values indicate higher precedence. At each processing phase, applicable rules are evaluated in ascending order of this value.",
							Computed:    true,
						},
						"traffic": schema.StringAttribute{
							Description: "The wirefilter expression used for traffic matching.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *TeamsRulesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *TeamsRulesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
