package sdkv2provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareApiTokenSchema() map[string]*schema.Schema {
	p := schema.Resource{
		Schema: map[string]*schema.Schema{
			"resources": {
				Type:        schema.TypeMap,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Describes what operations against which resources are allowed or denied.",
			},
			"permission_groups": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of permissions groups IDs. See [documentation](https://developers.cloudflare.com/api/tokens/create/permissions) for more information.",
			},
			"effect": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "allow",
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),
				Description:  fmt.Sprintf("Effect of the policy. %s", renderAvailableDocumentationValuesStringSlice([]string{"allow", "deny"})),
			},
		},
	}

	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the API Token.",
		},
		"policy": {
			Type:        schema.TypeSet,
			Set:         schema.HashResource(&p),
			Required:    true,
			Elem:        &p,
			Description: "Permissions policy. Multiple policy blocks can be defined.",
		},
		"condition": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Conditions under which the token should be considered valid.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"request_ip": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Request IP related conditions.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"in": {
									Type:        schema.TypeSet,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Optional:    true,
									Description: "List of IP addresses or CIDR notation where the token may be used from. If not specified, the token will be valid for all IP addresses.",
								},
								"not_in": {
									Type:        schema.TypeSet,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Optional:    true,
									Description: "List of IP addresses or CIDR notation where the token should not be used from.",
								},
							},
						},
					},
				},
			},
		},
		"not_before": {
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Description:  "The time before which the token MUST NOT be accepted for processing",
			Optional:     true,
		},
		"expires_on": {
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Description:  "The expiration time on or after which the token MUST NOT be accepted for processing",
			Optional:     true,
		},
		"value": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "The value of the API Token.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issued_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the token was issued.",
		},
		"modified_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the token was last modified.",
		},
	}
}
