package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

func dataSourceCloudflareDevicePostureRules() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Device Posture Rule.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(devicePostureRuleTypes, false),
				Description:  fmt.Sprintf("The device posture rule type. %s", renderAvailableDocumentationValuesStringSlice(devicePostureRuleTypes)),
			},
			"rules": {
				Description: "A list of matching Device Posture Rules.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID of the Device Posture Rule.",
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(devicePostureRuleTypes, false),
							Description:  fmt.Sprintf("The device posture rule type. %s", renderAvailableDocumentationValuesStringSlice(devicePostureRuleTypes)),
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the device posture rule.",
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"schedule": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tells the client when to run the device posture check. Must be in the format `1h` or `30m`. Valid units are `h` and `m`.",
						},
						"expiration": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Expire posture results after the specified amount of time. Must be in the format `1h` or `30m`. Valid units are `h` and `m`.",
						},
					},
				},
			},
		},
		Description: "Use this data source to lookup a list of [Device Posture Rule](https://developers.cloudflare.com/cloudflare-one/identity/devices)",
		ReadContext: dataSourceCloudflareDevicePostureRuleRead,
	}
}

func dataSourceCloudflareDevicePostureRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	if accountID == "" {
		return diag.FromErr(fmt.Errorf(`error determining resource: "account_id" required`))
	}

	name := d.Get("name")
	//type is a reserved keyword
	typ := d.Get("type")

	rules, _, err := client.DevicePostureRules(ctx, accountID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Device Posture Rules: %w", err))
	}
	if len(rules) == 0 {
		return diag.Errorf("no Device Posture Rules found")
	}
	matchedRules := make([]interface{}, 0)
	ruleIds := make([]string, 0)
	for _, rule := range rules {
		if name != "" && rule.Name != name {
			continue
		}
		if typ != "" && rule.Type != typ {
			continue
		}
		ruleIds = append(ruleIds, rule.ID)
		matchedRules = append(matchedRules, map[string]interface{}{
			"id":          rule.ID,
			"name":        rule.Name,
			"type":        rule.Type,
			"description": rule.Description,
			"schedule":    rule.Schedule,
			"expiration":  rule.Expiration,
		})
	}

	if len(matchedRules) == 0 {
		return diag.Errorf("no Device Posture Rules matching name %q type %q", name, typ)
	}

	if err = d.Set("rules", matchedRules); err != nil {
		return diag.FromErr(fmt.Errorf("error setting matched rules: %w", err))
	}

	d.SetId(stringListChecksum(ruleIds))
	return nil
}
