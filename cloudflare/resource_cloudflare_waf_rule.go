package cloudflare

import (
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCloudflareWAFRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWAFRuleCreate,
		Read:   resourceCloudflareWAFRuleRead,
		Update: resourceCloudflareWAFRuleUpdate,
		Delete: resourceCloudflareWAFRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWAFRuleImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"rule_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"zone": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "`zone` is deprecated in favour of explicit `zone_id` and will be removed in the next major release",
			},

			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"package_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"mode": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCloudflareWAFRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	ruleID := d.Get("rule_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	rule, err := client.WAFRule(zoneID, packageID, ruleID)
	if err != nil {
		return (err)
	}

	// Only need to set mode as that is the only attribute that could have changed
	d.Set("mode", rule.Mode)
	d.SetId(rule.ID)

	return nil
}

func resourceCloudflareWAFRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	ruleID := d.Get("rule_id").(string)
	zone := d.Get("zone").(string)
	zoneID := d.Get("zone_id").(string)
	mode := d.Get("mode").(string)

	// While we are deprecating `zone`, we need to perform the validation
	// inside the `Create` to ensure we get at least one of the expected
	// values.
	if zone == "" && zoneID == "" {
		return fmt.Errorf("either zone name or ID must be provided")
	}

	if zoneID == "" {
		var err error
		zoneID, err = client.ZoneIDByName(zone)
		if err != nil {
			return err
		}
	}

	packs, err := client.ListWAFPackages(zoneID)
	if err != nil {
		return err
	}

	for _, p := range packs {
		rule, err := client.WAFRule(zoneID, p.ID, ruleID)

		if err == nil {
			d.Set("zone", zone)
			d.Set("zone_id", zoneID)
			d.Set("package_id", rule.PackageID)
			d.Set("mode", mode)

			// Set the ID to the rule_id parameter passed in from the user.
			// All WAF rules already exist so we already know the rule_id e.g. 100000.
			//
			// This is a work around as we are not really "creating" a WAF Rule,
			// only associating it with our terraform config for future updates.
			d.SetId(ruleID)

			if rule.Mode != mode {
				return resourceCloudflareWAFRuleUpdate(d, meta)
			}

			return nil
		}
	}

	return fmt.Errorf("Unable to find WAF Rule %s", ruleID)
}

func resourceCloudflareWAFRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	ruleID := d.Get("rule_id").(string)
	zoneID := d.Get("zone_id").(string)
	packageID := d.Get("package_id").(string)

	rule, err := client.WAFRule(zoneID, packageID, ruleID)
	if err != nil {
		return err
	}

	// Can't delete WAF Rule so instead reset it to default
	if rule.Mode != "default" {
		_, err = client.UpdateWAFRule(zoneID, packageID, ruleID, "default")
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceCloudflareWAFRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	ruleID := d.Get("rule_id").(string)
	zoneID := d.Get("zone_id").(string)
	mode := d.Get("mode").(string)
	packageID := d.Get("package_id").(string)

	// We can only update the mode of a WAF Rule
	_, err := client.UpdateWAFRule(zoneID, packageID, ruleID, mode)
	if err != nil {
		return err
	}

	return nil
}

func resourceCloudflareWAFRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneName string
	var WAFID string
	if len(idAttr) == 2 {
		zoneName = idAttr[0]
		WAFID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneName/WAFID\" for import", d.Id())
	}

	zoneID, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return nil, fmt.Errorf("error finding zoneName %q: %s", zoneName, err)
	}

	packs, err := client.ListWAFPackages(zoneID)
	if err != nil {
		panic(err)
	}

	for _, p := range packs {
		rule, err := client.WAFRule(zoneID, p.ID, WAFID)
		if err == nil {
			d.Set("rule_id", rule.ID)
			d.Set("zone", zoneName)
			d.Set("zone_id", zoneID)
			d.Set("package_id", rule.PackageID)
			d.Set("mode", rule.Mode)

			// The ID is known by the user in advance
			d.SetId(WAFID)
		}
	}

	if d.Id() != WAFID {
		return nil, fmt.Errorf("Unable to find WAF Rule %s", WAFID)
	}

	return []*schema.ResourceData{d}, nil
}
