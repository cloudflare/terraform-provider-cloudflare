package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFirewallRule() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareFirewallRuleSchema(),
		CreateContext: resourceCloudflareFirewallRuleCreate,
		ReadContext: resourceCloudflareFirewallRuleRead,
		UpdateContext: resourceCloudflareFirewallRuleUpdate,
		DeleteContext: resourceCloudflareFirewallRuleDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareFirewallRuleImport,
		},
	}
}

func resourceCloudflareFirewallRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	var err error

	var newFirewallRule cloudflare.FirewallRule

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

	log.Printf("[DEBUG] Creating Cloudflare Firewall Rule from struct: %+v", newFirewallRule)

	var r []cloudflare.FirewallRule

	r, err = client.CreateFirewallRules(context.Background(), zoneID, []cloudflare.FirewallRule{newFirewallRule})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Firewall Rule for zone %q: %s", zoneID, err))
	}

	if len(r) == 0 {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r[0].ID)

	log.Printf("[INFO] Cloudflare Firewall Rule ID: %s", d.Id())

	return resourceCloudflareFirewallRuleRead(d, meta)
}

func resourceCloudflareFirewallRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	firewallRule, err := client.FirewallRule(context.Background(), zoneID, d.Id())

	log.Printf("[DEBUG] firewallRule: %#v", firewallRule)
	log.Printf("[DEBUG] firewallRule error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Firewall Rule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Firewall Rule %q: %s", d.Id(), err))
	}

	log.Printf("[DEBUG] Cloudflare Firewall Rule read configuration: %#v", firewallRule)

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
	zoneID := d.Get("zone_id").(string)

	var newFirewallRule cloudflare.FirewallRule
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

	log.Printf("[DEBUG] Updating Cloudflare Firewall Rule from struct: %+v", newFirewallRule)

	r, err := client.UpdateFirewallRule(context.Background(), zoneID, newFirewallRule)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Firewall Rule for zone %q: %s", zoneID, err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Update response; resource was empty"))
	}

	return resourceCloudflareFirewallRuleRead(d, meta)
}

func resourceCloudflareFirewallRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting Cloudflare Firewall Rule: id %s for zone %s", d.Id(), zoneID)

	err := client.DeleteFirewallRule(context.Background(), zoneID, d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Firewall Rule: %s", err))
	}

	return nil
}

func resourceCloudflareFirewallRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/ruleID\"", d.Id())
	}

	zoneID, ruleID := idAttr[0], idAttr[1]

	log.Printf("[DEBUG] Importing Cloudflare Firewall Rule: id %s for zone %s", ruleID, zoneID)

	d.Set("zone_id", zoneID)
	d.SetId(ruleID)

	resourceCloudflareFirewallRuleRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
