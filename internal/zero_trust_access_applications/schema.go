// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_applications

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r ZeroTrustAccessApplicationsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"app_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"domain": schema.StringAttribute{
				Description: "The primary hostname and path that Access will secure. If the app is visible in the App Launcher dashboard, this is the domain that will be displayed.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The application type.",
				Optional:    true,
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate to this application using their WARP session.  When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.",
				Optional:    true,
			},
			"allowed_idps": schema.StringAttribute{
				Description: "The identity providers your users can select when connecting to this application. Defaults to all IdPs configured in your account.",
				Optional:    true,
			},
			"app_launcher_visible": schema.BoolAttribute{
				Description: "Displays the application in the App Launcher.",
				Optional:    true,
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login. You must specify only one identity provider in allowed_idps.",
				Optional:    true,
			},
			"cors_headers": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"allow_all_headers": schema.BoolAttribute{
						Description: "Allows all HTTP request headers.",
						Optional:    true,
					},
					"allow_all_methods": schema.BoolAttribute{
						Description: "Allows all HTTP request methods.",
						Optional:    true,
					},
					"allow_all_origins": schema.BoolAttribute{
						Description: "Allows all origins.",
						Optional:    true,
					},
					"allow_credentials": schema.BoolAttribute{
						Description: "When set to `true`, includes credentials (cookies, authorization headers, or TLS client certificates) with requests.",
						Optional:    true,
					},
					"allowed_headers": schema.StringAttribute{
						Description: "Allowed HTTP request headers.",
						Optional:    true,
					},
					"allowed_methods": schema.StringAttribute{
						Description: "Allowed HTTP request methods.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("GET", "POST", "HEAD", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
						},
					},
					"allowed_origins": schema.StringAttribute{
						Description: "Allowed origins.",
						Optional:    true,
					},
					"max_age": schema.Float64Attribute{
						Description: "The maximum number of seconds the results of a preflight request can be cached.",
						Optional:    true,
					},
				},
			},
			"custom_deny_message": schema.StringAttribute{
				Description: "The custom error message shown to a user when they are denied access to the application.",
				Optional:    true,
			},
			"custom_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing identity-based rules.",
				Optional:    true,
			},
			"custom_non_identity_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing non-identity rules.",
				Optional:    true,
			},
			"custom_pages": schema.StringAttribute{
				Description: "The custom pages that will be displayed when applicable for this application",
				Optional:    true,
			},
			"enable_binding_cookie": schema.BoolAttribute{
				Description: "Enables the binding cookie, which increases security against compromised authorization tokens and CSRF attacks.",
				Optional:    true,
			},
			"http_only_cookie_attribute": schema.BoolAttribute{
				Description: "Enables the HttpOnly cookie attribute, which increases security against XSS attacks.",
				Optional:    true,
			},
			"logo_url": schema.StringAttribute{
				Description: "The image URL for the logo shown in the App Launcher dashboard.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the application.",
				Optional:    true,
			},
			"path_cookie_attribute": schema.BoolAttribute{
				Description: "Enables cookie paths to scope an application's JWT to the application path. If disabled, the JWT will scope to the hostname by default",
				Optional:    true,
			},
			"same_site_cookie_attribute": schema.StringAttribute{
				Description: "Sets the SameSite cookie setting, which provides increased security against CSRF attacks.",
				Optional:    true,
			},
			"self_hosted_domains": schema.StringAttribute{
				Description: "List of domains that Access will secure.",
				Optional:    true,
			},
			"service_auth_401_redirect": schema.BoolAttribute{
				Description: "Returns a 401 status code when the request is blocked by a Service Auth policy.",
				Optional:    true,
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for this application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
				Optional:    true,
			},
			"skip_interstitial": schema.BoolAttribute{
				Description: "Enables automatic authentication through cloudflared.",
				Optional:    true,
			},
			"tags": schema.StringAttribute{
				Description: "The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.",
				Optional:    true,
			},
			"saas_app": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Description: "Optional identifier indicating the authentication protocol used for the saas app. Required for OIDC. Default if unset is \"saml\"",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("saml", "oidc"),
						},
					},
					"consumer_service_url": schema.StringAttribute{
						Description: "The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.",
						Optional:    true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"custom_attributes": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description: "The name of the attribute.",
								Optional:    true,
							},
							"name_format": schema.StringAttribute{
								Description: "A globally unique name for an identity or service provider.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified", "urn:oasis:names:tc:SAML:2.0:attrname-format:basic", "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"),
								},
							},
							"source": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "The name of the IdP attribute.",
										Optional:    true,
									},
								},
							},
						},
					},
					"default_relay_state": schema.StringAttribute{
						Description: "The URL that the user will be redirected to after a successful login for IDP initiated logins.",
						Optional:    true,
					},
					"idp_entity_id": schema.StringAttribute{
						Description: "The unique identifier for your SaaS application.",
						Optional:    true,
					},
					"name_id_format": schema.StringAttribute{
						Description: "The format of the name identifier sent to the SaaS application.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "email"),
						},
					},
					"name_id_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.\n",
						Optional:    true,
					},
					"public_key": schema.StringAttribute{
						Description: "The Access public certificate that will be used to verify your identity.",
						Optional:    true,
					},
					"saml_attribute_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata] (https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate.  The output of this expression must be a JSON object.\n",
						Optional:    true,
					},
					"sp_entity_id": schema.StringAttribute{
						Description: "A globally unique name for an identity or service provider.",
						Optional:    true,
					},
					"sso_endpoint": schema.StringAttribute{
						Description: "The endpoint where your SaaS application will send login requests.",
						Optional:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
					"app_launcher_url": schema.StringAttribute{
						Description: "The URL where this applications tile redirects users",
						Optional:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "The application client id",
						Optional:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "The application client secret, only returned on POST request.",
						Optional:    true,
					},
					"grant_types": schema.ListAttribute{
						Description: "The OIDC flows supported by this application",
						Optional:    true,
						ElementType: types.StringType,
					},
					"group_filter_regex": schema.StringAttribute{
						Description: "A regex to filter Cloudflare groups returned in ID token and userinfo endpoint",
						Optional:    true,
					},
					"redirect_uris": schema.ListAttribute{
						Description: "The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens",
						Optional:    true,
						ElementType: types.StringType,
					},
					"scopes": schema.ListAttribute{
						Description: "Define the user information shared with access",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"id": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
			},
			"aud": schema.StringAttribute{
				Description: "Audience tag.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
