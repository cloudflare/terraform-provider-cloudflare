package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_response": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message_body": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringLenBetween(0, 1024),
									},

									"status_code": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"content_type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringLenBetween(0, 32),
									},

									"location": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringLenBetween(0, 2048),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareLoadBalancerStateUpgradeV1(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	if rawState["rules"] != nil {
		for i := range rawState["rules"].([]interface{}) {
			rawState["rules"].([]interface{})[i].(map[string]interface{})["fixed_response"] = []interface{}{rawState["rules"].([]interface{})[i].(map[string]interface{})["fixed_response"]}
		}
	}

	return rawState, nil
}
