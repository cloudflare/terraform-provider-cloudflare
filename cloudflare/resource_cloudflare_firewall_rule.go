package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudflareFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareFirewallRuleCreate,
		Read:   resourceCloudflareFirewallRuleRead,
		Update: resourceCloudflareFirewallRuleUpdate,
		Delete: resourceCloudflareFirewallRuleDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareFirewallRuleImport,
		},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"filter_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "allow", "js_challenge", "log"}, false),
			},
			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 2147483647),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 500),
			},
			"paused": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceCloudflareFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneName := d.Get("zone").(string)
	zoneID := d.Get("zone_id").(string)

	var err error

	if zoneName == "" && zoneID == "" {
		return fmt.Errorf("'zone' or 'zone_id' is required")
	} else if zoneID == "" {
		zoneID, err = client.ZoneIDByName(zoneName)
		if err != nil {
			return fmt.Errorf("Error finding zone %q: %s", zoneName, err)
		}
		d.Set("zone_id", zoneID)
	}

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

	log.Printf("[DEBUG] Creating Cloudflare Firewall Rule from struct: %+v", newFirewallRule)

	var r []cloudflare.FirewallRule

	r, err = client.CreateFirewallRules(zoneID, []cloudflare.FirewallRule{newFirewallRule})

	if err != nil {
		return fmt.Errorf("error creating Firewall Rule for zone %q: %s", zoneName, err)
	}

	if len(r) == 0 {
		return fmt.Errorf("failed to find id in Create response; resource was empty")
	}

	d.SetId(r[0].ID)

	log.Printf("[INFO] Cloudflare Firewall Rule ID: %s", d.Id())

	return resourceCloudflareFirewallRuleRead(d, meta)
}

func resourceCloudflareFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	zoneName := d.Get("zone").(string)
	if zoneID == "" {
		zoneID, _ = client.ZoneIDByName(zoneName)
	} else {
		zones, err := client.ListZones()
		if err != nil {
			return fmt.Errorf("failed to lookup all zones: %s", err)
		}
		for _, zone := range zones {
			if zone.ID == zoneID {
				zoneName = zone.Name
			}
		}
	}

	firewallRule, err := client.FirewallRule(zoneID, d.Id())

	log.Printf("[DEBUG] firewallRule: %#v", firewallRule)
	log.Printf("[DEBUG] firewallRule error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Firewall Rule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Firewall Rule %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Cloudflare Firewall Rule read configuration: %#v", firewallRule)

	d.Set("zone", zoneName)
	d.Set("zone_id", zoneID)
	d.Set("paused", firewallRule.Paused)
	d.Set("description", firewallRule.Description)
	d.Set("action", firewallRule.Action)
	d.Set("priority", firewallRule.Priority)
	d.Set("filter_id", firewallRule.Filter.ID)

	return nil
}

func resourceCloudflareFirewallRuleUpdate(d *schema.ResourceData, meta interface{}) error {
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

	log.Printf("[DEBUG] Updating Cloudflare Firewall Rule from struct: %+v", newFirewallRule)

	r, err := client.UpdateFirewallRule(zoneID, newFirewallRule)

	if err != nil {
		return fmt.Errorf("error updating Firewall Rule for zone %q: %s", zoneID, err)
	}

	if r.ID == "" {
		return fmt.Errorf("failed to find id in Update response; resource was empty")
	}

	return resourceCloudflareFirewallRuleRead(d, meta)
}

func resourceCloudflareFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting Cloudflare Firewall Rule: id %s for zone %s", d.Id(), zoneID)

	err := client.DeleteFirewallRule(zoneID, d.Id())

	if err != nil {
		return fmt.Errorf("Error deleting Cloudflare Firewall Rule: %s", err)
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
