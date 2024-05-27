// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_organization

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r AccessOrganizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"auth_domain": schema.StringAttribute{
				Description: "The unique subdomain assigned to your Zero Trust organization.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description:   "The name of your Zero Trust organization.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.",
				Optional:    true,
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_ui_read_only": schema.BoolAttribute{
				Description: "Lock all settings as Read-Only in the Dashboard, regardless of user permission. Updates may only be made via the API or Terraform for this account when enabled.",
				Optional:    true,
			},
			"login_design": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"background_color": schema.StringAttribute{
						Description: "The background color on your login page.",
						Optional:    true,
					},
					"footer_text": schema.StringAttribute{
						Description: "The text at the bottom of your login page.",
						Optional:    true,
					},
					"header_text": schema.StringAttribute{
						Description: "The text at the top of your login page.",
						Optional:    true,
					},
					"logo_path": schema.StringAttribute{
						Description: "The URL of the logo on your login page.",
						Optional:    true,
					},
					"text_color": schema.StringAttribute{
						Description: "The text color on your login page.",
						Optional:    true,
					},
				},
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for applications will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
				Optional:    true,
			},
			"ui_read_only_toggle_reason": schema.StringAttribute{
				Description: "A description of the reason why the UI read only field is being toggled.",
				Optional:    true,
			},
			"user_seat_expiration_inactive_time": schema.StringAttribute{
				Description: "The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.",
				Optional:    true,
			},
			"warp_auth_session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for applications will be valid. Must be in the format `30m` or `2h45m`. Valid time units are: m, h.",
				Optional:    true,
			},
		},
	}
}
