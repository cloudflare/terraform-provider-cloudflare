// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r TeamsLocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the location.",
				Required:    true,
			},
			"client_default": schema.BoolAttribute{
				Description: "True if the location is the default location.",
				Optional:    true,
			},
			"dns_destination_ips_id": schema.StringAttribute{
				Description: "The identifier of the pair of IPv4 addresses assigned to this location. When creating a location, if this field is absent or set with null, the pair of shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned. When updating a location, if the field is absent or set with null, the pre-assigned pair remains unchanged.",
				Optional:    true,
			},
			"ecs_support": schema.BoolAttribute{
				Description: "True if the location needs to resolve EDNS queries.",
				Optional:    true,
			},
			"endpoints": schema.SingleNestedAttribute{
				Description: "The destination endpoints configured for this location. When updating a location, if this field is absent or set with null, the endpoints configuration remains unchanged.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"doh": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "A list of allowed source IP network ranges for this endpoint. When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "The IP address or IP CIDR.",
											Required:    true,
										},
									},
								},
							},
							"require_token": schema.BoolAttribute{
								Description: "True if the endpoint requires [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user) authentication.",
								Optional:    true,
							},
						},
					},
					"dot": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "A list of allowed source IP network ranges for this endpoint. When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "The IP address or IP CIDR.",
											Required:    true,
										},
									},
								},
							},
						},
					},
					"ipv4": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Optional:    true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "A list of allowed source IPv6 network ranges for this endpoint. When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "The IPv6 address or IPv6 CIDR.",
											Required:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Description: "A list of network ranges that requests from this location would originate from. A non-empty list is only effective if the ipv4 endpoint is enabled for this location.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network": schema.StringAttribute{
							Description: "The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.",
							Required:    true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"doh_subdomain": schema.StringAttribute{
				Description: "The DNS over HTTPS domain to send DNS requests to. This field is auto-generated by Gateway.",
				Computed:    true,
			},
			"ip": schema.StringAttribute{
				Description: "IPV6 destination ip assigned to this location. DNS requests sent to this IP will counted as the request under this location. This field is auto-generated by Gateway.",
				Computed:    true,
			},
			"ipv4_destination": schema.StringAttribute{
				Description: "The primary destination IPv4 address from the pair identified by the dns_destination_ips_id. This field is read-only.",
				Computed:    true,
			},
			"ipv4_destination_backup": schema.StringAttribute{
				Description: "The backup destination IPv4 address from the pair identified by the dns_destination_ips_id. This field is read-only.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
