package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareDeviceManagedNetworksSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"tls"}, false),
			Description:  fmt.Sprintf("The type of Device Managed Network. %s", renderAvailableDocumentationValuesStringSlice([]string{"tls"})),
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the Device Managed Network. Must be unique.",
		},
		"config": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "The configuration containing information for the WARP client to detect the managed network.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"tls_sockaddr": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "A network address of the form \"host:port\" that the WARP client will use to detect the presence of a TLS host.",
					},
					"sha256": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.",
					},
				},
			},
		},
	}
}
