package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareWeb3HostnameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "The hostname that will point to the target gateway via CNAME.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"target": {
			Description:  "Target gateway of the hostname.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"ethereum", "ipfs"}, false),
		},
		"description": {
			Description:  "An optional description of the hostname.",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 500),
		},
		"dnslink": {
			Description: "DNSLink value used if the target is ipfs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status": {
			Description: "Status of the hostname's activation.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_on": {
			Description: "Creation time.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"modified_on": {
			Description: "Last modification time.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
