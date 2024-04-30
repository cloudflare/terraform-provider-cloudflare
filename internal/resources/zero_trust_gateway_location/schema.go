// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r ZeroTrustGatewayLocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the location.",
				Required:    true,
			},
			"client_default": schema.BoolAttribute{
				Description: "True if the location is the default location.",
				Optional:    true,
			},
			"ecs_support": schema.BoolAttribute{
				Description: "True if the location needs to resolve EDNS queries.",
				Optional:    true,
			},
			"networks": schema.ListNestedAttribute{
				Description: "A list of network ranges that requests from this location would originate from.",
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
		},
	}
}
