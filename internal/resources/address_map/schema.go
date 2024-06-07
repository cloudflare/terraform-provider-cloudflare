// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r AddressMapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description field which may be used to describe the types of IPs or zones on the map.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the Address Map is enabled or not. Cloudflare's DNS will not respond with IP addresses on an Address Map until the map is enabled.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"ips": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"memberships": schema.ListNestedAttribute{
				Description: "Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"can_delete": schema.BoolAttribute{
							Description: "Controls whether the membership can be deleted via the API or not.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"identifier": schema.StringAttribute{
							Description: "The identifier for the membership (eg. a zone or account tag).",
							Optional:    true,
						},
						"kind": schema.StringAttribute{
							Description: "The type of the membership.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("zone", "account"),
							},
						},
					},
				},
			},
		},
	}
}
