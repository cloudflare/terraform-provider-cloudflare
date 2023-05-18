package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFirewallRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareFirewallRuleSchema(),
		CreateContext: resourceCloudflareFirewallRuleCreate,
		ReadContext:   resourceCloudflareFirewallRuleRead,
		UpdateContext: resourceCloudflareFirewallRuleUpdate,
		DeleteContext: resourceCloudflareFirewallRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareFirewallRuleImport,
		},
		Description: heredoc.Doc(`
			Define Firewall rules using filter expressions for more control over
			how traffic is matched to the rule. A filter expression permits
			selecting traffic by multiple criteria allowing greater freedom in
			rule creation.

			Filter expressions needs to be created first before using Firewall
			Rule.
		`),
		DeprecationMessage: heredoc.Doc(fmt.Sprintf(`
			%s resource is in a deprecation phase that will
			last for one year (May 1st, 2024). During this time period, this
			resource is still fully supported but you are strongly advised
			to move to the %s resource. For more information, see
			https://developers.cloudflare.com/waf/reference/migration-guides/firewall-rules-to-custom-rules/#relevant-changes-for-terraform-users.
		`, "`cloudflare_firewall_rule`", "`cloudflare_ruleset`")),
	}
}

func resourceCloudflareFirewallRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var err error

	var newFirewallRule cloudflare.FirewallRuleCreateParams

	if paused, ok := d.GetOk("paused"); ok {
		newFirewallRule.Paused = paused.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newFirewallRule.Description = description.(string)
	}

	if action, ok := d.GetOk("action"); ok {
		newFirewallRule.Action = action.(string)
	}

	if priority, ok := d.GetOk("priority"); ok {
		newFirewallRule.Priority = priority.(int)
	}

	if filterID, ok := d.GetOk("filter_id"); ok {
		newFirewallRule.Filter = cloudflare.Filter{
			ID: filterID.(string),
		}
	}

	if products, ok := d.GetOk("products"); ok {
		newFirewallRule.Products = expandInterfaceToStringList(products.(*schema.Set).List())
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Firewall Rule from struct: %+v", newFirewallRule))

	var r []cloudflare.FirewallRule

	r, err = client.CreateFirewallRules(ctx, cloudflare.ZoneIdentifier(zoneID), []cloudflare.FirewallRuleCreateParams{newFirewallRule})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Firewall Rule for zone %q: %w", zoneID, err))
	}

	if len(r) == 0 {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r[0].ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Firewall Rule ID: %s", d.Id()))

	return resourceCloudflareFirewallRuleRead(ctx, d, meta)
}

func resourceCloudflareFirewallRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	firewallRule, err := client.FirewallRule(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())

	tflog.Debug(ctx, fmt.Sprintf("firewallRule: %#v", firewallRule))
	tflog.Debug(ctx, fmt.Sprintf("firewallRule error: %#v", err))

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Firewall Rule %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Firewall Rule %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Firewall Rule read configuration: %#v", firewallRule))

	products := expandStringListToSet(firewallRule.Products)
	d.Set("paused", firewallRule.Paused)
	d.Set("description", firewallRule.Description)
	d.Set("action", firewallRule.Action)
	d.Set("priority", firewallRule.Priority)
	d.Set("filter_id", firewallRule.Filter.ID)
	d.Set("products", products)

	return nil
}

func resourceCloudflareFirewallRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var newFirewallRule cloudflare.FirewallRuleUpdateParams
	newFirewallRule.ID = d.Id()

	if paused, ok := d.GetOk("paused"); ok {
		newFirewallRule.Paused = paused.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newFirewallRule.Description = description.(string)
	}

	if action, ok := d.GetOk("action"); ok {
		newFirewallRule.Action = action.(string)
	}

	if priority, ok := d.GetOk("priority"); ok {
		newFirewallRule.Priority = priority.(int)
	}

	if filterID, ok := d.GetOk("filter_id"); ok {
		newFirewallRule.Filter = cloudflare.Filter{
			ID: filterID.(string),
		}
	}

	if products, ok := d.GetOk("products"); ok {
		newFirewallRule.Products = expandInterfaceToStringList(products.(*schema.Set).List())
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Firewall Rule from struct: %+v", newFirewallRule))

	r, err := client.UpdateFirewallRule(ctx, cloudflare.ZoneIdentifier(zoneID), newFirewallRule)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Firewall Rule for zone %q: %w", zoneID, err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Update response; resource was empty"))
	}

	return resourceCloudflareFirewallRuleRead(ctx, d, meta)
}

func resourceCloudflareFirewallRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Firewall Rule: id %s for zone %s", d.Id(), zoneID))

	err := client.DeleteFirewallRule(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Firewall Rule: %w", err))
	}

	return nil
}

func resourceCloudflareFirewallRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/ruleID\"", d.Id())
	}

	zoneID, ruleID := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Firewall Rule: id %s for zone %s", ruleID, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(ruleID)

	resourceCloudflareFirewallRuleRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
