// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_ipsec_tunnels

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r MagicTransitIPSECTunnelsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"tunnel_identifier": schema.StringAttribute{
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
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("unidirectional", "bidirectional"),
						},
					},
					"enabled": schema.BoolAttribute{
						Description: "Determines whether to run healthchecks for a tunnel.",
						Optional:    true,
					},
					"rate": schema.StringAttribute{
						Description: "How frequent the health check is run. The default value is `mid`.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("low", "mid", "high"),
						},
					},
					"target": schema.StringAttribute{
						Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target.",
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("reply", "request"),
						},
					},
				},
			},
			"psk": schema.StringAttribute{
				Description: "A randomly generated or provided string for use in the IPsec tunnel.",
				Optional:    true,
			},
			"replay_protection": schema.BoolAttribute{
				Description: "If `true`, then IPsec replay protection will be supported in the Cloudflare-to-customer direction.",
				Optional:    true,
			},
			"ipsec_tunnels": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
						"id": schema.StringAttribute{
							Description: "Tunnel identifier tag.",
							Computed:    true,
						},
						"allow_null_cipher": schema.BoolAttribute{
							Description: "When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel (Phase 2).",
							Optional:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "The date and time the tunnel was created.",
							Computed:    true,
						},
						"customer_endpoint": schema.StringAttribute{
							Description: "The IP address assigned to the customer side of the IPsec tunnel. Not required, but must be set for proactive traceroutes to work.",
							Optional:    true,
						},
						"description": schema.StringAttribute{
							Description: "An optional description forthe IPsec tunnel.",
							Optional:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "The date and time the tunnel was last modified.",
							Computed:    true,
						},
						"psk_metadata": schema.SingleNestedAttribute{
							Description: "The PSK metadata that includes when the PSK was generated.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"last_generated_on": schema.StringAttribute{
									Description: "The date and time the tunnel was last modified.",
									Computed:    true,
								},
							},
						},
						"replay_protection": schema.BoolAttribute{
							Description: "If `true`, then IPsec replay protection will be supported in the Cloudflare-to-customer direction.",
							Optional:    true,
						},
						"tunnel_health_check": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Determines whether to run healthchecks for a tunnel.",
									Optional:    true,
								},
								"rate": schema.StringAttribute{
									Description: "How frequent the health check is run. The default value is `mid`.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("low", "mid", "high"),
									},
								},
								"target": schema.StringAttribute{
									Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`.",
									Optional:    true,
								},
								"type": schema.StringAttribute{
									Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("reply", "request"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
