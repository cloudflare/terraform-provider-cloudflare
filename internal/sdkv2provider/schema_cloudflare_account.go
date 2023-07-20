package sdkv2provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

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
			Description:  fmt.Sprintf("Account type. %s", renderAvailableDocumentationValuesStringSlice([]string{accountTypeEnterprise, accountTypeStandard})),
			Default:      accountTypeStandard,
			ValidateFunc: validation.StringInSlice([]string{accountTypeEnterprise, accountTypeStandard}, false),
			ForceNew:     true, // "Updating account type is not supported from client api"
		},
		"enforce_twofactor": {
			Description: "Whether 2FA is enforced on the account.",
			Type:        schema.TypeBool,
			Default:     false,
			Optional:    true,
		},
	}
}
