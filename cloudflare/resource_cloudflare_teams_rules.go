package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTeamsRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCloudflareTeamsRuleRead,
		Update: resourceCloudflareTeamsRuleUpdate,
		Create: resourceCloudflareTeamsRuleCreate,
		Delete: resourceCloudflareTeamsRuleDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareTeamsRuleImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"precedence": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(cloudflare.TeamsRulesActionValues(), false),
				Required:     true,
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"traffic": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"rule_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: teamsRuleSettings,
				},
			},
		},
	}
}

var teamsRuleSettings = map[string]*schema.Schema{
	"block_page_enabled": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"block_page_reason": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"override_ips": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
	"override_host": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"l4override": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsL4OverrideSettings,
		},
	},
	"biso_admin_controls": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: teamsBisoAdminControls,
		},
	},
}

var teamsL4OverrideSettings = map[string]*schema.Schema{
	"ip": {
		Type:     schema.TypeString,
		Required: true,
	},
	"port": {
		Type:     schema.TypeInt,
		Required: true,
	},
}

var teamsBisoAdminControls = map[string]*schema.Schema{
	"disable_printing": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"disable_copy_paste": {
		Type:     schema.TypeBool,
		Optional: true,
	},
}

func resourceCloudflareTeamsRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	rule, err := client.TeamsRule(context.Background(), accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 400") {
			log.Printf("[INFO] Teams Rule config %s doesnt exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Teams Rule %q: %s", d.Id(), err)
	}
	if err := d.Set("name", rule.Name); err != nil {
		return fmt.Errorf("error parsing rule name")
	}
	if err := d.Set("description", rule.Description); err != nil {
		return fmt.Errorf("error parsing rule description")
	}
	if err := d.Set("precedence", int64(rule.Precedence)); err != nil {
		return fmt.Errorf("error parsing rule precedence")
	}
	if err := d.Set("enabled", rule.Enabled); err != nil {
		return fmt.Errorf("error parsing rule enablement")
	}
	if err := d.Set("action", rule.Action); err != nil {
		return fmt.Errorf("error parsing rule action")
	}
	if err := d.Set("filters", rule.Filters); err != nil {
		return fmt.Errorf("error parsing rule filters")
	}
	if err := d.Set("traffic", rule.Traffic); err != nil {
		return fmt.Errorf("error parsing rule traffic")
	}
	if err := d.Set("identity", rule.Identity); err != nil {
		return fmt.Errorf("error parsing rule identity")
	}
	if err := d.Set("version", int64(rule.Version)); err != nil {
		return fmt.Errorf("error parsing rule version")
	}
	if err := d.Set("rule_settings", flattenTeamsRuleSettings(&rule.RuleSettings)); err != nil {
		return fmt.Errorf("error parsing rule settings")
	}
	return nil
}

func resourceCloudflareTeamsRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	accountID := d.Get("account_id").(string)
	settings := inflateTeamsRuleSettings(d.Get("rule_settings"))

	var filters []cloudflare.TeamsFilterType
	for _, f := range d.Get("filters").([]interface{}) {
		filters = append(filters, cloudflare.TeamsFilterType(f.(string)))
	}

	newTeamsRule := cloudflare.TeamsRule{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Precedence:  uint64(d.Get("precedence").(int)),
		Enabled:     d.Get("enabled").(bool),
		Action:      cloudflare.TeamsGatewayAction(d.Get("action").(string)),
		Filters:     filters,
		Traffic:     d.Get("traffic").(string),
		Identity:    d.Get("identity").(string),
		Version:     uint64(d.Get("version").(int)),
	}

	if settings != nil {
		newTeamsRule.RuleSettings = *settings
	}

	log.Printf("[DEBUG] Creating Cloudflare Teams Rule from struct: %+v", newTeamsRule)

	rule, err := client.TeamsCreateRule(context.Background(), accountID, newTeamsRule)
	if err != nil {
		return fmt.Errorf("error creating Teams rule for account %q: %s", accountID, err)
	}

	d.SetId(rule.ID)
	return resourceCloudflareTeamsRuleRead(d, meta)
}

func resourceCloudflareTeamsRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	settings := inflateTeamsRuleSettings(d.Get("rule_settings"))

	var filters []cloudflare.TeamsFilterType
	for _, f := range d.Get("filters").([]interface{}) {
		filters = append(filters, cloudflare.TeamsFilterType(f.(string)))
	}

	teamsRule := cloudflare.TeamsRule{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Precedence:  uint64(d.Get("precedence").(int)),
		Enabled:     d.Get("enabled").(bool),
		Action:      cloudflare.TeamsGatewayAction(d.Get("action").(string)),
		Filters:     filters,
		Traffic:     d.Get("traffic").(string),
		Identity:    d.Get("identity").(string),
		Version:     uint64(d.Get("version").(int)),
	}

	if settings != nil {
		teamsRule.RuleSettings = *settings
	}
	log.Printf("[DEBUG] Updating Cloudflare Teams rule from struct: %+v", teamsRule)

	updatedTeamsRule, err := client.TeamsUpdateRule(context.Background(), accountID, teamsRule.ID, teamsRule)
	if err != nil {
		return fmt.Errorf("error updating Teams rule for account %q: %s", accountID, err)
	}
	if updatedTeamsRule.ID == "" {
		return fmt.Errorf("failed to find Teams Rule ID in update response; resource was empty")
	}
	return resourceCloudflareTeamsRuleRead(d, meta)
}

func resourceCloudflareTeamsRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	id := d.Id()
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Teams Rule using ID: %s", id)

	err := client.TeamsDeleteRule(context.Background(), accountID, id)
	if err != nil {
		return fmt.Errorf("error deleting Teams Rule for account %q: %s", accountID, err)
	}

	return resourceCloudflareTeamsRuleRead(d, meta)
}

func resourceCloudflareTeamsRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsRuleID\"", d.Id())
	}

	accountID, teamsRuleID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Teams Rule: id %s for account %s", teamsRuleID, accountID)

	d.Set("account_id", accountID)
	d.SetId(teamsRuleID)

	err := resourceCloudflareTeamsRuleRead(d, meta)

	return []*schema.ResourceData{d}, err
}

func flattenTeamsRuleSettings(settings *cloudflare.TeamsRuleSettings) []interface{} {
	return []interface{}{map[string]interface{}{
		"block_page_enabled":  settings.BlockPageEnabled,
		"block_page_reason":   settings.BlockReason,
		"override_ips":        settings.OverrideIPs,
		"override_host":       settings.OverrideHost,
		"l4override":          flattenTeamsL4Override(settings.L4Override),
		"biso_admin_controls": flattenTeamsRuleBisoAdminControls(settings.BISOAdminControls),
	}}
}

func inflateTeamsRuleSettings(settings interface{}) *cloudflare.TeamsRuleSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	enabled := settingsMap["block_page_enabled"].(bool)
	reason := settingsMap["block_page_reason"].(string)

	var overrideIPs []string
	for _, ip := range settingsMap["override_ips"].([]interface{}) {
		overrideIPs = append(overrideIPs, ip.(string))
	}

	overrideHost := settingsMap["override_host"].(string)
	l4Override := inflateTeamsL4Override(settingsMap["l4override"].([]interface{}))

	bisoAdminControls := inflateTeamsRuleBisoAdminControls(settingsMap["biso_admin_controls"].([]interface{}))

	return &cloudflare.TeamsRuleSettings{
		BlockPageEnabled:  enabled,
		BlockReason:       reason,
		OverrideIPs:       overrideIPs,
		OverrideHost:      overrideHost,
		L4Override:        l4Override,
		BISOAdminControls: bisoAdminControls,
	}
}

func flattenTeamsRuleBisoAdminControls(settings *cloudflare.TeamsBISOAdminControlSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"disable_printing":   settings.DisablePrinting,
		"disable_copy_paste": settings.DisableCopyPaste,
	}}
}

func inflateTeamsRuleBisoAdminControls(settings interface{}) *cloudflare.TeamsBISOAdminControlSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	disablePrinting := settingsMap["disable_printing"].(bool)
	disableCopyPaste := settingsMap["disable_copy_paste"].(bool)
	return &cloudflare.TeamsBISOAdminControlSettings{
		DisablePrinting:  disablePrinting,
		DisableCopyPaste: disableCopyPaste,
	}
}

func flattenTeamsL4Override(settings *cloudflare.TeamsL4OverrideSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"ip":   settings.IP,
		"port": settings.Port,
	}}

}

func inflateTeamsL4Override(settings interface{}) *cloudflare.TeamsL4OverrideSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	ip := settingsMap["ip"].(string)
	port := settingsMap["port"].(int)
	return &cloudflare.TeamsL4OverrideSettings{
		IP:   ip,
		Port: port,
	}
}
