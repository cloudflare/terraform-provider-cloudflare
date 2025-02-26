// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicNetworkMonitoringRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The id of the rule. Must be unique.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9, underscore (_), dash (-), period (.), and tilde (~). You can’t have a space in the rule name. Max 256 characters.",
				Required:    true,
			},
			"automatic_advertisement": schema.BoolAttribute{
				Description: "Toggle on if you would like Cloudflare to automatically advertise the IP Prefixes within the rule via Magic Transit when the rule is triggered. Only available for users of Magic Transit.",
				Optional:    true,
			},
			"bandwidth": schema.Float64Attribute{
				Description: "The number of bits per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"packet_threshold": schema.Float64Attribute{
				Description: "The number of packets per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"prefixes": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"duration": schema.StringAttribute{
				Description: "The amount of time that the rule threshold must be exceeded to send an alert notification. The final value must be equivalent to one of the following 8 values [\"1m\",\"5m\",\"10m\",\"15m\",\"20m\",\"30m\",\"45m\",\"60m\"].\navailable values: \"1m\", \"5m\", \"10m\", \"15m\", \"20m\", \"30m\", \"45m\", \"60m\"",
				Computed:    true,
				Optional:    true,
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
				Default: stringdefault.StaticString("1m"),
			},
			"bandwidth_threshold": schema.Float64Attribute{
				Description: "The number of bits per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
			"prefix_match": schema.StringAttribute{
				Description: "Prefix match type to be applied for a prefix auto advertisement when using an advanced_ddos rule.\navailable values: \"exact\", \"subnet\", \"supernet\"",
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
				Description: "MNM rule type.\navailable values: \"threshold\", \"zscore\", \"advanced_ddos\"",
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
				Description: "Level of sensitivity set for zscore rules.\navailable values: \"low\", \"medium\", \"high\"",
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
				Description: "Target of the zscore rule analysis.\navailable values: \"bits\", \"packets\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("bits", "packets"),
				},
			},
		},
	}
}

func (r *MagicNetworkMonitoringRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicNetworkMonitoringRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
