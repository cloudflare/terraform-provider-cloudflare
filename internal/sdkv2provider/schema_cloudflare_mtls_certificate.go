package sdkv2provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareMTLSCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:        schema.TypeString,
			Description: "The account identifier to target for the resource.",
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Optional unique name for the certificate.",
			Optional:    true,
			ForceNew:    true,
		},
		"issuer": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"signature": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"certificates": {
			Type:        schema.TypeString,
			Description: "Certificate you intend to use with mTLS-enabled services.",
			Required:    true,
			ForceNew:    true,
		},
		"private_key": {
			Type:        schema.TypeString,
			Description: "The certificate's private key.",
			Optional:    true,
			ForceNew:    true,
		},
		"ca": {
			Type:        schema.TypeBool,
			Description: "Whether this is a CA or leaf certificate.",
			Required:    true,
			ForceNew:    true,
		},
		"uploaded_on": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
	}
}
