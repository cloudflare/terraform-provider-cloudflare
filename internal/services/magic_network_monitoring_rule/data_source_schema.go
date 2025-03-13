// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicNetworkMonitoringRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The id of the rule. Must be unique.",
				Computed:    true,
			},
			"rule_id": schema.StringAttribute{
				Description: "The id of the rule. Must be unique.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
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
				Description: "The amount of time that the rule threshold must be exceeded to send an alert notification. The final value must be equivalent to one of the following 8 values [\"1m\",\"5m\",\"10m\",\"15m\",\"20m\",\"30m\",\"45m\",\"60m\"].\nAvailable values: \"1m\", \"5m\", \"10m\", \"15m\", \"20m\", \"30m\", \"45m\", \"60m\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"1m",
						"5m",
						"10m",
						"15m",
						"20m",
						"30m",
						"45m",
						"60m",
					),
				},
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
			"prefix_match": schema.StringAttribute{
				Description: "Prefix match type to be applied for a prefix auto advertisement when using an advanced_ddos rule.\nAvailable values: \"exact\", \"subnet\", \"supernet\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"exact",
						"subnet",
						"supernet",
					),
				},
			},
			"type": schema.StringAttribute{
				Description: "MNM rule type.\nAvailable values: \"threshold\", \"zscore\", \"advanced_ddos\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"threshold",
						"zscore",
						"advanced_ddos",
					),
				},
			},
			"zscore_sensitivity": schema.StringAttribute{
				Description: "Level of sensitivity set for zscore rules.\nAvailable values: \"low\", \"medium\", \"high\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"low",
						"medium",
						"high",
					),
				},
			},
			"zscore_target": schema.StringAttribute{
				Description: "Target of the zscore rule analysis.\nAvailable values: \"bits\", \"packets\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("bits", "packets"),
				},
			},
			"prefixes": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *MagicNetworkMonitoringRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicNetworkMonitoringRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
