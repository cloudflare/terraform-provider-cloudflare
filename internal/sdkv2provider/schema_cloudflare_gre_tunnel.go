package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareGRETunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the GRE tunnel.",
		},
		"customer_gre_endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The IP address assigned to the customer side of the GRE tunnel.",
		},
		"cloudflare_gre_endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
		},
		"interface_address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the GRE tunnel intent",
		},
		"ttl": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Time To Live (TTL) in number of hops of the GRE tunnel.",
			ValidateFunc: validation.IntAtLeast(64),
		},
		"mtu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Maximum Transmission Unit (MTU) in bytes for the GRE tunnel.",
			ValidateFunc: validation.IntBetween(576, 1476),
		},
		"health_check_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Specifies if ICMP tunnel health checks are enabled.",
		},
		"health_check_target": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the customer endpoint that will receive tunnel health checks.",
		},
		"health_check_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice([]string{"request", "reply"}, false),
			Description:  fmt.Sprintf("Specifies the ICMP echo type for the health check. %s", renderAvailableDocumentationValuesStringSlice([]string{"request", "reply"})),
		},
	}
}
