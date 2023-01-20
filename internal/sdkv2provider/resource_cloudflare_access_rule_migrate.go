package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessRuleV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"configuration": {
				Type:             schema.TypeMap,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: configurationDiffSuppress,
				ValidateFunc:     validateAccessRuleConfiguration,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"ip", "ip6", "ip_range", "asn", "country"}, false),
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

func resourceCloudflareAccessRuleStateUpgradeV1(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState["configuration"] = []interface{}{rawState["configuration"]}
	return rawState, nil
}
