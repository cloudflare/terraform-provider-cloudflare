// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessIdentityProviderResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"name": schema.StringAttribute{
				Description: "The name of the identity provider, shown to users on the login page.",
				Required:    true,
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
						Sensitive:   true,
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
						Description: "Indicates the type of user interaction that is required. prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.\nAvailable values: \"login\", \"select_account\", \"none\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"login",
								"select_account",
								"none",
							),
						},
					},
					"support_groups": schema.BoolAttribute{
						Description: "Should Cloudflare try to load groups from your account",
						Optional:    true,
					},
					"centrify_account": schema.StringAttribute{
						Description: "Your centrify account url",
						Optional:    true,
					},
					"centrify_app_id": schema.StringAttribute{
						Description: "Your centrify app id",
						Optional:    true,
					},
					"apps_domain": schema.StringAttribute{
						Description: "Your companies TLD",
						Optional:    true,
					},
					"auth_url": schema.StringAttribute{
						Description: "The authorization_endpoint URL of your IdP",
						Optional:    true,
					},
					"certs_url": schema.StringAttribute{
						Description: "The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens",
						Optional:    true,
					},
					"pkce_enabled": schema.BoolAttribute{
						Description: "Enable Proof Key for Code Exchange (PKCE)",
						Optional:    true,
					},
					"scopes": schema.ListAttribute{
						Description: "OAuth scopes",
						Optional:    true,
						ElementType: types.StringType,
					},
					"token_url": schema.StringAttribute{
						Description: "The token_endpoint URL of your IdP",
						Optional:    true,
					},
					"authorization_server_id": schema.StringAttribute{
						Description: "Your okta authorization server id",
						Optional:    true,
					},
					"okta_account": schema.StringAttribute{
						Description: "Your okta account url",
						Optional:    true,
					},
					"onelogin_account": schema.StringAttribute{
						Description: "Your OneLogin account url",
						Optional:    true,
					},
					"ping_env_id": schema.StringAttribute{
						Description: "Your PingOne environment identifier",
						Optional:    true,
					},
					"attributes": schema.ListAttribute{
						Description: "A list of SAML attribute names that will be added to your signed JWT token and can be used in SAML policy rules.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"email_attribute_name": schema.StringAttribute{
						Description: "The attribute name for email in the SAML response.",
						Optional:    true,
					},
					"header_attributes": schema.ListNestedAttribute{
						Description: "Add a list of attribute names that will be returned in the response header from the Access callback.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "attribute name from the IDP",
									Optional:    true,
								},
								"header_name": schema.StringAttribute{
									Description: "header that will be added on the request to the origin",
									Optional:    true,
								},
							},
						},
					},
					"idp_public_certs": schema.ListAttribute{
						Description: "X509 certificate to verify the signature in the SAML authentication response",
						Optional:    true,
						ElementType: types.StringType,
					},
					"issuer_url": schema.StringAttribute{
						Description: "IdP Entity ID or Issuer URL",
						Optional:    true,
					},
					"sign_request": schema.BoolAttribute{
						Description: "Sign the SAML authentication request with Access credentials. To verify the signature, use the public key from the Access certs endpoints.",
						Optional:    true,
					},
					"sso_target_url": schema.StringAttribute{
						Description: "URL to send the SAML authentication requests to",
						Optional:    true,
					},
					"redirect_url": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"type": schema.StringAttribute{
				Required:   true,
				CustomType: jsontypes.NormalizedType{},
			},
			"scim_config": schema.StringAttribute{
				Optional:   true,
				CustomType: jsontypes.NormalizedType{},
			},
		},
	}
}

func (r *ZeroTrustAccessIdentityProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessIdentityProviderResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
