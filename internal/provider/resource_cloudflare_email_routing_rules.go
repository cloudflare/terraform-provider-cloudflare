package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
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
	}
}

func resourceCloudflareEmailRoutingRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	res, err := client.GetEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading email routing rule %q: %w", d.Id(), err))
	}
	d.SetId(res.Tag)
	d.Set("name", res.Name)
	d.Set("enabled", res.Enabled)
	d.Set("priority", res.Priority)

	return nil
}

func resourceCloudflareEmailRoutingRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	createParams := cloudflare.CreateEmailRoutingRuleParameters{}
	createParams.Name = d.Get("name").(string)
	createParams.Enabled = cloudflare.BoolPtr(d.Get("enabled").(bool))
	createParams.Priority = d.Get("priority").(int)

	if items, ok := d.GetOk("matchers"); ok {
		var matchRules []cloudflare.EmailRoutingRuleMatcher
		for _, item := range items.(*schema.Set).List() {
			matcher := item.(map[string]string)
			matchRules = append(matchRules, cloudflare.EmailRoutingRuleMatcher{
				Type:  matcher["type"],
				Field: matcher["field"],
				Value: matcher["value"],
			})
		}
		createParams.Matchers = matchRules
	}

	if items, ok := d.GetOk("actions"); ok {
		var actions []cloudflare.EmailRoutingRuleAction
		for _, item := range items.(*schema.Set).List() {
			action := item.(map[string]interface{})
			actions = append(actions, cloudflare.EmailRoutingRuleAction{
				Type:  action["type"].(string),
				Value: action["value"].([]string),
			})
		}
		createParams.Actions = actions
	}

	res, err := client.CreateEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), createParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating email routing rule %q: %w", createParams.Name, err))
	}

	d.SetId(res.Tag)
	return resourceCloudflareEmailRoutingRuleRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	createParams := cloudflare.UpdateEmailRoutingRuleParameters{}
	createParams.Name = d.Get("name").(string)
	createParams.Enabled = cloudflare.BoolPtr(d.Get("enabled").(bool))
	createParams.Priority = d.Get("priority").(int)

	if items, ok := d.GetOk("matchers"); ok {
		var matchRules []cloudflare.EmailRoutingRuleMatcher
		for _, item := range items.(*schema.Set).List() {
			matcher := item.(map[string]string)
			matchRules = append(matchRules, cloudflare.EmailRoutingRuleMatcher{
				Type:  matcher["type"],
				Field: matcher["field"],
				Value: matcher["value"],
			})
		}
		createParams.Matchers = matchRules
	}

	if items, ok := d.GetOk("actions"); ok {
		var actions []cloudflare.EmailRoutingRuleAction
		for _, item := range items.(*schema.Set).List() {
			action := item.(map[string]interface{})
			actions = append(actions, cloudflare.EmailRoutingRuleAction{
				Type:  action["type"].(string),
				Value: action["value"].([]string),
			})
		}
		createParams.Actions = actions
	}

	res, err := client.UpdateEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), createParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating email routing rule %q: %w", createParams.Name, err))
	}

	d.SetId(res.Tag)
	return resourceCloudflareEmailRoutingRuleRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	_, err := client.DeleteEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting email routing rule %q: %w", d.Id(), err))
	}

	return nil
}
