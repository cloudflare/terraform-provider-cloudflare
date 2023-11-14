package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Account ID to manage the zone resource in.",
		},
		"zone": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: zoneDiffFunc,
			Description:      "The DNS zone name which will be added.",
		},
		"jump_start": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to scan for DNS records on creation. Ignored after zone is created.",
		},
		"paused": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether this zone is paused (traffic bypasses Cloudflare)",
			Default:     false,
		},
		"vanity_name_servers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of Vanity Nameservers (if set).",
		},
		"plan": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ValidateFunc: validation.StringInSlice([]string{
				planIDFree,
				planIDLite,
				planIDPro,
				planIDProPlus,
				planIDBusiness,
				planIDEnterprise,
				planIDPartnerFree,
				planIDPartnerPro,
				planIDPartnerBusiness,
				planIDPartnerEnterprise,
			}, false),
			Description: fmt.Sprintf("The name of the commercial plan to apply to the zone. %s", renderAvailableDocumentationValuesStringSlice([]string{
				planIDFree,
				planIDLite,
				planIDPro,
				planIDProPlus,
				planIDBusiness,
				planIDEnterprise,
				planIDPartnerFree,
				planIDPartnerPro,
				planIDPartnerBusiness,
				planIDPartnerEnterprise,
			})),
		},
		"meta": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: fmt.Sprintf("Status of the zone. %s ", renderAvailableDocumentationValuesStringSlice([]string{"active", "pending", "initializing", "moved", "deleted", "deactivated"})),
		},
		"type": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"full", "partial", "secondary"}, false),
			Default:      "full",
			Optional:     true,
			Description:  fmt.Sprintf("A full zone implies that DNS is hosted with Cloudflare. A partial zone is typically a partner-hosted zone or a CNAME setup. %s", renderAvailableDocumentationValuesStringSlice([]string{"full", "partial", "secondary"})),
		},
		"name_servers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Cloudflare-assigned name servers. This is only populated for zones that use Cloudflare DNS.",
		},
		"verification_key": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Contains the TXT record value to validate domain ownership. This is only populated for zones of type `partial`.",
		},
	}
}
