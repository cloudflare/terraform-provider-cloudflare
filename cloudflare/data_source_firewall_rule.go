package cloudflare

import (
	"context"
	"fmt"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceCloudflareFirewallRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareFirewallRulesRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"filter": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntBetween(1, 2147483647),
						},
						"match_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"lt", "lte", "eq", "gte", "gt"}, false),
						},
						"paused": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},

			"firewall_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"paused": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareFirewallRulesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	filter, err := expandFilterFirewallRules(d.Get("filter"))
	if err != nil {
		return err
	}

	zoneID := d.Get("zone_id").(string)
	rules, err := client.FirewallRules(context.Background(), zoneID, cloudflare.PaginationOptions{})
	if err != nil {
		return fmt.Errorf("error listing Firewall Rules: %s", err)
	}

	filteredRules := make([]map[string]interface{}, 0)
	for _, v := range rules {
		if filter.paused != v.Paused {
			continue
		}

		rulePrio, ok := v.Priority.(int)
		if !ok {
			rulePrio = 2147483647
		}
		if !filter.matchPriority(rulePrio) {
			continue
		}

		filteredRules = append(filteredRules, map[string]interface{}{
			"id":       v.ID,
			"action":   v.Action,
			"priority": v.Priority,
			"paused":   v.Paused,
		})
	}

	err = d.Set("firewall_rules", filteredRules)
	if err != nil {
		return fmt.Errorf("error setting Firewall Rules: %s", err)
	}

	return nil
}

func expandFilterFirewallRules(d interface{}) (*firewallRulesFilter, error) {
	cfg := d.([]interface{})
	filter := &firewallRulesFilter{}

	m := cfg[0].(map[string]interface{})
	priority, ok := m["priority"]
	if ok {
		filter.priority = priority.(int)
	}

	matchType, ok := m["match_type"]
	if ok {
		filter.matchType = matchType.(string)
	}

	paused, ok := m["paused"]
	if ok {
		filter.paused = paused.(bool)
	}

	return filter, nil
}

type firewallRulesFilter struct {
	priority  int
	matchType string
	paused    bool
}

func (f *firewallRulesFilter) matchPriority(match int) bool {
	switch f.matchType {
	case "lte":
		return match <= f.priority
	case "lt":
		return match < f.priority
	case "eq":
		return match == f.priority
	case "gt":
		return match > f.priority
	case "gte":
		return match >= f.priority
	default:
		return true
	}
}
