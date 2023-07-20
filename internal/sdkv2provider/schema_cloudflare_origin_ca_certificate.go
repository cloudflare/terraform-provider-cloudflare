package sdkv2provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareOriginCACertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Origin CA certificate.",
		},
		"csr": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Required:     true,
			ValidateFunc: validateCSR,
			Description:  "The Certificate Signing Request. Must be newline-encoded.",
		},
		"expires_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The datetime when the certificate will expire.",
		},
		"hostnames": {
			Type:     schema.TypeSet,
			Required: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "A list of hostnames or wildcard names bound to the certificate.",
		},
		"request_type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"origin-rsa", "origin-ecc", "keyless-certificate"}, false),
			Description:  fmt.Sprintf("The signature type desired on the certificate. %s", renderAvailableDocumentationValuesStringSlice([]string{"origin-rsa", "origin-ecc", "keyless-certificate"})),
		},
		"requested_validity": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			ValidateFunc: validation.IntInSlice([]int{7, 30, 90, 365, 730, 1095, 5475}),
			Description:  fmt.Sprintf("The number of days for which the certificate should be valid. %s", renderAvailableDocumentationValuesIntSlice([]int{7, 30, 90, 365, 730, 1095, 5475})),
		},
		"min_days_for_renewal": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Number of days prior to the expiry to trigger a renewal of the certificate if a Terraform operation is run.",
		},
	}
}
