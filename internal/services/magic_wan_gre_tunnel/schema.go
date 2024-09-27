// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MagicWANGRETunnelResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"gre_tunnel_id": schema.StringAttribute{
				Description:   "Identifier",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cloudflare_gre_endpoint": schema.StringAttribute{
				Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
				Optional:    true,
			},
			"customer_gre_endpoint": schema.StringAttribute{
				Description: "The IP address assigned to the customer side of the GRE tunnel.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description of the GRE tunnel.",
				Optional:    true,
			},
			"interface_address": schema.StringAttribute{
				Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the tunnel. The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.",
				Optional:    true,
			},
			"mtu": schema.Int64Attribute{
				Description: "Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value is 576.",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(1476),
			},
			"ttl": schema.Int64Attribute{
				Description: "Time To Live (TTL) in number of hops of the GRE tunnel.",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(64),
			},
			"health_check": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelHealthCheckModel](ctx),
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
							stringvalidator.OneOfCaseInsensitive(
								"low",
								"mid",
								"high",
							),
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
			"deleted": schema.BoolAttribute{
				Computed: true,
			},
			"modified": schema.BoolAttribute{
				Computed: true,
			},
			"deleted_gre_tunnel": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelDeletedGRETunnelModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cloudflare_gre_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
						Computed:    true,
					},
					"customer_gre_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the customer side of the GRE tunnel.",
						Computed:    true,
					},
					"interface_address": schema.StringAttribute{
						Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the tunnel. The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Tunnel identifier tag.",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "The date and time the tunnel was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An optional description of the GRE tunnel.",
						Computed:    true,
					},
					"health_check": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelDeletedGRETunnelHealthCheckModel](ctx),
						Attributes: map[string]schema.Attribute{
							"direction": schema.StringAttribute{
								Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel. Note in the case of bidirecitonal healthchecks, the target field in health_check is ignored as the interface_address is used to send traffic into the tunnel.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("unidirectional", "bidirectional"),
								},
								Default: stringdefault.StaticString("unidirectional"),
							},
							"enabled": schema.BoolAttribute{
								Description: "Determines whether to run healthchecks for a tunnel.",
								Computed:    true,
								Default:     booldefault.StaticBool(true),
							},
							"rate": schema.StringAttribute{
								Description: "How frequent the health check is run. The default value is `mid`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"low",
										"mid",
										"high",
									),
								},
								Default: stringdefault.StaticString("mid"),
							},
							"target": schema.StringAttribute{
								Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target.",
								Computed:    true,
							},
							"type": schema.StringAttribute{
								Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("reply", "request"),
								},
								Default: stringdefault.StaticString("reply"),
							},
						},
					},
					"modified_on": schema.StringAttribute{
						Description: "The date and time the tunnel was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"mtu": schema.Int64Attribute{
						Description: "Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value is 576.",
						Computed:    true,
						Default:     int64default.StaticInt64(1476),
					},
					"ttl": schema.Int64Attribute{
						Description: "Time To Live (TTL) in number of hops of the GRE tunnel.",
						Computed:    true,
						Default:     int64default.StaticInt64(64),
					},
				},
			},
			"gre_tunnel": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelGRETunnelModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cloudflare_gre_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
						Computed:    true,
					},
					"customer_gre_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the customer side of the GRE tunnel.",
						Computed:    true,
					},
					"interface_address": schema.StringAttribute{
						Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the tunnel. The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Tunnel identifier tag.",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "The date and time the tunnel was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An optional description of the GRE tunnel.",
						Computed:    true,
					},
					"health_check": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelGRETunnelHealthCheckModel](ctx),
						Attributes: map[string]schema.Attribute{
							"direction": schema.StringAttribute{
								Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel. Note in the case of bidirecitonal healthchecks, the target field in health_check is ignored as the interface_address is used to send traffic into the tunnel.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("unidirectional", "bidirectional"),
								},
								Default: stringdefault.StaticString("unidirectional"),
							},
							"enabled": schema.BoolAttribute{
								Description: "Determines whether to run healthchecks for a tunnel.",
								Computed:    true,
								Default:     booldefault.StaticBool(true),
							},
							"rate": schema.StringAttribute{
								Description: "How frequent the health check is run. The default value is `mid`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"low",
										"mid",
										"high",
									),
								},
								Default: stringdefault.StaticString("mid"),
							},
							"target": schema.StringAttribute{
								Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target.",
								Computed:    true,
							},
							"type": schema.StringAttribute{
								Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("reply", "request"),
								},
								Default: stringdefault.StaticString("reply"),
							},
						},
					},
					"modified_on": schema.StringAttribute{
						Description: "The date and time the tunnel was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"mtu": schema.Int64Attribute{
						Description: "Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value is 576.",
						Computed:    true,
						Default:     int64default.StaticInt64(1476),
					},
					"ttl": schema.Int64Attribute{
						Description: "Time To Live (TTL) in number of hops of the GRE tunnel.",
						Computed:    true,
						Default:     int64default.StaticInt64(64),
					},
				},
			},
			"gre_tunnels": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[MagicWANGRETunnelGRETunnelsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cloudflare_gre_endpoint": schema.StringAttribute{
							Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
							Computed:    true,
						},
						"customer_gre_endpoint": schema.StringAttribute{
							Description: "The IP address assigned to the customer side of the GRE tunnel.",
							Computed:    true,
						},
						"interface_address": schema.StringAttribute{
							Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the tunnel. The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "Tunnel identifier tag.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "The date and time the tunnel was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "An optional description of the GRE tunnel.",
							Computed:    true,
						},
						"health_check": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelGRETunnelsHealthCheckModel](ctx),
							Attributes: map[string]schema.Attribute{
								"direction": schema.StringAttribute{
									Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel. Note in the case of bidirecitonal healthchecks, the target field in health_check is ignored as the interface_address is used to send traffic into the tunnel.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("unidirectional", "bidirectional"),
									},
									Default: stringdefault.StaticString("unidirectional"),
								},
								"enabled": schema.BoolAttribute{
									Description: "Determines whether to run healthchecks for a tunnel.",
									Computed:    true,
									Default:     booldefault.StaticBool(true),
								},
								"rate": schema.StringAttribute{
									Description: "How frequent the health check is run. The default value is `mid`.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"low",
											"mid",
											"high",
										),
									},
									Default: stringdefault.StaticString("mid"),
								},
								"target": schema.StringAttribute{
									Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("reply", "request"),
									},
									Default: stringdefault.StaticString("reply"),
								},
							},
						},
						"modified_on": schema.StringAttribute{
							Description: "The date and time the tunnel was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"mtu": schema.Int64Attribute{
							Description: "Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value is 576.",
							Computed:    true,
							Default:     int64default.StaticInt64(1476),
						},
						"ttl": schema.Int64Attribute{
							Description: "Time To Live (TTL) in number of hops of the GRE tunnel.",
							Computed:    true,
							Default:     int64default.StaticInt64(64),
						},
					},
				},
			},
			"modified_gre_tunnel": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelModifiedGRETunnelModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cloudflare_gre_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
						Computed:    true,
					},
					"customer_gre_endpoint": schema.StringAttribute{
						Description: "The IP address assigned to the customer side of the GRE tunnel.",
						Computed:    true,
					},
					"interface_address": schema.StringAttribute{
						Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the tunnel. The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Tunnel identifier tag.",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "The date and time the tunnel was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An optional description of the GRE tunnel.",
						Computed:    true,
					},
					"health_check": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelModifiedGRETunnelHealthCheckModel](ctx),
						Attributes: map[string]schema.Attribute{
							"direction": schema.StringAttribute{
								Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel. Note in the case of bidirecitonal healthchecks, the target field in health_check is ignored as the interface_address is used to send traffic into the tunnel.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("unidirectional", "bidirectional"),
								},
								Default: stringdefault.StaticString("unidirectional"),
							},
							"enabled": schema.BoolAttribute{
								Description: "Determines whether to run healthchecks for a tunnel.",
								Computed:    true,
								Default:     booldefault.StaticBool(true),
							},
							"rate": schema.StringAttribute{
								Description: "How frequent the health check is run. The default value is `mid`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"low",
										"mid",
										"high",
									),
								},
								Default: stringdefault.StaticString("mid"),
							},
							"target": schema.StringAttribute{
								Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target.",
								Computed:    true,
							},
							"type": schema.StringAttribute{
								Description: "The type of healthcheck to run, reply or request. The default value is `reply`.",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("reply", "request"),
								},
								Default: stringdefault.StaticString("reply"),
							},
						},
					},
					"modified_on": schema.StringAttribute{
						Description: "The date and time the tunnel was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"mtu": schema.Int64Attribute{
						Description: "Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value is 576.",
						Computed:    true,
						Default:     int64default.StaticInt64(1476),
					},
					"ttl": schema.Int64Attribute{
						Description: "Time To Live (TTL) in number of hops of the GRE tunnel.",
						Computed:    true,
						Default:     int64default.StaticInt64(64),
					},
				},
			},
		},
	}
}

func (r *MagicWANGRETunnelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicWANGRETunnelResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
