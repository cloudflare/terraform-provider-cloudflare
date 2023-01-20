package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessMutualTLSCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the certificate.",
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The Root CA for your certificates.",
		},
		"associated_hostnames": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The hostnames that will be prompted for this certificate.",
		},
		"fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
