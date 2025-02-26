// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PageShieldPolicyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"policy_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"action": schema.StringAttribute{
				Description: "The action to take if the expression matches\navailable values: \"allow\", \"log\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "log"),
				},
			},
			"description": schema.StringAttribute{
				Description: "A description for the policy",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the policy is enabled",
				Computed:    true,
			},
			"expression": schema.StringAttribute{
				Description: "The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax",
				Computed:    true,
			},
			"value": schema.StringAttribute{
				Description: "The policy which will be applied",
				Computed:    true,
			},
		},
	}
}

func (d *PageShieldPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PageShieldPolicyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
