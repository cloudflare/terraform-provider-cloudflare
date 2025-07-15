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
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cloudflare_gre_endpoint": schema.StringAttribute{
				Description: "The IP address assigned to the Cloudflare side of the GRE tunnel.",
				Required:    true,
			},
			"customer_gre_endpoint": schema.StringAttribute{
				Description: "The IP address assigned to the customer side of the GRE tunnel.",
				Required:    true,
			},
			"interface_address": schema.StringAttribute{
				Description: "A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel. Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the tunnel. The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description of the GRE tunnel.",
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
						Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel.\nAvailable values: \"unidirectional\", \"bidirectional\".",
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
						Description: "How frequent the health check is run. The default value is `mid`.\nAvailable values: \"low\", \"mid\", \"high\".",
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
					"target": schema.SingleNestedAttribute{
						Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target. Must be in object form if the x-magic-new-hc-target header is set to true and string form if x-magic-new-hc-target is absent or set to false.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[MagicWANGRETunnelHealthCheckTargetModel](ctx),
						Attributes: map[string]schema.Attribute{
							"effective": schema.StringAttribute{
								Description: "The effective health check target. If 'saved' is empty, then this field will be populated with the calculated default value on GET requests. Ignored in POST, PUT, and PATCH requests.",
								Computed:    true,
							},
							"saved": schema.StringAttribute{
								Description: "The saved health check target. Setting the value to the empty string indicates that the calculated default value will be used.",
								Optional:    true,
							},
						},
					},
					"type": schema.StringAttribute{
						Description: "The type of healthcheck to run, reply or request. The default value is `reply`.\nAvailable values: \"reply\", \"request\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("reply", "request"),
						},
						Default: stringdefault.StaticString("reply"),
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "The date and time the tunnel was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.BoolAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The date and time the tunnel was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"gre_tunnel": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelGRETunnelModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
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
								Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel.\nAvailable values: \"unidirectional\", \"bidirectional\".",
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
								Description: "How frequent the health check is run. The default value is `mid`.\nAvailable values: \"low\", \"mid\", \"high\".",
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
							"target": schema.SingleNestedAttribute{
								Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target. Must be in object form if the x-magic-new-hc-target header is set to true and string form if x-magic-new-hc-target is absent or set to false.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[MagicWANGRETunnelGRETunnelHealthCheckTargetModel](ctx),
								Attributes: map[string]schema.Attribute{
									"effective": schema.StringAttribute{
										Description: "The effective health check target. If 'saved' is empty, then this field will be populated with the calculated default value on GET requests. Ignored in POST, PUT, and PATCH requests.",
										Computed:    true,
									},
									"saved": schema.StringAttribute{
										Description: "The saved health check target. Setting the value to the empty string indicates that the calculated default value will be used.",
										Computed:    true,
									},
								},
							},
							"type": schema.StringAttribute{
								Description: "The type of healthcheck to run, reply or request. The default value is `reply`.\nAvailable values: \"reply\", \"request\".",
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
			"modified_gre_tunnel": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANGRETunnelModifiedGRETunnelModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
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
								Description: "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel.\nAvailable values: \"unidirectional\", \"bidirectional\".",
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
								Description: "How frequent the health check is run. The default value is `mid`.\nAvailable values: \"low\", \"mid\", \"high\".",
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
							"target": schema.SingleNestedAttribute{
								Description: "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target. Must be in object form if the x-magic-new-hc-target header is set to true and string form if x-magic-new-hc-target is absent or set to false.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[MagicWANGRETunnelModifiedGRETunnelHealthCheckTargetModel](ctx),
								Attributes: map[string]schema.Attribute{
									"effective": schema.StringAttribute{
										Description: "The effective health check target. If 'saved' is empty, then this field will be populated with the calculated default value on GET requests. Ignored in POST, PUT, and PATCH requests.",
										Computed:    true,
									},
									"saved": schema.StringAttribute{
										Description: "The saved health check target. Setting the value to the empty string indicates that the calculated default value will be used.",
										Computed:    true,
									},
								},
							},
							"type": schema.StringAttribute{
								Description: "The type of healthcheck to run, reply or request. The default value is `reply`.\nAvailable values: \"reply\", \"request\".",
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
	resp.Schema = customResourceSchema(ctx)
}

func (r *MagicWANGRETunnelResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
