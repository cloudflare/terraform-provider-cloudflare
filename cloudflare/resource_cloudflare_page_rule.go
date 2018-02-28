package cloudflare

import (
	"fmt"
	"log"

	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudFlarePageRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlarePageRuleCreate,
		Read:   resourceCloudFlarePageRuleRead,
		Update: resourceCloudFlarePageRuleUpdate,
		Delete: resourceCloudFlarePageRuleDelete,
		// TODO Importer

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},

			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"target": {
				Type:     schema.TypeString,
				Required: true,
			},

			"effective_target": {
				Type:     schema.TypeString,
				Computed: true, // better to have a diffsuppressfunc on 'api'  but its unclear what changes the api can make
			},

			"actions": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					SchemaVersion: 1,
					Schema: map[string]*schema.Schema{
						"always_online": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"automatic_https_rewrites": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"browser_check": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"email_obfuscation": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"ip_geolocation": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"opportunistic_encryption": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"server_side_exclude": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"smart_errors": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"always_use_https": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"disable_apps": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						// may not be used with rocket loader
						// n.b. ConflictsWith doesn't seem to work on nested schemas
						"disable_performance": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"disable_security": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},

						"browser_cache_ttl": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtMost(31536000),
						},

						"edge_cache_ttl": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtMost(31536000),
						},

						"cache_level": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "basic", "simplified", "aggressive", "cache_everything"}, false),
						},

						"forwarding_url": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								SchemaVersion: 1,
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Required: true,
									},

									"status_code": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(301, 302),
									},
								},
							},
						},

						// may not be used with disable_performance
						"rocket_loader": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"off", "manual", "automatic"}, false),
						},

						"security_level": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"essentially_off", "low", "medium", "high", "under_attack"}, false),
						},

						"ssl": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict"}, false),
						},
					},
				},
			},

			"priority": {
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},

			"status": {
				Type:         schema.TypeString,
				Default:      "active",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "paused"}, false),
			},
		},
	}
}

func resourceCloudFlarePageRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	domain := d.Get("domain").(string)

	newPageRuleTargets := []cloudflare.PageRuleTarget{
		{
			Target: "url",
			Constraint: struct {
				Operator string `json:"operator"`
				Value    string `json:"value"`
			}{
				Operator: "matches",
				Value:    d.Get("target").(string),
			},
		},
	}

	actions := d.Get("actions").([]interface{})
	newPageRuleActions := make([]cloudflare.PageRuleAction, 0, len(actions))

	log.Printf("[DEBUG] Actions found in config: %#v", actions)
	for _, action := range actions {
		for id, value := range action.(map[string]interface{}) {
			newPageRuleAction, err := transformToCloudFlarePageRuleAction(id, value)
			if err != nil {
				return err
			} else if newPageRuleAction.Value == nil {
				continue
			}
			newPageRuleActions = append(newPageRuleActions, newPageRuleAction)
		}
	}

	newPageRule := cloudflare.PageRule{
		Targets:  newPageRuleTargets,
		Actions:  newPageRuleActions,
		Priority: d.Get("priority").(int),
		Status:   d.Get("status").(string),
	}

	zoneID, err := client.ZoneIDByName(domain)
	if err != nil {
		return fmt.Errorf("Error finding zone %q: %s", domain, err)
	}

	d.Set("zone_id", zoneID)
	log.Printf("[DEBUG] CloudFlare Page Rule create configuration: %#v", newPageRule)

	r, err := client.CreatePageRule(zoneID, newPageRule)
	if err != nil {
		return fmt.Errorf("Failed to create page rule: %s", err)
	}

	if r.ID == "" {
		return fmt.Errorf("Failed to find page rule in Create response; ID was empty")
	}

	d.SetId(r.ID)

	return resourceCloudFlarePageRuleRead(d, meta)
}

func resourceCloudFlarePageRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	pageRule, err := client.PageRule(zoneID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Invalid Page Rule identifier") || // api bug - this indicates non-existing resource
			strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Page Rule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("Error finding page rule %q: %s", d.Id(), err)
		}
	}
	log.Printf("[DEBUG] CloudFlare Page Rule read configuration: %#v", pageRule)

	// Cloudflare presently only has one target type, and its Operator is always
	// "matches"; so we can just read the first element's Value.
	d.Set("effective_target", pageRule.Targets[0].Constraint.Value)

	actions := map[string]interface{}{}
	for _, pageRuleAction := range pageRule.Actions {
		key, value, err := transformFromCloudFlarePageRuleAction(&pageRuleAction)
		if err != nil {
			return fmt.Errorf("Failed to parse page rule action: %s", err)
		}
		actions[key] = value
	}
	log.Printf("[DEBUG] CloudFlare Page Rule actions configuration: %#v", actions)
	d.Set("actions", []map[string]interface{}{actions})

	d.Set("priority", pageRule.Priority)
	d.Set("status", pageRule.Status)

	return nil
}

func resourceCloudFlarePageRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	updatePageRule := cloudflare.PageRule{}

	if target, ok := d.GetOk("target"); ok {
		updatePageRule.Targets = []cloudflare.PageRuleTarget{
			cloudflare.PageRuleTarget{
				Target: "url",
				Constraint: struct {
					Operator string `json:"operator"`
					Value    string `json:"value"`
				}{
					Operator: "matches",
					Value:    target.(string),
				},
			},
		}
	}

	if v, ok := d.GetOk("actions"); ok {
		actions := v.([]interface{})
		newPageRuleActions := make([]cloudflare.PageRuleAction, 0, len(actions))

		for _, action := range actions {
			for id, value := range action.(map[string]interface{}) {
				newPageRuleAction, err := transformToCloudFlarePageRuleAction(id, value)
				if err != nil {
					return err
				} else if newPageRuleAction.Value == nil {
					continue
				}
				newPageRuleActions = append(newPageRuleActions, newPageRuleAction)
			}
		}

		updatePageRule.Actions = newPageRuleActions
	}

	if priority, ok := d.GetOk("priority"); ok {
		updatePageRule.Priority = priority.(int)
	}

	if status, ok := d.GetOk("status"); ok {
		updatePageRule.Status = status.(string)
	}

	log.Printf("[DEBUG] Cloudflare Page Rule update configuration: %#v", updatePageRule)

	// contrary to docs, change page rule actually does a full replace
	// this part of the api needs some work, so it may change in future
	if err := client.ChangePageRule(zoneID, d.Id(), updatePageRule); err != nil {
		return fmt.Errorf("Failed to update Cloudflare Page Rule: %s", err)
	}

	return resourceCloudFlarePageRuleRead(d, meta)
}

func resourceCloudFlarePageRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	domain := d.Get("domain").(string)

	log.Printf("[INFO] Deleting Cloudflare Page Rule: %s, %s", domain, d.Id())

	if err := client.DeletePageRule(zoneID, d.Id()); err != nil {
		return fmt.Errorf("Error deleting Cloudflare Page Rule: %s", err)
	}

	return nil
}

func transformFromCloudFlarePageRuleAction(pageRuleAction *cloudflare.PageRuleAction) (key string, value interface{}, err error) {
	key = pageRuleAction.ID

	switch {
	case contains(pageRuleAPIOnOffFields, pageRuleAction.ID):
		if strings.EqualFold(pageRuleAction.Value.(string), "on") {
			value = true
		} else {
			value = false
		}
		break

	case contains(pageRuleAPINilFields, pageRuleAction.ID):
		// api returns a nil value
		value = true
		break

	case contains(pageRuleAPIFloatFields, pageRuleAction.ID):
		value = pageRuleAction.Value.(float64)
		break

	case contains(pageRuleAPIStringFields, pageRuleAction.ID):
		value = pageRuleAction.Value.(string)
		break

	case pageRuleAction.ID == "forwarding_url":
		value = pageRuleAction.Value.(map[string]interface{})
		break

	default:
		// User supplied ID is already validated, so this is always an internal error
		err = fmt.Errorf("Unimplemented action ID %q - this is always an internal error", pageRuleAction.ID)
	}
	return
}

func transformToCloudFlarePageRuleAction(id string, value interface{}) (pageRuleAction cloudflare.PageRuleAction, err error) {

	pageRuleAction.ID = id

	if strValue, ok := value.(string); ok {
		if strValue == "" || strValue == "off" {
			pageRuleAction.Value = nil
		} else {
			pageRuleAction.Value = strValue
		}
	} else if unitValue, ok := value.(bool); ok {
		if !unitValue {
			pageRuleAction.Value = nil
		} else {
			if contains(pageRuleAPIOnOffFields, id) {
				pageRuleAction.Value = "on"
			} else {
				pageRuleAction.Value = true
			}
		}
	} else if intValue, ok := value.(int); ok {
		if intValue == 0 {
			// This happens when not set by the user
			pageRuleAction.Value = nil
		} else {
			pageRuleAction.Value = intValue
		}
	} else if id == "forwarding_url" {
		forwardActionSchema := value.([]interface{})

		log.Printf("[DEBUG] forwarding_url action to be applied: %#v", forwardActionSchema)

		if len(forwardActionSchema) != 0 {
			fwd := forwardActionSchema[0].(map[string]interface{})

			pageRuleAction.Value = map[string]interface{}{
				"url":         fwd["url"].(string),
				"status_code": fwd["status_code"].(int),
			}
		}
	} else {
		err = fmt.Errorf("Bad value for %s: %s", id, value)
	}

	return
}

var pageRuleAPIOnOffFields = []string{"always_online", "automatic_https_rewrites", "browser_check", "email_obfuscation", "ip_geolocation", "opportunistic_encryption", "server_side_exclude", "smart_errors"}
var pageRuleAPINilFields = []string{"always_use_https", "disable_apps", "disable_performance", "disable_security"}
var pageRuleAPIFloatFields = []string{"browser_cache_ttl", "edge_cache_ttl"}
var pageRuleAPIStringFields = []string{"cache_level", "rocket_loader", "security_level", "ssl"}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
