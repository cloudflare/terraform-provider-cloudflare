package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCloudflareWAFOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWAFOverrideCreate,
		Read:   resourceCloudflareWAFOverrideRead,
		Update: resourceCloudflareWAFOverrideUpdate,
		Delete: resourceCloudflareWAFOverrideDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWAFOverrideImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"urls": {
				Required: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rules": {
				Required: true,
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"paused": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(-1000000000, 1000000000),
			},
			"groups": {
				Optional: true,
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rewrite_action": {
				Optional: true,
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceCloudflareWAFOverrideRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	_, err := client.WAFOverride(zoneID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "wafuriconfig.api.not_found") {
			log.Printf("[INFO] WAF override %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("failed to find WAF override %s: %s", d.Id(), err)
	}

	return nil
}

func resourceCloudflareWAFOverrideCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	newOverride := cloudflare.WAFOverride{}

	urls := d.Get("urls").([]interface{})
	for _, url := range urls {
		newOverride.URLs = append(newOverride.URLs, url.(string))
	}

	rules := d.Get("rules").(map[string]interface{})
	rulesMap := make(map[string]string)
	for ruleID, state := range rules {
		rulesMap[ruleID] = state.(string)
	}
	newOverride.Rules = rulesMap

	if pausedValue, ok := d.GetOk("paused"); ok {
		newOverride.Paused = pausedValue.(bool)
	}

	if descriptionValue, ok := d.GetOk("description"); ok {
		newOverride.Description = descriptionValue.(string)
	}

	if priorityValue, ok := d.GetOk("priority"); ok {
		newOverride.Priority = priorityValue.(int)
	}

	if groupsValue, ok := d.GetOk("groups"); ok {
		groupsMap := make(map[string]string)
		for groupID, state := range groupsValue.(map[string]interface{}) {
			groupsMap[groupID] = state.(string)
		}
		newOverride.Groups = groupsMap
	}

	if rewriteActionValue, ok := d.GetOk("rewrite_action"); ok {
		rewriteActions := make(map[string]string)
		for rewriteOriginal, rewriteWant := range rewriteActionValue.(map[string]interface{}) {
			rewriteActions[rewriteOriginal] = rewriteWant.(string)
		}
		newOverride.RewriteAction = rewriteActions
	}

	override, err := client.CreateWAFOverride(zoneID, newOverride)
	if err != nil {
		return fmt.Errorf("failed to create WAF override: %s", err)
	}

	d.SetId(override.ID)

	return resourceCloudflareWAFOverrideRead(d, meta)
}

func resourceCloudflareWAFOverrideDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	overrideID := d.Get("override_id").(string)
	zoneID := d.Get("zone_id").(string)

	err := client.DeleteWAFOverride(zoneID, overrideID)
	if err != nil {
		return fmt.Errorf("failed to delete WAF override ID %s: %s", overrideID, err)
	}

	return resourceCloudflareWAFOverrideRead(d, meta)
}
