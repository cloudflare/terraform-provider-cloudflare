package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

func resourceCloudflareAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the account that is displayed in the Cloudflare dashboard.",
		},
		"type": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Valid values are standard (default) and enterprise. For self-serve customers, use standard. For enterprise customers, use enterprise.",
			Default:      accountTypeStandard,
			ValidateFunc: validation.StringInSlice([]string{accountTypeEnterprise, accountTypeStandard}, false),
			ForceNew:     true,
		},
		"enforce_twofactor": {
			Description: "Whether 2FA is enforced on the account.",
			Type:        schema.TypeBool,
			Default:     false,
			Optional:    true,
		},
	}
}
