package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareDevicePostureIntegrationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the device posture integration.",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{ws1, uptycs, crowdstrike, intune, kolide, sentinelone}, false),
			Description:  fmt.Sprintf("The device posture integration type. %s", renderAvailableDocumentationValuesStringSlice([]string{ws1, uptycs, crowdstrike, intune, kolide, sentinelone, tanium})),
		},
		"identifier": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"interval": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Indicates the frequency with which to poll the third-party API. Must be in the format `1h` or `30m`.",
		},
		"config": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The device posture integration's connection authorization parameters.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"auth_url": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsURLWithHTTPS,
						Description:  "The third-party authorization API URL.",
					},
					"api_url": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsURLWithHTTPS,
						Description:  "The third-party API's URL.",
					},
					"client_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The client identifier for authenticating API calls.",
					},
					"client_secret": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						Description: "The client secret for authenticating API calls.",
					},
					"customer_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The customer identifier for authenticating API calls.",
					},
					"client_key": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						Description: "The client key for authenticating API calls.",
					},
				},
			},
		},
	}
}
