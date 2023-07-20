package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTotalTLSSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"enabled": {
			Description: "Enable Total TLS for the zone.",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"certificate_authority": {
			Description:  fmt.Sprintf("The Certificate Authority that Total TLS certificates will be issued through. %s", renderAvailableDocumentationValuesStringSlice([]string{"google", "lets_encrypt"})),
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"google", "lets_encrypt"}, false),
		},
	}
}
