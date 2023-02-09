package sdkv2provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareRulesets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareRulesetsRead,

		Description: heredoc.Doc(`
			Use this datasource to lookup Rulesets in an account or zone.
		`),

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description:  "The account identifier to target for the resource.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"zone_id", "account_id"},
			},
			"zone_id": {
				Description:  "The zone identifier to target for the resource.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"zone_id", "account_id"},
			},

			"include_rules": {
				Description: "Include rule data in response",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the Ruleset to target.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the ruleset.",
						},
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
						},
						"kind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: fmt.Sprintf("Type of Ruleset to create. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetKindValues())),
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Version of the ruleset to filter on.",
						},
					},
				},
			},

			"rulesets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID of the ruleset.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the ruleset.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Brief summary of the ruleset and its intended use.",
						},
						"kind": {
							Type:        schema.TypeString,
							Required:    true,
							Description: fmt.Sprintf("Type of Ruleset. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetKindValues())),
						},
						"version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Version of the ruleset.",
						},
						"phase": {
							Type:        schema.TypeString,
							Required:    true,
							Description: fmt.Sprintf("Point in the request/response lifecycle where the ruleset executes. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
						},
						"rules": resourceCloudflareRulesetSchema()["rules"],
					},
				},
			},
		},
	}
}

func dataSourceCloudflareRulesetsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	accountID := d.Get("account_id").(string)
	includeRules := d.Get("include_rules").(bool)
	filter, err := expandFilterRulesets(d.Get("filter"))
	if err != nil {
		return diag.FromErr(err)
	}

	var rulesetsList []cloudflare.Ruleset
	if accountID != "" {
		rulesetsList, err = client.ListAccountRulesets(ctx, accountID)
	} else {
		rulesetsList, err = client.ListZoneRulesets(ctx, zoneID)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	rulesets := make([]interface{}, 0)
	rulesetIds := make([]string, 0)
	for _, ruleset := range rulesetsList {
		if filter.ID != "" && filter.ID != ruleset.ID {
			continue
		}
		if filter.Phase != "" && filter.Phase != ruleset.Phase {
			continue
		}
		if filter.Kind != "" && filter.Kind != ruleset.Kind {
			continue
		}
		if filter.Version != "" && filter.Version != ruleset.Version {
			continue
		}
		if filter.Name != nil && !filter.Name.Match([]byte(ruleset.Name)) {
			continue
		}

		rulesetIds = append(rulesetIds, ruleset.ID)
		resultRuleset := map[string]interface{}{
			"id":          ruleset.ID,
			"name":        ruleset.Name,
			"description": ruleset.Description,
			"kind":        ruleset.Kind,
			"version":     ruleset.Version,
			"phase":       ruleset.Phase,
		}

		if includeRules {
			var fullRuleset cloudflare.Ruleset
			if accountID != "" {
				fullRuleset, err = client.GetAccountRuleset(ctx, accountID, ruleset.ID)
			} else {
				fullRuleset, err = client.GetZoneRuleset(ctx, zoneID, ruleset.ID)
			}
			if err != nil {
				return diag.FromErr(err)
			}

			rules := make([]interface{}, 0)
			for _, rule := range fullRuleset.Rules {
				fullRulesetRule := map[string]interface{}{
					"id":           rule.ID,
					"version":      rule.Version,
					"action":       rule.Action,
					"expression":   rule.Expression,
					"description":  rule.Description,
					"last_updated": rule.LastUpdated.String(),
					"ref":          rule.Ref,
					"enabled":      rule.Enabled,
				}

				if rule.RateLimit != nil {
					rl := make([]interface{}, 0)
					fullRulesetRule["ratelimit"] = append(rl, map[string]interface{}{
						"characteristics":     rule.RateLimit.Characteristics,
						"requests_per_period": rule.RateLimit.RequestsPerPeriod,
						"period":              rule.RateLimit.Period,
						"mitigation_timeout":  rule.RateLimit.MitigationTimeout,
						"counting_expression": rule.RateLimit.CountingExpression,
						"requests_to_origin":  rule.RateLimit.RequestsToOrigin,
					})
				}

				if rule.ExposedCredentialCheck != nil {
					ecc := make([]interface{}, 0)
					fullRulesetRule["exposed_credential_check"] = append(ecc, map[string]interface{}{
						"username_expression": rule.ExposedCredentialCheck.UsernameExpression,
						"password_expression": rule.ExposedCredentialCheck.PasswordExpression,
					})
				}
				if rule.Logging != nil {
					lg := make([]interface{}, 0)
					fullRulesetRule["logging"] = append(lg, map[string]interface{}{
						"enabled": rule.Logging.Enabled,
						"status":  loggingStatus(rule.Logging),
					})
				}
				rules = append(rules, fullRulesetRule)
			}
			resultRuleset["rules"] = rules
		}

		rulesets = append(rulesets, resultRuleset)
	}

	err = d.Set("rulesets", rulesets)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting Ruleset Ids: %w", err))
	}

	d.SetId(stringListChecksum(rulesetIds))
	return nil
}

func loggingStatus(logging *cloudflare.RulesetRuleLogging) string {
	if logging == nil || logging.Enabled == nil {
		return "default"
	}
	if *logging.Enabled {
		return "enabled"
	}
	return "disabled"
}

func expandFilterRulesets(d interface{}) (*searchFilterRulesets, error) {
	cfg := d.([]interface{})
	filter := &searchFilterRulesets{}
	if len(cfg) == 0 || cfg[0] == nil {
		return filter, nil
	}

	m := cfg[0].(map[string]interface{})
	if name, ok := m["name"]; ok {
		match, err := regexp.Compile(name.(string))
		if err != nil {
			return nil, err
		}

		filter.Name = match
	}

	if kind, ok := m["kind"]; ok {
		filter.Kind = kind.(string)
	}
	if phase, ok := m["phase"]; ok {
		filter.Phase = phase.(string)
	}
	if id, ok := m["id"]; ok {
		filter.ID = id.(string)
	}
	if version, ok := m["version"]; ok {
		filter.Version = version.(string)
	}

	return filter, nil
}

type searchFilterRulesets struct {
	Name    *regexp.Regexp
	Kind    string
	Phase   string
	ID      string
	Version string
}
