package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareEmailRoutingRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareEmailRoutingRuleSchema(),
		ReadContext:   resourceCloudflareEmailRoutingRuleRead,
		CreateContext: resourceCloudflareEmailRoutingRuleCreate,
		UpdateContext: resourceCloudflareEmailRoutingRuleUpdate,
		DeleteContext: resourceCloudflareEmailRoutingRuleDelete,
		Description: heredoc.Doc(`
			Provides a resource for managing Email Routing rules.
		`),
	}
}

func buildMatchersAndActions(d *schema.ResourceData) (matchers []cloudflare.EmailRoutingRuleMatcher, actions []cloudflare.EmailRoutingRuleAction) {
	if items, ok := d.GetOk("matcher"); ok {
		for _, item := range items.(*schema.Set).List() {
			matcher := item.(map[string]interface{})
			matcherStruct := cloudflare.EmailRoutingRuleMatcher{
				Type: matcher["type"].(string),
			}
			if val, ok := matcher["field"]; ok {
				matcherStruct.Field = val.(string)
			}
			if val, ok := matcher["value"]; ok {
				matcherStruct.Value = val.(string)
			}
			matchers = append(matchers, matcherStruct)
		}
	}

	if items, ok := d.GetOk("action"); ok {
		for _, item := range items.(*schema.Set).List() {
			action := item.(map[string]interface{})
			ruleAction := cloudflare.EmailRoutingRuleAction{}
			ruleAction.Type = action["type"].(string)
			for _, value := range action["value"].([]interface{}) {
				ruleAction.Value = append(ruleAction.Value, value.(string))
			}

			actions = append(actions, ruleAction)
		}
	}
	return
}

func resourceCloudflareEmailRoutingRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	res, err := client.GetEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading email routing rule %q: %w", d.Id(), err))
	}
	d.SetId(res.Tag)
	d.Set("name", res.Name)
	d.Set("enabled", cloudflare.Bool(res.Enabled))
	d.Set("priority", res.Priority)

	return nil
}

func resourceCloudflareEmailRoutingRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	createParams := cloudflare.CreateEmailRoutingRuleParameters{
		Name:     d.Get("name").(string),
		Enabled:  cloudflare.BoolPtr(d.Get("enabled").(bool)),
		Priority: d.Get("priority").(int),
	}

	createParams.Matchers, createParams.Actions = buildMatchersAndActions(d)

	res, err := client.CreateEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), createParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating email routing rule %q: %w", createParams.Name, err))
	}

	d.SetId(res.Tag)
	return resourceCloudflareEmailRoutingRuleRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	updateParams := cloudflare.UpdateEmailRoutingRuleParameters{
		RuleID:   d.Id(),
		Name:     d.Get("name").(string),
		Enabled:  cloudflare.BoolPtr(d.Get("enabled").(bool)),
		Priority: d.Get("priority").(int),
	}

	updateParams.Matchers, updateParams.Actions = buildMatchersAndActions(d)

	_, err := client.UpdateEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), updateParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating email routing rule %q: %w", updateParams.Name, err))
	}

	return resourceCloudflareEmailRoutingRuleRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	_, err := client.DeleteEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting email routing rule %q: %w", d.Id(), err))
	}

	return nil
}
