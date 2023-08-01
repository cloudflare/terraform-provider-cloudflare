package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWebAnalyticsRule() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareWebAnalyticsRuleSchema(),

		CreateContext: resourceCloudflareWebAnalyticsRuleCreate,
		ReadContext:   resourceCloudflareWebAnalyticsRuleRead,
		UpdateContext: resourceCloudflareWebAnalyticsRuleUpdate,
		DeleteContext: resourceCloudflareWebAnalyticsRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWebAnalyticsRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
		},
		Description: "Provides a Cloudflare Web Analytics Rule resource.",
	}
}

func buildCreateWebAnalyticsRuleParams(d *schema.ResourceData) cloudflare.CreateWebAnalyticsRuleParams {
	paths := make([]string, 0)
	newPaths := d.Get("paths").([]interface{})
	for _, p := range newPaths {
		paths = append(paths, p.(string))
	}
	return cloudflare.CreateWebAnalyticsRuleParams{
		RulesetID: d.Get("ruleset_id").(string),
		Rule: cloudflare.CreateWebAnalyticsRule{
			Host:      d.Get("host").(string),
			Paths:     paths,
			Inclusive: d.Get("inclusive").(bool),
			IsPaused:  d.Get("is_paused").(bool),
		},
	}
}

func buildUpdateWebAnalyticsRuleParams(d *schema.ResourceData) cloudflare.UpdateWebAnalyticsRuleParams {
	paths := make([]string, 0)
	newPaths := d.Get("paths").([]interface{})
	for _, p := range newPaths {
		paths = append(paths, p.(string))
	}
	return cloudflare.UpdateWebAnalyticsRuleParams{
		RuleID:    d.Id(),
		RulesetID: d.Get("ruleset_id").(string),
		Rule: cloudflare.CreateWebAnalyticsRule{
			Host:      d.Get("host").(string),
			Paths:     paths,
			Inclusive: d.Get("inclusive").(bool),
			IsPaused:  d.Get("is_paused").(bool),
		},
	}
}

func resourceCloudflareWebAnalyticsRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := buildCreateWebAnalyticsRuleParams(d)
	rule, err := client.CreateWebAnalyticsRule(ctx, cloudflare.AccountIdentifier(accountID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating web analytics rule%q: %w", d.Id(), err))
	}

	d.SetId(rule.ID)

	return resourceCloudflareWebAnalyticsRuleRead(ctx, d, meta)
}

func resourceCloudflareWebAnalyticsRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := buildUpdateWebAnalyticsRuleParams(d)
	rule, err := client.UpdateWebAnalyticsRule(ctx, cloudflare.AccountIdentifier(accountID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating web analytics rule %q: %w", d.Id(), err))
	}

	d.SetId(rule.ID)

	return resourceCloudflareWebAnalyticsRuleRead(ctx, d, meta)
}

func resourceCloudflareWebAnalyticsRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	rules, err := client.ListWebAnalyticsRules(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWebAnalyticsRulesParams{
		RulesetID: d.Get("ruleset_id").(string),
	})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing web analytics rule from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error getting web analytics rule %q: %w", d.Id(), err))
	}
	for _, rule := range rules.Rules {
		if d.Id() == rule.ID {
			d.SetId(rule.ID)
			d.Set("host", rule.Host)
			d.Set("paths", rule.Paths)
			d.Set("inclusive", rule.Inclusive)
			d.Set("is_paused", rule.IsPaused)
			return nil
		}
	}

	return diag.FromErr(fmt.Errorf("error getting web analytics rule %q: %w", d.Id(), errors.New("web analytics rule not found in ruleset")))
}

func resourceCloudflareWebAnalyticsRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	ruleID := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := cloudflare.DeleteWebAnalyticsRuleParams{
		RulesetID: d.Get("ruleset_id").(string),
		RuleID:    ruleID,
	}
	_, err := client.DeleteWebAnalyticsRule(ctx, cloudflare.AccountIdentifier(accountID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting web analytics rule %q: %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareWebAnalyticsRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), "/", 3)
	var accountID string
	var rulesetID string
	var ruleID string
	if len(idAttr) == 3 {
		accountID = idAttr[0]
		rulesetID = idAttr[1]
		ruleID = idAttr[2]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/rulesetID/ruleID\" for import", d.Id())
	}

	rules, err := client.ListWebAnalyticsRules(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWebAnalyticsRulesParams{
		RulesetID: rulesetID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch web analytics rule: %s", ruleID)
	}

	for _, rule := range rules.Rules {
		if rule.ID == ruleID {
			d.SetId(rule.ID)
			d.Set("ruleset_id", rulesetID)
			d.Set(consts.AccountIDSchemaKey, accountID)
			resourceCloudflareWebAnalyticsRuleRead(ctx, d, meta)
			return []*schema.ResourceData{d}, nil
		}
	}

	return nil, fmt.Errorf("error importing web analytics rule %q: %w", d.Id(), errors.New("web analytics rule not found in ruleset"))
}
