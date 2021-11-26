package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAuthenticatedOriginPullsCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"certificate": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"private_key": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			ForceNew:  true,
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
		},
	}
}
