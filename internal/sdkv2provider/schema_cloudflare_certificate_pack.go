package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCertificatePackSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"advanced"}, false),
			Description:  fmt.Sprintf("Certificate pack configuration type. %s", renderAvailableDocumentationValuesStringSlice([]string{"advanced"})),
		},
		"hosts": {
			Type:     schema.TypeSet,
			Required: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of hostnames to provision the certificate pack for. The zone name must be included as a host. Note: If using Let's Encrypt, you cannot use individual subdomains and only a wildcard for subdomain is available.",
		},
		"validation_method": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"txt", "http", "email"}, false),
			Description:  fmt.Sprintf("Which validation method to use in order to prove domain ownership. %s", renderAvailableDocumentationValuesStringSlice([]string{"txt", "http", "email"})),
		},
		"validity_days": {
			Type:         schema.TypeInt,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IntInSlice([]int{14, 30, 90, 365}),
			Description:  fmt.Sprintf("How long the certificate is valid for. Note: If using Let's Encrypt, this value can only be 90 days. %s", renderAvailableDocumentationValuesIntSlice([]int{14, 30, 90, 365})),
		},
		"certificate_authority": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"digicert", "lets_encrypt", "google"}, false),
			Default:      nil,
			Description:  fmt.Sprintf("Which certificate authority to issue the certificate pack. %s", renderAvailableDocumentationValuesStringSlice([]string{"digicert", "lets_encrypt", "google"})),
		},
		"validation_records": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem:     sslValidationRecordsSchema(),
		},
		"validation_errors": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem:     sslValidationErrorsSchema(),
		},
		"cloudflare_branding": {
			Type:        schema.TypeBool,
			Optional:    true,
			ForceNew:    true,
			Description: "Whether or not to include Cloudflare branding. This will add `sni.cloudflaressl.com` as the Common Name if set to `true`.",
		},
		"wait_for_active_status": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Whether or not to wait for a certificate pack to reach status `active` during creation.",
		},
	}
}
