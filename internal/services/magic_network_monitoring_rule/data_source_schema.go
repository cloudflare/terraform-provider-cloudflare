// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicNetworkMonitoringRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"rule_id": schema.StringAttribute{
				Description: "The id of the rule. Must be unique.",
				Optional:    true,
			},
			"automatic_advertisement": schema.BoolAttribute{
				Description: "Toggle on if you would like Cloudflare to automatically advertise the IP Prefixes within the rule via Magic Transit when the rule is triggered. Only available for users of Magic Transit.",
				Computed:    true,
			},
			"bandwidth_threshold": schema.Float64Attribute{
				Description: "The number of bits per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"duration": schema.StringAttribute{
				Description: "The amount of time that the rule threshold must be exceeded to send an alert notification. The final value must be equivalent to one of the following 8 values [\"1m\",\"5m\",\"10m\",\"15m\",\"20m\",\"30m\",\"45m\",\"60m\"]. The format is AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at least one unit must be provided.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "The id of the rule. Must be unique.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9, underscore (_), dash (-), period (.), and tilde (~). You can’t have a space in the rule name. Max 256 characters.",
				Computed:    true,
			},
			"packet_threshold": schema.Float64Attribute{
				Description: "The number of packets per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"prefixes": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *MagicNetworkMonitoringRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicNetworkMonitoringRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("rule_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("rule_id")),
	}
}
