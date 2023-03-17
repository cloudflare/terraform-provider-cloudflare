package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareListItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"list_id": {
			Description: "The list identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ip": {
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"ip", "redirect"},
			ForceNew:     true,
		},
		"redirect": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"source_url": {
						Description: "The source url of the redirect.",
						Type:        schema.TypeString,
						Required:    true,
					},
					"target_url": {
						Description: "The target url of the redirect.",
						Type:        schema.TypeString,
						Required:    true,
					},
					"include_subdomains": {
						Description:  fmt.Sprintf("Whether the redirect also matches subdomains of the source url. %s", renderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
					},
					"subpath_matching": {
						Description:  fmt.Sprintf("Whether the redirect also matches subpaths of the source url. %s", renderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
					},
					"status_code": {
						Description: "The status code to be used when redirecting a request.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"preserve_query_string": {
						Description:  fmt.Sprintf("Whether the redirect target url should keep the query string of the request's url. %s", renderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
					},
					"preserve_path_suffix": {
						Description:  fmt.Sprintf("Whether to preserve the path suffix when doing subpath matching. %s", renderAvailableDocumentationValuesStringSlice([]string{"disabled", "enabled"})),
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
					},
				},
			},
		},
		"comment": {
			Description: "An optional comment for the item.",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
