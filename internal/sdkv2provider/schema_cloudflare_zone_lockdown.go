package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareZoneLockdownSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"paused": {
			Type:        schema.TypeBool,
			Default:     false,
			Optional:    true,
			Description: "Boolean of whether this zone lockdown is currently paused",
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
			Description:  "A description about the lockdown entry. Typically used as a reminder or explanation for the lockdown.",
		},
		"urls": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "A list of simple wildcard patterns to match requests against. The order of the urls is unimportant.",
		},
		"configurations": {
			Type:        schema.TypeSet,
			MinItems:    1,
			Required:    true,
			Description: "A list of IP addresses or IP ranges to match the request against specified in target, value pairs.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"target": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"ip", "ip_range"}, false),
						Description:  fmt.Sprintf("The request property to target. %s", renderAvailableDocumentationValuesStringSlice([]string{"ip", "ip_range"})),
					},
					"value": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The value to target. Depends on target's type. IP addresses should just be standard IPv4/IPv6 notation i.e. `192.0.2.1` or `2001:db8::/32` and IP ranges in CIDR format i.e. `192.0.2.0/24`",
					},
				},
			},
		},
	}
}
