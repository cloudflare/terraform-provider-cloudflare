// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDNSLocationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloudflare Zero Trust Secure DNS Locations Write",
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Specify the location name.",
				Required:    true,
			},
			"endpoints": schema.SingleNestedAttribute{
				Description: "Configure the destination endpoints for this location.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"doh": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the DOH endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "Specify the list of allowed source IP network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationEndpointsDOHNetworksModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "Specify the IP address or IP CIDR.",
											Required:    true,
										},
									},
								},
							},
							"require_token": schema.BoolAttribute{
								Description: "Specify whether the DOH endpoint requires user identity authentication.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"dot": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the DOT endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "Specify the list of allowed source IP network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationEndpointsDOTNetworksModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "Specify the IP address or IP CIDR.",
											Required:    true,
										},
									},
								},
							},
						},
					},
					"ipv4": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the IPv4 endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the IPV6 endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "Specify the list of allowed source IPv6 network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationEndpointsIPV6NetworksModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "Specify the IPv6 address or IPv6 CIDR.",
											Required:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"client_default": schema.BoolAttribute{
				Description: "Indicate whether this location is the default location.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"dns_destination_ips_id": schema.StringAttribute{
				Description: "Specify the identifier of the pair of IPv4 addresses assigned to this location. When creating a location, if this field is absent or set to null, the pair of shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned. When updating a location, if this field is absent or set to null, the pre-assigned pair remains unchanged.",
				Computed:    true,
				Optional:    true,
			},
			"ecs_support": schema.BoolAttribute{
				Description: "Indicate whether the location must resolve EDNS queries.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"max_ttl": schema.SingleNestedAttribute{
				Description: "Controls how DNS response TTLs are capped for this location relative to the account `max_ttl_secs` setting. Omitting `max_ttl` on update resets it to `inherit`.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDNSLocationMaxTTLModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "`inherit` uses the account `max_ttl_secs`. `override` uses this location's `ttl_secs`. `disabled` leaves returned TTLs unchanged.\nAvailable values: \"inherit\", \"override\", \"disabled\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"inherit",
								"override",
								"disabled",
							),
						},
					},
					"ttl_secs": schema.Int64Attribute{
						Description: "Location-specific cap on DNS response TTLs, in seconds. Required when `mode` is `override`. Must be omitted when `mode` is `inherit` or `disabled`.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(60, 36000),
						},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Description: "Specify the list of network ranges from which requests at this location originate. The list takes effect only if it is non-empty and the IPv4 endpoint is enabled for this location.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationNetworksModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network": schema.StringAttribute{
							Description: "Specify the IPv4 address or IPv4 CIDR. Limit IPv4 CIDRs to a maximum of /24.",
							Required:    true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"dns_destination_ipv6_block_id": schema.StringAttribute{
				Description: "Specify the UUID of the IPv6 block brought to the gateway so that this location's IPv6 address is allocated from the Bring Your Own IPv6 (BYOIPv6) block rather than the standard Cloudflare IPv6 block.",
				Computed:    true,
			},
			"doh_subdomain": schema.StringAttribute{
				Description: "Specify the DNS over HTTPS domain that receives DNS requests. Gateway automatically generates this value.",
				Computed:    true,
			},
			"ip": schema.StringAttribute{
				Description: "Defines the automatically generated IPv6 destination IP assigned to this location. Gateway counts all DNS requests sent to this IP as requests under this location.",
				Computed:    true,
			},
			"ipv4_destination": schema.StringAttribute{
				Description: "Show the primary destination IPv4 address from the pair identified dns_destination_ips_id. This field read-only.",
				Computed:    true,
			},
			"ipv4_destination_backup": schema.StringAttribute{
				Description: "Show the backup destination IPv4 address from the pair identified dns_destination_ips_id. This field read-only.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustDNSLocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDNSLocationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
