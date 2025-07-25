package access_identity_provider

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ResourceSchema returns the Terraform schema for the deprecated cloudflare_access_identity_provider resource.
func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID.",
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
			"type": schema.StringAttribute{
				Description:   "The type of identity provider. To determine the value for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).\nAvailable values: \"onetimepin\", \"azureAD\", \"saml\", \"centrify\", \"facebook\", \"github\", \"google-apps\", \"google\", \"linkedin\", \"oidc\", \"okta\", \"onelogin\", \"pingone\", \"yandex\".",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
						WriteOnly:   true,
					},
					"conditional_access_enabled": schema.BoolAttribute{
						Description: "Should Cloudflare try to load authentication contexts from your account",
						Optional:    true,
						Validators: []validator.Bool{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "azureAD"),
						},
					},
					"directory_id": schema.StringAttribute{
						Description: "Your Azure directory uuid",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "azureAD"),
						},
					},
					"email_claim_name": schema.StringAttribute{
						Description: "The claim name for email in the id_token response.",
						Optional:    true,
					},
					"prompt": schema.StringAttribute{
						Description: "The prompt parameter for the OAuth provider's authorization endpoint.",
						Optional:    true,
					},
					"support_groups": schema.BoolAttribute{
						Description: "Should Cloudflare try to load groups from your account",
						Optional:    true,
						Validators: []validator.Bool{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "azureAD"),
						},
					},
					"centrify_account": schema.StringAttribute{
						Description: "Your Centrify account domain",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "centrify"),
						},
					},
					"centrify_app_id": schema.StringAttribute{
						Description: "Your Centrify app id",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "centrify"),
						},
					},
					"apps_domain": schema.StringAttribute{
						Description: "Your Google Apps domain",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "google-apps"),
						},
					},
					"auth_url": schema.StringAttribute{
						Description: "Your OAuth authorization URL",
						Optional:    true,
					},
					"certs_url": schema.StringAttribute{
						Description: "Your OAuth certificates URL",
						Optional:    true,
					},
					"pkce_enabled": schema.BoolAttribute{
						Description: "Enable PKCE",
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"scopes": schema.ListAttribute{
						Description: "The scopes for your OAuth provider",
						Optional:    true,
						ElementType: types.StringType,
					},
					"token_url": schema.StringAttribute{
						Description: "Your OAuth token URL",
						Optional:    true,
					},
					"authorization_server_id": schema.StringAttribute{
						Description: "Your Okta authorization server id",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "okta"),
						},
					},
					"okta_account": schema.StringAttribute{
						Description: "Your Okta account domain",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "okta"),
						},
					},
					"onelogin_account": schema.StringAttribute{
						Description: "Your OneLogin account domain",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "onelogin"),
						},
					},
					"ping_env_id": schema.StringAttribute{
						Description: "Your PingOne environment identifier",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "pingone"),
						},
					},
					"attributes": schema.ListAttribute{
						Description: "Custom attributes",
						Optional:    true,
						ElementType: types.StringType,
					},
					"email_attribute_name": schema.StringAttribute{
						Description: "The attribute name for email in the SAML response",
						Optional:    true,
					},
					"header_attributes": schema.ListNestedAttribute{
						Description: "Add your custom headers when you make requests to the authorization URL",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the attribute",
									Optional:    true,
								},
								"header_name": schema.StringAttribute{
									Description: "The name of the header",
									Optional:    true,
								},
							},
						},
					},
					"idp_public_certs": schema.ListAttribute{
						Description: "Your IdP's public certificates",
						Optional:    true,
						ElementType: types.StringType,
					},
					"issuer_url": schema.StringAttribute{
						Description: "Your IdP's issuer URL",
						Optional:    true,
					},
					"sign_request": schema.BoolAttribute{
						Description: "Should Cloudflare sign the request",
						Optional:    true,
					},
					"sso_target_url": schema.StringAttribute{
						Description: "Your IdP's SSO target URL",
						Optional:    true,
					},
					"redirect_url": schema.StringAttribute{
						Description: "The callback URL for your OAuth provider",
						Computed:    true,
					},
				},
			},
			"scim_config": schema.SingleNestedAttribute{
				Description: "The SCIM configuration for the identity provider.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether SCIM is enabled for this identity provider.",
						Optional:    true,
						Computed:    true,
					},
					"identity_update_behavior": schema.StringAttribute{
						Description: "The behavior for identity updates.",
						Optional:    true,
						Computed:    true,
					},
					"scim_base_url": schema.StringAttribute{
						Description: "The SCIM base URL.",
						Computed:    true,
					},
					"seat_deprovision": schema.BoolAttribute{
						Description: "Whether to deprovision seats when users are removed.",
						Optional:    true,
						Computed:    true,
					},
					"secret": schema.StringAttribute{
						Description: "The SCIM secret.",
						Computed:    true,
						Sensitive:   true,
					},
					"user_deprovision": schema.BoolAttribute{
						Description: "Whether to deprovision users when they are removed.",
						Optional:    true,
						Computed:    true,
					},
				},
			},
		},
	}
}

// Duplicate Schema and ConfigValidators implementations were removed to avoid
// redeclaration errors. The implementations reside in `resource.go`.
