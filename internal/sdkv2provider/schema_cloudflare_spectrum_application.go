package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareSpectrumApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},

		"protocol": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The port configuration at Cloudflare's edge. e.g. `tcp/22`.",
		},

		"traffic_type": {
			Type:         schema.TypeString,
			Default:      "direct",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"direct", "http", "https"}, false),
			Description:  fmt.Sprintf("Sets application type. %s", renderAvailableDocumentationValuesStringSlice([]string{"direct", "http", "https"})),
		},

		"dns": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The name and type of DNS record for the Spectrum application.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The type of DNS record associated with the application.",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The name of the DNS record associated with the application",
					},
				},
			},
		},

		"origin_direct": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of destination addresses to the origin. e.g. `tcp://192.0.2.1:22`.",
		},

		"origin_dns": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "A destination DNS addresses to the origin.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Fully qualified domain name of the origin.",
					},
				},
			},
		},

		"origin_port": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"origin_port_range"},
			ValidateFunc:  validation.IntBetween(0, 65535),
			Description:   "Origin port to proxy traffice to.",
		},

		"origin_port_range": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			ConflictsWith: []string{"origin_port"},
			Description:   "Origin port range to proxy traffice to. When using a range, the protocol field must also specify a range, e.g. `tcp/22-23`.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"start": {
						Type:         schema.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(0, 65535),
						Description:  "Lower bound of the origin port range.",
					},
					"end": {
						Type:         schema.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(0, 65535),
						Description:  "Upper bound of the origin port range.",
					},
				},
			},
		},

		"tls": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "off",
			ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict"}, false),
			Description:  fmt.Sprintf("TLS configuration option for Cloudflare to connect to your origin. %s", renderAvailableDocumentationValuesStringSlice([]string{"off", "flexible", "full", "strict"})),
		},

		"ip_firewall": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Enables the IP Firewall for this application.",
		},

		"proxy_protocol": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "off",
			Description:  fmt.Sprintf("Enables a proxy protocol to the origin. %s", renderAvailableDocumentationValuesStringSlice([]string{"off", "v1", "v2", "simple"})),
			ValidateFunc: validation.StringInSlice([]string{"off", "v1", "v2", "simple"}, false),
		},

		"edge_ips": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The anycast edge IP configuration for the hostname of this application.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"dynamic", "static"}, false),
						Description:  fmt.Sprintf("The type of edge IP configuration specified. %s", renderAvailableDocumentationValuesStringSlice([]string{"dynamic", "static"})),
					},
					"connectivity": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"all", "ipv4", "ipv6"}, false),
						Description:  fmt.Sprintf("The IP versions supported for inbound connections on Spectrum anycast IPs. Required when `type` is not `static`. %s", renderAvailableDocumentationValuesStringSlice([]string{"all", "ipv4", "ipv6"})),
					},
					"ips": {
						Type:        schema.TypeSet,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Optional:    true,
						Description: "The collection of customer owned IPs to broadcast via anycast for this hostname and application. Requires [Bring Your Own IP](https://developers.cloudflare.com/spectrum/getting-started/byoip/) provisioned.",
					},
				},
			},
		},

		"argo_smart_routing": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enables Argo Smart Routing.",
		},
	}
}
