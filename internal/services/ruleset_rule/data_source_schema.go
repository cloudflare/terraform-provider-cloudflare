// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_rule

import (
	"context"
	"regexp"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func DataSourceSchema(ctx context.Context) schema.Schema {
	// Get the parent ruleset schema to extract the rule schema
	rulesetSchema := ruleset.DataSourceSchema(ctx)
	rulesListAttr := rulesetSchema.Attributes["rules"].(schema.ListNestedAttribute)
	ruleAttributes := rulesListAttr.NestedObject.Attributes

	// Create schema for individual rule datasource
	return schema.Schema{
		Description: "Use this data source to lookup a single rule within a Cloudflare Ruleset.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique ID of the rule.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			"rule_id": schema.StringAttribute{
				Description: "The unique ID of the rule to lookup.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			"ruleset_id": schema.StringAttribute{
				Description: "The unique ID of the ruleset containing the rule.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			"account_id": schema.StringAttribute{
				Description: "The unique ID of the account.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.MatchRoot("zone_id")),
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			"zone_id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			// Include all rule attributes from the parent ruleset schema except "id"
			// which we've already defined above
			"action":                   ruleAttributes["action"],
			"action_parameters":        ruleAttributes["action_parameters"],
			"description":              ruleAttributes["description"],
			"enabled":                  ruleAttributes["enabled"],
			"exposed_credential_check": ruleAttributes["exposed_credential_check"],
			"expression":               ruleAttributes["expression"],
			"logging":                  ruleAttributes["logging"],
			"ratelimit":                ruleAttributes["ratelimit"],
			"ref":                      ruleAttributes["ref"],
			"categories":               ruleAttributes["categories"],
		},
	}
}

func (d *RulesetRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *RulesetRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
