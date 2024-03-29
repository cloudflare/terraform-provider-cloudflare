// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r FirewallRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the firewall rule.",
				Optional:    true,
			},
			"action": schema.StringAttribute{
				Description: "The action to apply to a matched request. The `log` action is only available on an Enterprise plan.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge", "allow", "log", "bypass"),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The unique identifier of the filter.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the filter.",
						Optional:    true,
					},
					"expression": schema.StringAttribute{
						Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
						Optional:    true,
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the filter is currently paused.",
						Optional:    true,
					},
					"ref": schema.StringAttribute{
						Description: "A short reference tag. Allows you to select related filters.",
						Optional:    true,
					},
					"deleted": schema.BoolAttribute{
						Description: "When true, indicates that the firewall rule was deleted.",
						Required:    true,
					},
				},
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the firewall rule is currently paused.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the firewall rule.",
				Optional:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.",
				Optional:    true,
			},
			"products": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"ref": schema.StringAttribute{
				Description: "A short reference tag. Allows you to select related firewall rules.",
				Optional:    true,
			},
		},
	}
}
