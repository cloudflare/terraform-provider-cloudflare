package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudflareAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccessRuleCreate,
		Read:   resourceCloudflareAccessRuleRead,
		Update: resourceCloudflareAccessRuleUpdate,
		Delete: resourceCloudflareAccessRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "whitelist", "js_challenge"}, false),
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration": {
				Type:             schema.TypeMap,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: configurationDiffSuppress,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"ip", "ip_range", "asn", "country"}, false),
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareAccessRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zone := d.Get("zone").(string)
	zone_id := d.Get("zone_id").(string)

	newRule := cloudflare.AccessRule{
		Notes: d.Get("notes").(string),
		Mode:  d.Get("mode").(string),
	}

	if configuration, configurationOk := d.GetOk("configuration"); configurationOk {
		config := configuration.(map[string]interface{})

		newRule.Configuration = cloudflare.AccessRuleConfiguration{
			Target: config["target"].(string),
			Value:  config["value"].(string),
		}
	}

	var r *cloudflare.AccessRuleResponse
	var err error

	if zone == "" && zone_id == "" {
		if client.OrganizationID != "" {
			r, err = client.CreateOrganizationAccessRule(client.OrganizationID, newRule)
		} else {
			r, err = client.CreateUserAccessRule(newRule)
		}
	} else {
		var zoneID string

		if zone_id != "" {
			zoneID = zone_id
		} else {
			zoneID, err = client.ZoneIDByName(zone)
			if err != nil {
				return fmt.Errorf("Error finding zone %q: %s", zone, err)
			}

			d.Set("zone_id", zoneID)
		}

		r, err = client.CreateZoneAccessRule(zoneID, newRule)
	}

	if err != nil {
		return fmt.Errorf("Failed to create access rule: %s", err)
	}

	if r.Result.ID == "" {
		return fmt.Errorf("Failed to find access rule in Create response; ID was empty")
	}

	d.SetId(r.Result.ID)

	return resourceCloudflareAccessRuleRead(d, meta)
}

func resourceCloudflareAccessRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	var accessRuleResponse *cloudflare.AccessRuleResponse
	var err error

	if zoneID == "" {
		if client.OrganizationID != "" {
			accessRuleResponse, err = client.OrganizationAccessRule(client.OrganizationID, d.Id())
		} else {
			accessRuleResponse, err = client.UserAccessRule(d.Id())
		}
	} else {
		accessRuleResponse, err = client.ZoneAccessRule(zoneID, d.Id())
	}

	log.Printf("[DEBUG] accessRuleResponse: %#v", accessRuleResponse)
	log.Printf("[DEBUG] accessRuleResponse error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Page Rule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding access rule %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Cloudflare Access Rule read configuration: %#v", accessRuleResponse)

	d.Set("mode", accessRuleResponse.Result.Mode)
	d.Set("notes", accessRuleResponse.Result.Notes)
	log.Printf("[DEBUG] read configuration: %#v", d.Get("configuration"))

	configuration := map[string]interface{}{}
	configuration["target"] = accessRuleResponse.Result.Configuration.Target
	configuration["value"] = accessRuleResponse.Result.Configuration.Value

	d.Set("configuration", configuration)

	return nil
}

func resourceCloudflareAccessRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	newRule := cloudflare.AccessRule{
		Notes: d.Get("notes").(string),
		Mode:  d.Get("mode").(string),
	}

	if configuration, configurationOk := d.GetOk("configuration"); configurationOk {
		config := configuration.(map[string]interface{})

		newRule.Configuration = cloudflare.AccessRuleConfiguration{
			Target: config["target"].(string),
			Value:  config["value"].(string),
		}
	}

	// var accessRuleResponse *cloudflare.AccessRuleResponse
	var err error

	if zoneID == "" {
		if client.OrganizationID != "" {
			_, err = client.UpdateOrganizationAccessRule(client.OrganizationID, d.Id(), newRule)
		} else {
			_, err = client.UpdateUserAccessRule(d.Id(), newRule)
		}
	} else {
		_, err = client.UpdateZoneAccessRule(zoneID, d.Id(), newRule)
	}

	if err != nil {
		return fmt.Errorf("Failed to update Access Rule: %s", err)
	}

	return resourceCloudflareAccessRuleRead(d, meta)
}

func resourceCloudflareAccessRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting Cloudflare Access Rule: id %s for zone_id %s", d.Id(), zoneID)

	var err error

	if zoneID == "" {
		if client.OrganizationID != "" {
			_, err = client.DeleteOrganizationAccessRule(client.OrganizationID, d.Id())
		} else {
			_, err = client.DeleteUserAccessRule(d.Id())
		}
	} else {
		_, err = client.DeleteZoneAccessRule(zoneID, d.Id())
	}

	if err != nil {
		return fmt.Errorf("Error deleting Cloudflare Access Rule: %s", err)
	}

	return nil
}

func configurationDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	switch {
	case d.Get("configuration.target") == "country" &&
		k == "configuration.value":
		return strings.ToUpper(old) == strings.ToUpper(new)
	case d.Get("configuration.target") == "asn" &&
		k == "configuration.value":

		if !strings.HasPrefix(strings.ToUpper(new), "AS") {
			new = "AS" + strings.ToUpper(new)
		}

		return strings.ToUpper(old) == strings.ToUpper(new)
	}

	return false
}
