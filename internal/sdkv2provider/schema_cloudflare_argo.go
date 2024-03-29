package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareArgoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"tiered_caching": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
			Optional:     true,
			Description:  fmt.Sprintf("Whether tiered caching is enabled. %s", renderAvailableDocumentationValuesStringSlice([]string{"on", "off"})),
			Deprecated:   "tiered_caching has been deprecated in favour of using `cloudflare_tiered_cache` resource instead.",
		},
		"smart_routing": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
			Optional:     true,
			Description:  fmt.Sprintf("Whether smart routing is enabled. %s", renderAvailableDocumentationValuesStringSlice([]string{"on", "off"})),
		},
	}
}
