// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustOrganizationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The name of your Zero Trust organization.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The name of your Zero Trust organization.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"auth_domain": schema.StringAttribute{
				Description: "The unique subdomain assigned to your Zero Trust organization.",
				Required:    true,
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.",
				Optional:    true,
			},
			"is_ui_read_only": schema.BoolAttribute{
				Description: "Lock all settings as Read-Only in the Dashboard, regardless of user permission. Updates may only be made via the API or Terraform for this account when enabled.",
				Optional:    true,
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
				Description: "The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.",
				Optional:    true,
			},
			"warp_auth_session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for applications will be valid. Must be in the format `30m` or `2h45m`. Valid time units are: m, h.",
				Optional:    true,
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"custom_pages": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustOrganizationCustomPagesModel](ctx),
				Attributes: map[string]schema.Attribute{
					"forbidden": schema.StringAttribute{
						Description: "The uid of the custom page to use when a user is denied access after failing a non-identity rule.",
						Optional:    true,
					},
					"identity_denied": schema.StringAttribute{
						Description: "The uid of the custom page to use when a user is denied access.",
						Optional:    true,
					},
				},
			},
			"login_design": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustOrganizationLoginDesignModel](ctx),
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
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustOrganizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustOrganizationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
