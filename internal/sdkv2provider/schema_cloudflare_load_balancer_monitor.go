package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},

		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Free text description.",
		},

		"interval": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60,
			Description: "The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.",
		},
		// interval has to be larger than (retries+1) * probe_timeout:

		"method": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The method to use for the health check.",
		},

		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 65535),
			Description:  "The port number to use for the healthcheck, required when creating a TCP monitor.",
		},

		"retries": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      2,
			ValidateFunc: validation.IntBetween(1, 5),
			Description:  "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
		},

		"timeout": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      5,
			ValidateFunc: validation.IntBetween(1, 10),
			Description:  "The timeout (in seconds) before marking the health check as failed.",
		},

		"type": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "http",
			ValidateFunc: validation.StringInSlice([]string{"http", "https", "tcp", "udp_icmp", "icmp_ping", "smtp"}, false),
			Description:  fmt.Sprintf("The protocol to use for the healthcheck. %s", renderAvailableDocumentationValuesStringSlice([]string{"http", "https", "tcp", "udp_icmp", "icmp_ping", "smtp"})),
		},

		"created_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the load balancer monitor was created.",
		},

		"modified_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the load balancer monitor was last modified.",
		},

		"allow_insecure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Do not validate the certificate when monitor use HTTPS.  Only valid if `type` is \"http\" or \"https\"",
		},

		"consecutive_down": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "To be marked unhealthy the monitored origin must fail this healthcheck N consecutive times.",
		},

		"consecutive_up": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "To be marked healthy the monitored origin must pass this healthcheck N consecutive times.",
		},

		"expected_body": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. Only valid if `type` is \"http\" or \"https\".",
		},

		"expected_codes": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The expected HTTP response code or code range of the health check. Eg `2xx`. Only valid and required if `type` is \"http\" or \"https\".",
		},

		"follow_redirects": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Follow redirects if returned by the origin. Only valid if `type` is \"http\" or \"https\".",
		},

		"header": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"header": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The header name.",
					},

					"values": {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "A list of values for the header.",
					},
				},
			},
		},

		"path": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The endpoint path to health check against.",
		},

		"probe_zone": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Assign this monitor to emulate the specified zone while probing. Only valid if `type` is \"http\" or \"https\".",
		},
	}
}
