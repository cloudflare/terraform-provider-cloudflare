// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitSiteACLResource)(nil)

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
			"site_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the ACL.",
				Required:    true,
			},
			"lan_1": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"lan_id": schema.StringAttribute{
						Description: "The identifier for the LAN you want to create an ACL policy with.",
						Required:    true,
					},
					"lan_name": schema.StringAttribute{
						Description: "The name of the LAN based on the provided lan_id.",
						Optional:    true,
					},
					"port_ranges": schema.ListAttribute{
						Description: "Array of port ranges on the provided LAN that will be included in the ACL. If no ports or port rangess are provided, communication on any port on this LAN is allowed.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ports": schema.ListAttribute{
						Description: "Array of ports on the provided LAN that will be included in the ACL. If no ports or port ranges are provided, communication on any port on this LAN is allowed.",
						Optional:    true,
						ElementType: types.Int64Type,
					},
					"subnets": schema.ListAttribute{
						Description: "Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"lan_2": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"lan_id": schema.StringAttribute{
						Description: "The identifier for the LAN you want to create an ACL policy with.",
						Required:    true,
					},
					"lan_name": schema.StringAttribute{
						Description: "The name of the LAN based on the provided lan_id.",
						Optional:    true,
					},
					"port_ranges": schema.ListAttribute{
						Description: "Array of port ranges on the provided LAN that will be included in the ACL. If no ports or port rangess are provided, communication on any port on this LAN is allowed.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ports": schema.ListAttribute{
						Description: "Array of ports on the provided LAN that will be included in the ACL. If no ports or port ranges are provided, communication on any port on this LAN is allowed.",
						Optional:    true,
						ElementType: types.Int64Type,
					},
					"subnets": schema.ListAttribute{
						Description: "Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "Description for the ACL.",
				Optional:    true,
			},
			"forward_locally": schema.BoolAttribute{
				Description: `The desired forwarding action for this ACL policy. If set to "false", the policy will forward traffic to Cloudflare. If set to "true", the policy will forward traffic locally on the Magic Connector. If not included in request, will default to false.`,
				Optional:    true,
			},
			"unidirectional": schema.BoolAttribute{
				Description: `The desired traffic direction for this ACL policy. If set to "false", the policy will allow bidirectional traffic. If set to "true", the policy will only allow traffic in one direction. If not included in request, will default to false.`,
				Optional:    true,
			},
			"protocols": schema.ListAttribute{
				Optional: true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"tcp",
							"udp",
							"icmp",
						),
					),
				},
				ElementType: types.StringType,
			},
		},
	}
}

func (r *MagicTransitSiteACLResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitSiteACLResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
