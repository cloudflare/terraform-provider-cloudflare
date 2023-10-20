package sdkv2provider

import (
	"github.com/MakeNowJust/heredoc/v2"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAPIShieldOperationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"method": {
			Description: "The HTTP method used to access the endpoint",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"host": {
			Description: "RFC3986-compliant host",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"endpoint": {
			Description: heredoc.Doc("The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with `{varN}`, starting with `{var1}`. This will then be [Cloudflare-normalized](https://developers.cloudflare.com/rules/normalization/how-it-works/)"),
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
	}
}
