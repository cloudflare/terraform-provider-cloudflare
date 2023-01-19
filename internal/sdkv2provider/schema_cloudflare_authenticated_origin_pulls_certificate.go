package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAuthenticatedOriginPullsCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"certificate": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The public client certificate.",
		},
		"private_key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			ForceNew:    true,
			Description: "The private key of the client certificate.",
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
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"uploaded_on": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"type": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"per-zone", "per-hostname"}, false),
			Required:     true,
			ForceNew:     true,
			Description:  fmt.Sprintf("The form of Authenticated Origin Pulls to upload the certificate to. %s", renderAvailableDocumentationValuesStringSlice([]string{"per-zone", "per-hostname"})),
		},
	}
}
