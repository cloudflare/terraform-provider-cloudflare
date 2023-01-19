package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var healthcheckRegions = []string{"WNAM", "ENAM", "WEU", "EEU", "NSAM", "SSAM", "OC", "ME", "NAF", "SAF", "IN", "SEAS", "NEAS", "ALL_REGIONS"}
var healthcheckType = []string{"TCP", "HTTP", "HTTPS"}
var healthcheckMethod = []string{"connection_established", "GET", "HEAD"}

func resourceCloudflareHealthcheckSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "A short name to identify the health check. Only alphanumeric characters, hyphens, and underscores are allowed.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "A human-readable description of the health check.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"suspended": {
			Description: "If suspended, no health checks are sent to the origin.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"address": {
			Description: "The hostname or IP address of the origin server to run health checks on.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"consecutive_fails": {
			Description: "The number of consecutive fails required from a health check before changing the health to unhealthy.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
		},
		"consecutive_successes": {
			Description: "The number of consecutive successes required from a health check before changing the health to healthy.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
		},
		"retries": {
			Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     2,
		},
		"timeout": {
			Description: "The timeout (in seconds) before marking the health check as failed.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     5,
		},
		"interval": {
			Description: "The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase the load on the origin as we check from multiple locations.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60,
		},
		"check_regions": {
			Description: fmt.Sprintf("A list of regions from which to run health checks. If not set, Cloudflare will pick a default region. %s", renderAvailableDocumentationValuesStringSlice(healthcheckRegions)),
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,

			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(healthcheckRegions, false),
			},
		},
		"type": {
			Description:  fmt.Sprintf("The protocol to use for the health check. %s", renderAvailableDocumentationValuesStringSlice(healthcheckType)),
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(healthcheckType, false),
		},
		"method": {
			Description:  fmt.Sprintf("The HTTP method to use for the health check. %s", renderAvailableDocumentationValuesStringSlice(healthcheckMethod)),
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice(healthcheckMethod, false),
		},
		"port": {
			Description:  "Port number to connect to for the health check.",
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      80,
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"path": {
			Description: "The endpoint path to health check against.",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "/",
		},
		"expected_codes": {
			Description: "The expected HTTP response codes (e.g. '200') or code ranges (e.g. '2xx' for all codes starting with 2) of the health check.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"expected_body": {
			Description: "A case-insensitive sub-string to look for in the response body. If this string is not found the origin will be marked as unhealthy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"follow_redirects": {
			Description: "Follow redirects if the origin returns a 3xx status code.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"allow_insecure": {
			Description: "Do not validate the certificate when the health check uses HTTPS.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"header": {
			Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"header": {
						Description: "The header name.",
						Type:        schema.TypeString,
						Required:    true,
					},
					"values": {
						Description: "A list of string values for the header.",
						Type:        schema.TypeSet,
						Required:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		"created_on": {
			Description: "Creation time.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"modified_on": {
			Description: "Last modified time.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
