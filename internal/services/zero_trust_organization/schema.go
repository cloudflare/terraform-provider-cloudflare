// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustOrganizationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Access: Organizations, Identity Providers, and Groups Read",
				"Access: Organizations, Identity Providers, and Groups Revoke",
				"Access: Organizations, Identity Providers, and Groups Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
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
				Optional:    true,
			},
			"deny_unmatched_requests": schema.BoolAttribute{
				Description: "Determines whether to deny all requests to Cloudflare-protected resources that lack an associated Access application. If enabled, you must explicitly configure an Access application and policy to allow traffic to your Cloudflare-protected resources. For domains you want to be public across all subdomains, add the domain to the `deny_unmatched_requests_exempted_zone_names` array.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of your Zero Trust organization.",
				Optional:    true,
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for applications will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
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
			"deny_unmatched_requests_exempted_zone_names": schema.ListAttribute{
				Description: "Contains zone names to exempt from the `deny_unmatched_requests` feature. Requests to a subdomain in an exempted zone will block unauthenticated traffic by default if there is a configured Access application and policy that matches the request.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"custom_pages": schema.SingleNestedAttribute{
				Optional: true,
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
			"mfa_config": schema.SingleNestedAttribute{
				Description: "Configures multi-factor authentication (MFA) settings for an organization.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"allowed_authenticators": schema.ListAttribute{
						Description: "Lists the MFA methods that users can authenticate with.",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"totp",
									"biometrics",
									"security_key",
									"ssh_piv_key",
								),
							),
						},
						ElementType: types.StringType,
					},
					"amr_matching_session_duration": schema.StringAttribute{
						Description: `Allows a user to skip MFA via Authentication Method Reference (AMR) matching when the AMR claim provided by the IdP the user used to authenticate contains "mfa". Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days).`,
						Optional:    true,
					},
					"required_aaguids": schema.StringAttribute{
						Description: "Specifies a Cloudflare List of required FIDO2 authenticator device AAGUIDs.",
						Optional:    true,
					},
					"session_duration": schema.StringAttribute{
						Description: "Defines the duration of an MFA session. Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.",
						Optional:    true,
					},
				},
			},
			"mfa_ssh_piv_key_requirements": schema.SingleNestedAttribute{
				Description: "Configures SSH PIV key requirements for MFA using hardware security keys.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"pin_policy": schema.StringAttribute{
						Description: "Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN required), `once` (PIN required once per session), `always` (PIN required for each use).\nAvailable values: \"never\", \"once\", \"always\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"never",
								"once",
								"always",
							),
						},
					},
					"require_fips_device": schema.BoolAttribute{
						Description: "Requires the SSH PIV key to be stored on a FIPS 140-2 Level 1 or higher validated device.",
						Optional:    true,
					},
					"ssh_key_size": schema.ListAttribute{
						Description: "Specifies the allowed SSH key sizes in bits. Valid sizes depend on key type. Ed25519 has a fixed key size and does not accept this parameter.",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueInt64sAre(
								int64validator.OneOf(
									256,
									384,
									521,
									2048,
									3072,
									4096,
								),
							),
						},
						ElementType: types.Int64Type,
					},
					"ssh_key_type": schema.ListAttribute{
						Description: "Specifies the allowed SSH key types. Valid values are `ecdsa`, `ed25519`, and `rsa`.",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"ecdsa",
									"ed25519",
									"rsa",
								),
							),
						},
						ElementType: types.StringType,
					},
					"touch_policy": schema.StringAttribute{
						Description: "Defines when physical touch is required to use the SSH key. Valid values: `never` (no touch required), `always` (touch required for each use), `cached` (touch cached for 15 seconds).\nAvailable values: \"never\", \"always\", \"cached\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"never",
								"always",
								"cached",
							),
						},
					},
				},
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_ui_read_only": schema.BoolAttribute{
				Description: "Lock all settings as Read-Only in the Dashboard, regardless of user permission. Updates may only be made via the API or Terraform for this account when enabled.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"mfa_required_for_all_apps": schema.BoolAttribute{
				Description: "Determines whether global MFA settings apply to applications by default. The organization must have MFA enabled with at least one authentication method and a session duration configured.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"ui_read_only_toggle_reason": schema.StringAttribute{
				Description: "A description of the reason why the UI read only field is being toggled.",
				Computed:    true,
				Optional:    true,
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
