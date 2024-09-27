// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessInfrastructureTargetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Target identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description: "A non-unique field that refers to a target. Case insensitive, maximum\nlength of 255 characters, supports the use of special characters dash\nand period, does not support spaces, and must start and end with an\nalphanumeric character.",
				Required:    true,
			},
			"ip": schema.SingleNestedAttribute{
				Description: "The IPv4/IPv6 address that identifies where to reach a target",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.SingleNestedAttribute{
						Description: "The target's IPv4 address",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetIPIPV4Model](ctx),
						Attributes: map[string]schema.Attribute{
							"ip_addr": schema.StringAttribute{
								Description: "IP address of the target",
								Computed:    true,
								Optional:    true,
							},
							"virtual_network_id": schema.StringAttribute{
								Description: "Private virtual network identifier for the target",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						Description: "The target's IPv6 address",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetIPIPV6Model](ctx),
						Attributes: map[string]schema.Attribute{
							"ip_addr": schema.StringAttribute{
								Description: "IP address of the target",
								Computed:    true,
								Optional:    true,
							},
							"virtual_network_id": schema.StringAttribute{
								Description: "Private virtual network identifier for the target",
								Computed:    true,
								Optional:    true,
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Description: "Date and time at which the target was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Description: "Date and time at which the target was modified",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustAccessInfrastructureTargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessInfrastructureTargetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
