// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_identity_provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r ZeroTrustIdentityProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "The configuration parameters for the identity provider. To view the required parameters for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"claims": schema.ListAttribute{
						Description: "Custom claims",
						Optional:    true,
						ElementType: types.StringType,
					},
					"client_id": schema.StringAttribute{
						Description: "Your OAuth Client ID",
						Optional:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "Your OAuth Client Secret",
						Optional:    true,
					},
					"conditional_access_enabled": schema.BoolAttribute{
						Description: "Should Cloudflare try to load authentication contexts from your account",
						Optional:    true,
					},
					"directory_id": schema.StringAttribute{
						Description: "Your Azure directory uuid",
						Optional:    true,
					},
					"email_claim_name": schema.StringAttribute{
						Description: "The claim name for email in the id_token response.",
						Optional:    true,
					},
					"prompt": schema.StringAttribute{
						Description: "Indicates the type of user interaction that is required. prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("login", "select_account", "none"),
						},
					},
					"support_groups": schema.BoolAttribute{
						Description: "Should Cloudflare try to load groups from your account",
						Optional:    true,
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the identity provider, shown to users on the login page.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of identity provider. To determine the value for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("onetimepin", "azureAD", "saml", "centrify", "facebook", "github", "google-apps", "google", "linkedin", "oidc", "okta", "onelogin", "pingone", "yandex"),
				},
			},
			"scim_config": schema.SingleNestedAttribute{
				Description: "The configuration settings for enabling a System for Cross-Domain Identity Management (SCIM) with the identity provider.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "A flag to enable or disable SCIM for the identity provider.",
						Optional:    true,
					},
					"group_member_deprovision": schema.BoolAttribute{
						Description: "A flag to revoke a user's session in Access and force a reauthentication on the user's Gateway session when they have been added or removed from a group in the Identity Provider.",
						Optional:    true,
					},
					"seat_deprovision": schema.BoolAttribute{
						Description: "A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.  This cannot be enabled unless user_deprovision is also enabled.",
						Optional:    true,
					},
					"secret": schema.StringAttribute{
						Description: "A read-only token generated when the SCIM integration is enabled for the first time.  It is redacted on subsequent requests.  If you lose this you will need to refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.",
						Optional:    true,
					},
					"user_deprovision": schema.BoolAttribute{
						Description: "A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.",
						Optional:    true,
					},
				},
			},
		},
	}
}
