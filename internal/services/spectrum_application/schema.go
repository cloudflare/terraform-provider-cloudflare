// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.ResourceWithConfigValidators = (*SpectrumApplicationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"protocol": schema.StringAttribute{
				Description: "The port configuration at Cloudflare's edge. May specify a single port, for example `\"tcp/1000\"`, or a range of ports, for example `\"tcp/1000-2000\"`.",
				Required:    true,
			},
			"dns": schema.SingleNestedAttribute{
				Description: "The name and type of DNS record for the Spectrum application.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "The name of the DNS record associated with the application.",
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of DNS record associated with the application.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("CNAME", "ADDRESS"),
						},
					},
				},
			},
			"ip_firewall": schema.BoolAttribute{
				Description: "Enables IP Access Rules for this application.\nNotes: Only available for TCP applications.",
				Optional:    true,
			},
			"tls": schema.StringAttribute{
				Description: "The type of TLS termination associated with the application.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"flexible",
						"full",
						"strict",
					),
				},
			},
			"origin_direct": schema.ListAttribute{
				Description: "List of origin IP addresses. Array may contain multiple IP addresses for load balancing.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"origin_port": schema.DynamicAttribute{
				Description: "The destination port at the origin. Only specified in conjunction with origin_dns. May use an integer to specify a single origin port, for example `1000`, or a string to specify a range of origin ports, for example `\"1000-2000\"`.\nNotes: If specifying a port range, the number of ports in the range must match the number of ports specified in the \"protocol\" field.",
				Optional:    true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.Int64Type{}, basetypes.StringType{}),
				},
			},
			"argo_smart_routing": schema.BoolAttribute{
				Description: "Enables Argo Smart Routing for this application.\nNotes: Only available for TCP applications with traffic_type set to \"direct\".",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"proxy_protocol": schema.StringAttribute{
				Description: "Enables Proxy Protocol to the origin. Refer to [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/) for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple Proxy Protocol.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"v1",
						"v2",
						"simple",
					),
				},
				Default: stringdefault.StaticString("off"),
			},
			"traffic_type": schema.StringAttribute{
				Description: "Determines how data travels from the edge to your origin. When set to \"direct\", Spectrum will send traffic directly to your origin, and the application's type is derived from the `protocol`. When set to \"http\" or \"https\", Spectrum will apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and the application type matches this property exactly.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"direct",
						"http",
						"https",
					),
				},
				Default: stringdefault.StaticString("direct"),
			},
			"edge_ips": schema.SingleNestedAttribute{
				Description: "The anycast edge IP configuration for the hostname of this application.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[SpectrumApplicationEdgeIPsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"connectivity": schema.StringAttribute{
						Description: "The IP versions supported for inbound connections on Spectrum anycast IPs.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"all",
								"ipv4",
								"ipv6",
							),
						},
					},
					"type": schema.StringAttribute{
						Description: "The type of edge IP configuration specified. Dynamically allocated edge IPs use Spectrum anycast IPs in accordance with the connectivity you specify. Only valid with CNAME DNS names.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("dynamic", "static"),
						},
					},
					"ips": schema.ListAttribute{
						Description: "The array of customer owned IPs we broadcast via anycast for this hostname and application.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"origin_dns": schema.SingleNestedAttribute{
				Description: "The name and type of DNS record for the Spectrum application.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[SpectrumApplicationOriginDNSModel](ctx),
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "The name of the DNS record associated with the origin.",
						Optional:    true,
					},
					"ttl": schema.Int64Attribute{
						Description: "The TTL of our resolution of your DNS record in seconds.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(600),
						},
					},
					"type": schema.StringAttribute{
						Description: "The type of DNS record associated with the origin. \"\" is used to specify a combination of A/AAAA records.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"",
								"A",
								"AAAA",
								"SRV",
							),
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the Application was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the Application was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *SpectrumApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SpectrumApplicationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
