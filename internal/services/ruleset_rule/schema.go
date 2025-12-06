// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_rule

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"
)

var _ resource.ResourceWithConfigValidators = (*RulesetRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	attributes := ruleset.RuleAttributes(ctx)
	attributes["ruleset_id"] = schema.StringAttribute{
		Description: "The unique ID of the ruleset.",
		Required:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("^[0-9a-f]{32}$"),
				"value must be a 32-character hexadecimal string",
			),
		},
	}
	attributes["account_id"] = schema.StringAttribute{
		Description: "The unique ID of the account.",
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("zone_id")),
			stringvalidator.RegexMatches(
				regexp.MustCompile("^[0-9a-f]{32}$"),
				"value must be a 32-character hexadecimal string",
			),
		},
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	}
	attributes["zone_id"] = schema.StringAttribute{
		Description: "The unique ID of the zone.",
		Optional:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("^[0-9a-f]{32}$"),
				"value must be a 32-character hexadecimal string",
			),
		},
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	}
	attributes["position"] = schema.SingleNestedAttribute{
		Description: "Specifies where to place the rule. Only one of `index`, `before`, or `after` can be set.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"index": schema.Int64Attribute{
				Description: "The absolute index position for the rule (starting at 1).",
				Optional:    true,
			},
			"before": schema.StringAttribute{
				Description: "Place this rule immediately before the rule with the given ID.",
				Optional:    true,
			},
			"after": schema.StringAttribute{
				Description: "Place this rule immediately after the rule with the given ID.",
				Optional:    true,
			},
		},
	}

	return schema.Schema{
		Attributes: attributes,
	}
}

func (r *RulesetRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *RulesetRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		// Ensure exactly one of account_id or zone_id is specified
	}
}
