package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareDeviceDexTestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the Device Dex Test. Must be unique.",
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Additional details about the test.",
		},
		"interval": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "How often the test will run.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Determines whether or not the test is active.",
		},
		"data": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "The configuration object which contains the details for the WARP client to conduct the test.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"kind": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"http", "traceroute"}, false),
						Description:  fmt.Sprintf("The type of Device Dex Test. %s", renderAvailableDocumentationValuesStringSlice([]string{"http", "traceroute"})),
					},
					"method": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"GET"}, false),
						Description:  fmt.Sprintf("The http request method. %s", renderAvailableDocumentationValuesStringSlice([]string{"GET"})),
					},
					"host": {
						Type:        schema.TypeString,
						Required:    true,
						Description: fmt.Sprint("The host URL for `http` test `kind`. For `traceroute`, it must be a valid hostname or IP address."),
					},
				},
			},
		},
		"updated": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the Dex Test was last updated.",
		},
		"created": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the Dex Test was created.",
		},
	}
}
