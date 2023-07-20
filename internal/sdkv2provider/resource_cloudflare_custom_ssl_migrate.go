package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomSSLV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"custom_ssl_options": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate": {
							Type:     schema.TypeString,
							Required: true,
						},
						"private_key": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"bundle_method": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"ubiquitous", "optimal", "force"}, false),
						},
						"geo_restrictions": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"us", "eu", "highest_security"}, false),
						},
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"legacy_custom", "sni_custom"}, false),
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareCustomSSLStateUpgradeV1(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState["custom_ssl_options"] = []interface{}{rawState["custom_ssl_options"]}
	return rawState, nil
}
