// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ipsec_tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r IPSECTunnelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ipsec_tunnel_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"cloudflare_endpoint": schema.StringAttribute{
				Description: "The IP address assigned to the Cloudflare side of the IPsec tunnel.",
				Required:    true,
			},
			"interface_address": schema.StringAttribute{
				Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the IPsec tunnel. The name cannot share a name with other tunnels.",
				Required:    true,
			},
			"customer_endpoint": schema.StringAttribute{
				Description: "The IP address assigned to the customer side of the IPsec tunnel. Not required, but must be set for proactive traceroutes to work.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description forthe IPsec tunnel.",
				Optional:    true,
			},
			"health_check": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel. Note in the case of bidirecitonal healthchecks, the target field in health_check is ignored as the interface_address is used to send traffic into the tunnel.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("unidirectional", "bidirectional"),
						},
						Default: stringdefault.StaticString("unidirectional"),
					},
					"enabled": schema.BoolAttribute{
						Description: "Determines whether to run healthchecks for a tunnel.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
					},
					"rate": schema.StringAttribute{
						Description: "How frequent the health check is run. The default value is `mid`.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("low", "mid", "high"),
						},
						Default: stringdefault.StaticString("mid"),
					},
					"target": schema.StringAttribute{
						Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target.",
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("reply", "request"),
						},
						Default: stringdefault.StaticString("reply"),
					},
				},
			},
			"psk": schema.StringAttribute{
				Description: "A randomly generated or provided string for use in the IPsec tunnel.",
				Optional:    true,
			},
			"replay_protection": schema.BoolAttribute{
				Description: "If `true`, then IPsec replay protection will be supported in the Cloudflare-to-customer direction.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}
