package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomSslSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"custom_ssl_priority": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"priority": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"custom_ssl_options": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			ForceNew:    true,
			Description: "The certificate associated parameters.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"certificate": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Certificate certificate and the intermediate(s).",
					},
					"private_key": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						Description: "Certificate's private key.",
					},
					"bundle_method": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"ubiquitous", "optimal", "force"}, false),
						Description:  fmt.Sprintf("Method of building intermediate certificate chain. A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. %s", renderAvailableDocumentationValuesStringSlice([]string{"ubiquitous", "optimal", "force"})),
					},
					"geo_restrictions": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"us", "eu", "highest_security"}, false),
						Description:  fmt.Sprintf("Specifies the region where your private key can be held locally. %s", renderAvailableDocumentationValuesStringSlice([]string{"us", "eu", "highest_security"})),
					},
					"type": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"legacy_custom", "sni_custom"}, false),
						Description:  fmt.Sprintf("Whether to enable support for legacy clients which do not include SNI in the TLS handshake. %s", renderAvailableDocumentationValuesStringSlice([]string{"legacy_custom", "sni_custom"})),
					},
				},
			},
		},
		"hosts": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"issuer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"signature": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uploaded_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
