// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitSiteWANResource)(nil)

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
			"physport": schema.Int64Attribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"priority": schema.Int64Attribute{
				Optional: true,
			},
			"vlan_tag": schema.Int64Attribute{
				Description: "VLAN ID. Use zero for untagged.",
				Optional:    true,
			},
			"static_addressing": schema.SingleNestedAttribute{
				Description: "(optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[MagicTransitSiteWANStaticAddressingModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Required:    true,
					},
					"gateway_address": schema.StringAttribute{
						Description: "A valid IPv4 address.",
						Required:    true,
					},
					"secondary_address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Optional:    true,
					},
				},
			},
			"health_check_rate": schema.StringAttribute{
				Description: "Magic WAN health check rate for tunnels created on this link. The default value is `mid`.\nAvailable values: \"low\", \"mid\", \"high\".",
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
		},
	}
}

func (r *MagicTransitSiteWANResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitSiteWANResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
