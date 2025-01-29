// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessIdentityProviderDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
			},
			"identity_provider_id": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the identity provider, shown to users on the login page.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of identity provider. To determine the value for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"onetimepin",
						"azureAD",
						"saml",
						"centrify",
						"facebook",
						"github",
						"google-apps",
						"google",
						"linkedin",
						"oidc",
						"okta",
						"onelogin",
						"pingone",
						"yandex",
					),
				},
			},
			"config": schema.SingleNestedAttribute{
				Description: "The configuration parameters for the identity provider. To view the required parameters for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessIdentityProviderConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"claims": schema.ListAttribute{
						Description: "Custom claims",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"client_id": schema.StringAttribute{
						Description: "Your OAuth Client ID",
						Computed:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "Your OAuth Client Secret",
						Computed:    true,
					},
					"conditional_access_enabled": schema.BoolAttribute{
						Description: "Should Cloudflare try to load authentication contexts from your account",
						Computed:    true,
					},
					"directory_id": schema.StringAttribute{
						Description: "Your Azure directory uuid",
						Computed:    true,
					},
					"email_claim_name": schema.StringAttribute{
						Description: "The claim name for email in the id_token response.",
						Computed:    true,
					},
					"prompt": schema.StringAttribute{
						Description: "Indicates the type of user interaction that is required. prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.",
						Computed:    true,
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
						Computed:    true,
					},
					"centrify_account": schema.StringAttribute{
						Description: "Your centrify account url",
						Computed:    true,
					},
					"centrify_app_id": schema.StringAttribute{
						Description: "Your centrify app id",
						Computed:    true,
					},
					"apps_domain": schema.StringAttribute{
						Description: "Your companies TLD",
						Computed:    true,
					},
					"auth_url": schema.StringAttribute{
						Description: "The authorization_endpoint URL of your IdP",
						Computed:    true,
					},
					"certs_url": schema.StringAttribute{
						Description: "The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens",
						Computed:    true,
					},
					"pkce_enabled": schema.BoolAttribute{
						Description: "Enable Proof Key for Code Exchange (PKCE)",
						Computed:    true,
					},
					"scopes": schema.ListAttribute{
						Description: "OAuth scopes",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"token_url": schema.StringAttribute{
						Description: "The token_endpoint URL of your IdP",
						Computed:    true,
					},
					"authorization_server_id": schema.StringAttribute{
						Description: "Your okta authorization server id",
						Computed:    true,
					},
					"okta_account": schema.StringAttribute{
						Description: "Your okta account url",
						Computed:    true,
					},
					"onelogin_account": schema.StringAttribute{
						Description: "Your OneLogin account url",
						Computed:    true,
					},
					"ping_env_id": schema.StringAttribute{
						Description: "Your PingOne environment identifier",
						Computed:    true,
					},
					"attributes": schema.ListAttribute{
						Description: "A list of SAML attribute names that will be added to your signed JWT token and can be used in SAML policy rules.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"email_attribute_name": schema.StringAttribute{
						Description: "The attribute name for email in the SAML response.",
						Computed:    true,
					},
					"header_attributes": schema.ListNestedAttribute{
						Description: "Add a list of attribute names that will be returned in the response header from the Access callback.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessIdentityProviderConfigHeaderAttributesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "attribute name from the IDP",
									Computed:    true,
								},
								"header_name": schema.StringAttribute{
									Description: "header that will be added on the request to the origin",
									Computed:    true,
								},
							},
						},
					},
					"idp_public_certs": schema.ListAttribute{
						Description: "X509 certificate to verify the signature in the SAML authentication response",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"issuer_url": schema.StringAttribute{
						Description: "IdP Entity ID or Issuer URL",
						Computed:    true,
					},
					"sign_request": schema.BoolAttribute{
						Description: "Sign the SAML authentication request with Access credentials. To verify the signature, use the public key from the Access certs endpoints.",
						Computed:    true,
					},
					"sso_target_url": schema.StringAttribute{
						Description: "URL to send the SAML authentication requests to",
						Computed:    true,
					},
					"redirect_url": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"scim_config": schema.SingleNestedAttribute{
				Description: "The configuration settings for enabling a System for Cross-Domain Identity Management (SCIM) with the identity provider.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessIdentityProviderSCIMConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "A flag to enable or disable SCIM for the identity provider.",
						Computed:    true,
					},
					"identity_update_behavior": schema.StringAttribute{
						Description: "Indicates how a SCIM event updates a user identity used for policy evaluation. Use \"automatic\" to automatically update a user's identity and augment it with fields from the SCIM user resource. Use \"reauth\" to force re-authentication on group membership updates, user identity update will only occur after successful re-authentication. With \"reauth\" identities will not contain fields from the SCIM user resource. With \"no_action\" identities will not be changed by SCIM updates in any way and users will not be prompted to reauthenticate.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"automatic",
								"reauth",
								"no_action",
							),
						},
					},
					"scim_base_url": schema.StringAttribute{
						Description: "The base URL of Cloudflare's SCIM V2.0 API endpoint.",
						Computed:    true,
					},
					"seat_deprovision": schema.BoolAttribute{
						Description: "A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.  This cannot be enabled unless user_deprovision is also enabled.",
						Computed:    true,
					},
					"secret": schema.StringAttribute{
						Description: "A read-only token generated when the SCIM integration is enabled for the first time.  It is redacted on subsequent requests.  If you lose this you will need to refresh it at /access/identity_providers/:idpID/refresh_scim_secret.",
						Computed:    true,
					},
					"user_deprovision": schema.BoolAttribute{
						Description: "A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.",
						Computed:    true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"scim_enabled": schema.StringAttribute{
						Description: "Indicates to Access to only retrieve identity providers that have the System for Cross-Domain Identity Management (SCIM) enabled.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessIdentityProviderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessIdentityProviderDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("identity_provider_id"), path.MatchRoot("filter")),
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
