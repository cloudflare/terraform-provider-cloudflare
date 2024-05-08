package sdkv2provider

import (
	"fmt"
	"regexp"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudflareAccessApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   consts.AccountIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   consts.ZoneIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"aud": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Application Audience (AUD) Tag of the application.",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Friendly name of the Access Application.",
			Optional:    true,
		},
		"domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The primary hostname and path that Access will secure. If the app is visible in the App Launcher dashboard, this is the domain that will be displayed.",
		},
		"self_hosted_domains": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of domains that access will secure. Only present for self_hosted, vnc, and ssh applications. Always includes the value set as `domain`",
		},
		"type": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "self_hosted",
			ValidateFunc: validation.StringInSlice([]string{"app_launcher", "bookmark", "biso", "dash_sso", "saas", "self_hosted", "ssh", "vnc", "warp"}, false),
			Description:  fmt.Sprintf("The application type. %s", renderAvailableDocumentationValuesStringSlice([]string{"app_launcher", "bookmark", "biso", "dash_sso", "saas", "self_hosted", "ssh", "vnc", "warp"})),
		},
		"policies": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			Description: "The policies associated with the application, in ascending order of precedence." +
				" When omitted, the application policies are not be updated." +
				" Warning: Do not use this field while you still have this application ID referenced as `application_id`" +
				" in an `cloudflare_access_policy` resource, as it can result in an inconsistent state.",
		},
		"session_duration": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "24h",
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				appType := d.Get("type").(string)
				// Suppress the diff if it's a bookmark app type. Bookmarks don't have a session duration
				// field which always creates a diff because of the default '24h' value.
				if appType == "bookmark" {
					return true
				}

				return oldValue == newValue
			},
			ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
				v := val.(string)
				_, err := time.ParseDuration(v)
				if err != nil {
					errs = append(errs, fmt.Errorf(`%q only supports "ns", "us" (or "µs"), "ms", "s", "m", or "h" as valid units`, key))
				}
				return
			},
			Description: "How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`",
		},
		"cors_headers": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "CORS configuration for the Access Application. See below for reference structure.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"allowed_methods": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of methods to expose via CORS.",
					},
					"allowed_origins": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of origins permitted to make CORS requests.",
					},
					"allowed_headers": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of HTTP headers to expose via CORS.",
					},
					"allow_all_methods": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Value to determine whether all methods are exposed.",
					},
					"allow_all_origins": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Value to determine whether all origins are permitted to make CORS requests.",
					},
					"allow_all_headers": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Value to determine whether all HTTP headers are exposed.",
					},
					"allow_credentials": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Value to determine if credentials (cookies, authorization headers, or TLS client certificates) are included with requests.",
					},
					"max_age": {
						Type:         schema.TypeInt,
						Optional:     true,
						ValidateFunc: validation.IntBetween(-1, 86400),
						Description:  "The maximum time a preflight request will be cached.",
					},
				},
			},
		},
		"saas_app": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "SaaS configuration for the Access Application.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					// shared values
					"auth_type": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"oidc", "saml"}, false),
						Description:  "",
					},
					"public_key": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The public certificate that will be used to verify identities.",
					},

					// OIDC options
					"client_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The application client id",
					},
					"client_secret": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The application client secret, only returned on initial apply",
						Sensitive:   true,
					},
					"redirect_uris": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens",
					},
					"grant_types": {
						Type:     schema.TypeSet,
						Optional: true,
						Computed: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "The OIDC flows supported by this application",
					},
					"scopes": {
						Type:     schema.TypeSet,
						Optional: true,
						Computed: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "Define the user information shared with access",
					},
					"app_launcher_url": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The URL where this applications tile redirects users",
					},
					"group_filter_regex": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A regex to filter Cloudflare groups returned in ID token and userinfo endpoint",
					},

					// SAML options
					"sp_entity_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A globally unique name for an identity or service provider.",
					},
					"consumer_service_url": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.",
					},
					"name_id_format": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"email", "id"}, false),
						Description:  "The format of the name identifier sent to the SaaS application.",
					},
					"custom_attribute": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Custom attribute mapped from IDPs.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The name of the attribute as provided to the SaaS app.",
								},
								"name_format": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "A globally unique name for an identity or service provider.",
								},
								"friendly_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "A friendly name for the attribute as provided to the SaaS app.",
								},
								"required": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "True if the attribute must be always present.",
								},
								"source": {
									Type:     schema.TypeList,
									Required: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Required:    true,
												Description: "The name of the attribute as provided by the IDP.",
											},
										},
									},
								},
							},
						},
					},
					"idp_entity_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The unique identifier for the SaaS application.",
					},
					"sso_endpoint": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The endpoint where the SaaS application will send login requests.",
					},
					"default_relay_state": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The relay state used if not provided by the identity provider.",
					},
					"name_id_transform_jsonata": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.",
					},
					"saml_attribute_transform_jsonata": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.",
					},
				},
			},
		},
		"auto_redirect_to_identity": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Option to skip identity provider selection if only one is configured in `allowed_idps`.",
		},
		"enable_binding_cookie": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional \"binding\" cookie on requests.",
		},
		"allowed_idps": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The identity providers selected for the application.",
		},
		"custom_deny_message": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Option that returns a custom error message when a user is denied access to the application.",
		},
		"custom_deny_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Option that redirects to a custom URL when a user is denied access to the application via identity based rules.",
		},
		"custom_non_identity_deny_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Option that redirects to a custom URL when a user is denied access to the application via non identity rules.",
		},
		"http_only_cookie_attribute": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Option to add the `HttpOnly` cookie flag to access tokens.",
		},
		"same_site_cookie_attribute": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"none", "lax", "strict"}, false),
			Description:  fmt.Sprintf("Defines the same-site cookie setting for access tokens. %s", renderAvailableDocumentationValuesStringSlice(([]string{"none", "lax", "strict"}))),
		},
		"logo_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Image URL for the logo shown in the app launcher dashboard.",
		},
		"skip_interstitial": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Option to skip the authorization interstitial when using the CLI.",
		},
		"app_launcher_visible": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Option to show/hide applications in App Launcher.",
		},
		"service_auth_401_redirect": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Option to return a 401 status code in service authentication rules on failed requests.",
		},
		"custom_pages": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The custom pages selected for the application.",
		},
		"tags": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The itags associated with the application.",
		},
		"app_launcher_logo_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The logo URL of the app launcher.",
		},
		"header_bg_color": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The background color of the header bar in the app launcher.",
		},
		"bg_color": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The background color of the app launcher.",
		},
		"footer_links": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The name of the footer link.",
					},
					"url": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The URL of the footer link.",
					},
				},
			},
			Description: "The footer links of the app launcher.",
		},
		"landing_page_design": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The landing page design of the app launcher.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"title": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The title of the landing page.",
					},
					"message": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The message of the landing page.",
					},
					"button_text_color": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The button text color of the landing page.",
					},
					"button_color": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The button color of the landing page.",
					},
					"image_url": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The URL of the image to be displayed in the landing page.",
					},
				},
			},
		},
		"allow_authenticate_via_warp": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When set to true, users can authenticate to this application using their WARP session. When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.",
		},
		"options_preflight_bypass": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Allows options preflight requests to bypass Access authentication and go directly to the origin. Cannot turn on if cors_headers is set.",
		},
		"scim_config": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Configuration for provisioning to this application via SCIM. This is currently in closed beta.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether SCIM provisioning is turned on for this application.",
					},
					"remote_uri": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The base URI for the application's SCIM-compatible API.",
					},
					"idp_uid": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The UID of the IdP to use as the source for SCIM resources to provision to this application.",
					},
					"deactivate_on_delete": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "If false, propagates DELETE requests to the target application for SCIM resources. If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.",
					},
					"authentication": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Attributes for configuring HTTP Basic, OAuth Bearer token, or OAuth 2 authentication schemes for SCIM provisioning to an application.",
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								// Common Attributes
								"scheme": {
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"httpbasic", "oauthbearertoken", "oauth2"}, false),
									Description:  "The authentication scheme to use when making SCIM requests to this application.",
								},
								// HTTP Basic Authentication Attributes
								"user": {
									Type:          schema.TypeString,
									Optional:      true,
									RequiredWith:  []string{"scim_config.0.authentication.0.password"},
									ConflictsWith: []string{"scim_config.0.authentication.0.token", "scim_config.0.authentication.0.client_id", "scim_config.0.authentication.0.client_secret", "scim_config.0.authentication.0.authorization_url", "scim_config.0.authentication.0.token_url", "scim_config.0.authentication.0.scopes"},
									Description:   "User name used to authenticate with the remote SCIM service.",
								},
								"password": {
									Type:          schema.TypeString,
									Optional:      true,
									RequiredWith:  []string{"scim_config.0.authentication.0.user"},
									ConflictsWith: []string{"scim_config.0.authentication.0.token", "scim_config.0.authentication.0.client_id", "scim_config.0.authentication.0.client_secret", "scim_config.0.authentication.0.authorization_url", "scim_config.0.authentication.0.token_url", "scim_config.0.authentication.0.scopes"},
									StateFunc: func(val interface{}) string {
										return CONCEALED_STRING
									},
								},
								// OAuth Bearer Token Authentication Attributes
								"token": {
									Type:          schema.TypeString,
									Optional:      true,
									Description:   "Token used to authenticate with the remote SCIM service.",
									ConflictsWith: []string{"scim_config.0.authentication.0.user", "scim_config.0.authentication.0.password", "scim_config.0.authentication.0.client_id", "scim_config.0.authentication.0.client_secret", "scim_config.0.authentication.0.authorization_url", "scim_config.0.authentication.0.token_url", "scim_config.0.authentication.0.scopes"},
									StateFunc: func(val interface{}) string {
										return CONCEALED_STRING
									},
								},
								// OAuth 2 Authentication Attributes
								"client_id": {
									Type:          schema.TypeString,
									Optional:      true,
									Description:   "Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.",
									RequiredWith:  []string{"scim_config.0.authentication.0.client_secret", "scim_config.0.authentication.0.authorization_url", "scim_config.0.authentication.0.token_url"},
									ConflictsWith: []string{"scim_config.0.authentication.0.user", "scim_config.0.authentication.0.password", "scim_config.0.authentication.0.token"},
								},
								"client_secret": {
									Type:          schema.TypeString,
									Optional:      true,
									Description:   "Secret used to authenticate when generating a token for authenticating with the remove SCIM service.",
									RequiredWith:  []string{"scim_config.0.authentication.0.client_id", "scim_config.0.authentication.0.authorization_url", "scim_config.0.authentication.0.token_url"},
									ConflictsWith: []string{"scim_config.0.authentication.0.user", "scim_config.0.authentication.0.password", "scim_config.0.authentication.0.token"},
									StateFunc: func(val interface{}) string {
										return CONCEALED_STRING
									},
								},
								"authorization_url": {
									Type:          schema.TypeString,
									Optional:      true,
									Description:   "URL used to generate the auth code used during token generation.",
									RequiredWith:  []string{"scim_config.0.authentication.0.client_secret", "scim_config.0.authentication.0.client_id", "scim_config.0.authentication.0.token_url"},
									ConflictsWith: []string{"scim_config.0.authentication.0.user", "scim_config.0.authentication.0.password", "scim_config.0.authentication.0.token"},
								},
								"token_url": {
									Type:          schema.TypeString,
									Optional:      true,
									Description:   "URL used to generate the token used to authenticate with the remote SCIM service.",
									RequiredWith:  []string{"scim_config.0.authentication.0.client_secret", "scim_config.0.authentication.0.authorization_url", "scim_config.0.authentication.0.client_id"},
									ConflictsWith: []string{"scim_config.0.authentication.0.user", "scim_config.0.authentication.0.password", "scim_config.0.authentication.0.token"},
								},
								"scopes": {
									Type:          schema.TypeSet,
									Description:   "The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.",
									Optional:      true,
									ConflictsWith: []string{"scim_config.0.authentication.0.user", "scim_config.0.authentication.0.password", "scim_config.0.authentication.0.token"},
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
							},
						},
					},
					"mappings": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "A list of mappings to apply to SCIM resources before provisioning them in this application. These can transform or filter the resources to be provisioned.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"schema": {
									Type:         schema.TypeString,
									Required:     true,
									Description:  "Which SCIM resource type this mapping applies to.",
									ValidateFunc: validation.StringMatch(regexp.MustCompile(`urn:.*`), "schema must begin with \"urn:\""),
								},
								"enabled": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Whether or not this mapping is enabled.",
								},
								"filter": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.",
								},
								"transform_jsonata": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.",
								},
								"operations": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Whether or not this mapping applies to creates, updates, or deletes.",
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"create": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Whether or not this mapping applies to create (POST) operations.",
											},
											"update": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Whether or not this mapping applies to update (PATCH/PUT) operations.",
											},
											"delete": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Whether or not this mapping applies to DELETE operations.",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func convertCORSSchemaToStruct(d *schema.ResourceData) (*cloudflare.AccessApplicationCorsHeaders, error) {
	CORSConfig := cloudflare.AccessApplicationCorsHeaders{}

	if _, ok := d.GetOk("cors_headers"); ok {
		if allowedMethods, ok := d.GetOk("cors_headers.0.allowed_methods"); ok {
			CORSConfig.AllowedMethods = expandInterfaceToStringList(allowedMethods.(*schema.Set).List())
		}

		if allowedHeaders, ok := d.GetOk("cors_headers.0.allowed_headers"); ok {
			CORSConfig.AllowedHeaders = expandInterfaceToStringList(allowedHeaders.(*schema.Set).List())
		}

		if allowedOrigins, ok := d.GetOk("cors_headers.0.allowed_origins"); ok {
			CORSConfig.AllowedOrigins = expandInterfaceToStringList(allowedOrigins.(*schema.Set).List())
		}

		CORSConfig.AllowAllMethods = d.Get("cors_headers.0.allow_all_methods").(bool)
		CORSConfig.AllowAllHeaders = d.Get("cors_headers.0.allow_all_headers").(bool)
		CORSConfig.AllowAllOrigins = d.Get("cors_headers.0.allow_all_origins").(bool)
		CORSConfig.AllowCredentials = d.Get("cors_headers.0.allow_credentials").(bool)
		CORSConfig.MaxAge = d.Get("cors_headers.0.max_age").(int)

		// Prevent misconfigurations of CORS when `Access-Control-Allow-Origin` is
		// a wildcard (aka all origins) and using credentials.
		// See https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS/Errors/CORSNotSupportingCredentials
		if CORSConfig.AllowCredentials {
			if contains(CORSConfig.AllowedOrigins, "*") || CORSConfig.AllowAllOrigins {
				return nil, errors.New("CORS credentials are not permitted when all origins are allowed")
			}
		}

		// Ensure that should someone forget to set allowed methods (either
		// individually or *), we throw an error to prevent getting into an
		// unrecoverable state.
		if CORSConfig.AllowAllOrigins || len(CORSConfig.AllowedOrigins) > 1 {
			if CORSConfig.AllowAllMethods == false && len(CORSConfig.AllowedMethods) == 0 {
				return nil, errors.New("must set allowed_methods or allow_all_methods")
			}
		}

		// Ensure that should someone forget to set allowed origins (either
		// individually or *), we throw an error to prevent getting into an
		// unrecoverable state.
		if CORSConfig.AllowAllMethods || len(CORSConfig.AllowedMethods) > 1 {
			if CORSConfig.AllowAllOrigins == false && len(CORSConfig.AllowedOrigins) == 0 {
				return nil, errors.New("must set allowed_origins or allow_all_origins")
			}
		}
	}

	return &CORSConfig, nil
}

func convertCORSStructToSchema(d *schema.ResourceData, headers *cloudflare.AccessApplicationCorsHeaders) []interface{} {
	if _, ok := d.GetOk("cors_headers"); !ok {
		return []interface{}{}
	}

	if headers == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"allow_all_methods": headers.AllowAllMethods,
		"allow_all_headers": headers.AllowAllHeaders,
		"allow_all_origins": headers.AllowAllOrigins,
		"allow_credentials": headers.AllowCredentials,
		"max_age":           headers.MaxAge,
	}

	m["allowed_methods"] = flattenStringList(headers.AllowedMethods)
	m["allowed_headers"] = flattenStringList(headers.AllowedHeaders)
	m["allowed_origins"] = flattenStringList(headers.AllowedOrigins)

	return []interface{}{m}
}

func convertSAMLAttributeSchemaToStruct(data map[string]interface{}) cloudflare.SAMLAttributeConfig {
	var cfg cloudflare.SAMLAttributeConfig
	cfg.Name, _ = data["name"].(string)
	cfg.NameFormat, _ = data["name_format"].(string)
	cfg.Required, _ = data["required"].(bool)
	cfg.FriendlyName, _ = data["friendly_name"].(string)
	sourcesSlice, _ := data["source"].([]interface{})
	if len(sourcesSlice) != 0 {
		sourceMap, ok := sourcesSlice[0].(map[string]interface{})
		if ok {
			cfg.Source.Name, _ = sourceMap["name"].(string)
		}
	}

	return cfg
}

func convertSaasSchemaToStruct(d *schema.ResourceData) *cloudflare.SaasApplication {
	SaasConfig := cloudflare.SaasApplication{}
	if _, ok := d.GetOk("saas_app"); ok {
		authType := "saml"
		if rawAuthType, ok := d.GetOk("saas_app.0.auth_type"); ok {
			authType = rawAuthType.(string)
		}
		SaasConfig.AuthType = authType
		if authType == "oidc" {
			SaasConfig.ClientID = d.Get("saas_app.0.client_id").(string)
			SaasConfig.AppLauncherURL = d.Get("saas_app.0.app_launcher_url").(string)
			SaasConfig.RedirectURIs = expandInterfaceToStringList(d.Get("saas_app.0.redirect_uris").(*schema.Set).List())
			SaasConfig.GrantTypes = expandInterfaceToStringList(d.Get("saas_app.0.grant_types").(*schema.Set).List())
			SaasConfig.Scopes = expandInterfaceToStringList(d.Get("saas_app.0.scopes").(*schema.Set).List())
			SaasConfig.GroupFilterRegex = d.Get("saas_app.0.group_filter_regex").(string)
		} else {
			SaasConfig.SPEntityID = d.Get("saas_app.0.sp_entity_id").(string)
			SaasConfig.ConsumerServiceUrl = d.Get("saas_app.0.consumer_service_url").(string)
			SaasConfig.NameIDFormat = d.Get("saas_app.0.name_id_format").(string)
			SaasConfig.DefaultRelayState = d.Get("saas_app.0.default_relay_state").(string)
			SaasConfig.NameIDTransformJsonata = d.Get("saas_app.0.name_id_transform_jsonata").(string)
			SaasConfig.SamlAttributeTransformJsonata = d.Get("saas_app.0.saml_attribute_transform_jsonata").(string)

			customAttributes, _ := d.Get("saas_app.0.custom_attribute").([]interface{})
			for _, customAttributes := range customAttributes {
				attributeAsMap := customAttributes.(map[string]interface{})
				SaasConfig.CustomAttributes = append(SaasConfig.CustomAttributes, convertSAMLAttributeSchemaToStruct(attributeAsMap))
			}
		}
	}
	return &SaasConfig
}

func convertLandingPageDesignSchemaToStruct(d *schema.ResourceData) *cloudflare.AccessLandingPageDesign {
	LandingPageDesign := cloudflare.AccessLandingPageDesign{}
	if _, ok := d.GetOk("landing_page_design"); ok {
		LandingPageDesign.ButtonColor = d.Get("landing_page_design.0.button_color").(string)
		LandingPageDesign.ButtonTextColor = d.Get("landing_page_design.0.button_text_color").(string)
		LandingPageDesign.Title = d.Get("landing_page_design.0.title").(string)
		LandingPageDesign.Message = d.Get("landing_page_design.0.message").(string)
		LandingPageDesign.ImageURL = d.Get("landing_page_design.0.image_url").(string)
	}
	return &LandingPageDesign
}

func convertFooterLinksSchemaToStruct(d *schema.ResourceData) []cloudflare.AccessFooterLink {
	var footerLinks []cloudflare.AccessFooterLink
	if _, ok := d.GetOk("footer_links"); ok {
		footerLinksInterface := d.Get("footer_links").(*schema.Set).List()
		for _, footerLinkInterface := range footerLinksInterface {
			footerLink := footerLinkInterface.(map[string]interface{})
			footerLinks = append(footerLinks, cloudflare.AccessFooterLink{
				Name: footerLink["name"].(string),
				URL:  footerLink["url"].(string),
			})
		}
	}
	return footerLinks
}

func convertSCIMConfigSchemaToStruct(d *schema.ResourceData) *cloudflare.AccessApplicationSCIMConfig {
	scimConfig := new(cloudflare.AccessApplicationSCIMConfig)

	if _, ok := d.GetOk("scim_config"); ok {
		scimConfig.Enabled = cloudflare.BoolPtr(d.Get("scim_config.0.enabled").(bool))
		scimConfig.RemoteURI = d.Get("scim_config.0.remote_uri").(string)
		scimConfig.IdPUID = d.Get("scim_config.0.idp_uid").(string)
		scimConfig.DeactivateOnDelete = cloudflare.BoolPtr(d.Get("scim_config.0.deactivate_on_delete").(bool))

		if _, ok := d.GetOk("scim_config.0.authentication"); ok {
			scimConfig.Authentication = convertScimConfigAuthenticationSchemaToStruct(d)
		}

		mappings := d.Get("scim_config.0.mappings").([]interface{})

		for _, mapping := range mappings {
			mappingMap := mapping.(map[string]interface{})
			scimConfig.Mappings = append(scimConfig.Mappings, convertScimConfigMappingsSchemaToStruct(mappingMap))
		}
	}

	return scimConfig
}

func convertScimConfigMappingsSchemaToStruct(mappingData map[string]interface{}) *cloudflare.AccessApplicationScimMapping {
	mapping := new(cloudflare.AccessApplicationScimMapping)

	if mappingSchema, ok := mappingData["schema"]; ok {
		mapping.Schema = mappingSchema.(string)
	}

	if enabled, ok := mappingData["enabled"]; ok {
		mapping.Enabled = cloudflare.BoolPtr(enabled.(bool))
	}

	if filter, ok := mappingData["filter"]; ok {
		mapping.Filter = filter.(string)
	}

	if transformJsonata, ok := mappingData["transform_jsonata"]; ok {
		mapping.TransformJsonata = transformJsonata.(string)
	}

	if operations, ok := mappingData["operations"]; ok {
		ops := new(cloudflare.AccessApplicationScimMappingOperations)

		operationsArr := operations.([]interface{})

		if len(operationsArr) != 0 {
			operationsData := operationsArr[0].(map[string]interface{})

			if create, ok := operationsData["create"]; ok {
				ops.Create = cloudflare.BoolPtr(create.(bool))
			}

			if update, ok := operationsData["update"]; ok {
				ops.Update = cloudflare.BoolPtr(update.(bool))
			}

			if del, ok := operationsData["delete"]; ok {
				ops.Delete = cloudflare.BoolPtr(del.(bool))
			}
		}

		mapping.Operations = ops
	}

	return mapping
}

func convertScimConfigAuthenticationSchemaToStruct(d *schema.ResourceData) *cloudflare.AccessApplicationScimAuthenticationJson {
	auth := new(cloudflare.AccessApplicationScimAuthenticationJson)

	if _, ok := d.GetOk("scim_config.0.authentication"); ok {
		scheme := cloudflare.AccessApplicationScimAuthenticationScheme(d.Get("scim_config.0.authentication.0.scheme").(string))
		switch scheme {
		case cloudflare.AccessApplicationScimAuthenticationSchemeHttpBasic:
			base := &cloudflare.AccessApplicationScimAuthenticationHttpBasic{
				User:     d.Get("scim_config.0.authentication.0.user").(string),
				Password: d.Get("scim_config.0.authentication.0.password").(string),
			}
			base.Scheme = scheme
			auth.Value = base
			break
		case cloudflare.AccessApplicationScimAuthenticationSchemeOauthBearerToken:
			base := &cloudflare.AccessApplicationScimAuthenticationOauthBearerToken{
				Token: d.Get("scim_config.0.authentication.0.token").(string),
			}
			base.Scheme = scheme
			auth.Value = base
			break
		case cloudflare.AccessApplicationScimAuthenticationSchemeOauth2:
			base := &cloudflare.AccessApplicationScimAuthenticationOauth2{
				ClientID:         d.Get("scim_config.0.authentication.0.client_id").(string),
				ClientSecret:     d.Get("scim_config.0.authentication.0.client_secret").(string),
				AuthorizationURL: d.Get("scim_config.0.authentication.0.authorization_url").(string),
				TokenURL:         d.Get("scim_config.0.authentication.0.token_url").(string),
				Scopes:           expandInterfaceToStringList(d.Get("scim_config.0.authentication.0.scopes").(*schema.Set).List()),
			}
			base.Scheme = scheme
			auth.Value = base
			break
		}
	}

	return auth
}

func convertLandingPageDesignStructToSchema(d *schema.ResourceData, design *cloudflare.AccessLandingPageDesign) []interface{} {
	if _, ok := d.GetOk("landing_page_design"); !ok {
		return []interface{}{}
	}

	if design == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"button_color":      design.ButtonColor,
		"button_text_color": design.ButtonTextColor,
		"title":             design.Title,
		"message":           design.Message,
		"image_url":         design.ImageURL,
	}

	return []interface{}{m}
}

func convertFooterLinksStructToSchema(d *schema.ResourceData, footerLinks []cloudflare.AccessFooterLink) []interface{} {
	if _, ok := d.GetOk("footer_links"); !ok {
		return []interface{}{}
	}

	if footerLinks == nil {
		return []interface{}{}
	}

	var footerLinksInterface []interface{}
	for _, footerLink := range footerLinks {
		footerLinksInterface = append(footerLinksInterface, map[string]interface{}{
			"name": footerLink.Name,
			"url":  footerLink.URL,
		})
	}

	return footerLinksInterface
}

func convertSAMLAttributeStructToSchema(attr cloudflare.SAMLAttributeConfig) map[string]interface{} {
	m := make(map[string]interface{})
	if attr.Name != "" {
		m["name"] = attr.Name
	}
	if attr.NameFormat != "" {
		m["name_format"] = attr.NameFormat
	}
	if attr.Required {
		m["required"] = true
	}
	if attr.FriendlyName != "" {
		m["friendly_name"] = attr.FriendlyName
	}
	if attr.Source.Name != "" {
		m["source"] = []interface{}{map[string]interface{}{"name": attr.Source.Name}}
	}
	return m
}

func convertSaasStructToSchema(d *schema.ResourceData, app *cloudflare.SaasApplication) []interface{} {
	if app == nil {
		return []interface{}{}
	}
	if app.AuthType == "oidc" {
		m := map[string]interface{}{
			"auth_type":          app.AuthType,
			"client_id":          app.ClientID,
			"redirect_uris":      app.RedirectURIs,
			"grant_types":        app.GrantTypes,
			"scopes":             app.Scopes,
			"public_key":         app.PublicKey,
			"group_filter_regex": app.GroupFilterRegex,
			"app_launcher_url":   app.AppLauncherURL,
		}
		// client secret is only returned on create, if it is present in the state, preserve it
		if client_secret, ok := d.GetOk("saas_app.0.client_secret"); ok {
			m["client_secret"] = client_secret.(string)
		}
		return []interface{}{m}
	} else {
		m := map[string]interface{}{
			"sp_entity_id":                     app.SPEntityID,
			"consumer_service_url":             app.ConsumerServiceUrl,
			"name_id_format":                   app.NameIDFormat,
			"idp_entity_id":                    app.IDPEntityID,
			"public_key":                       app.PublicKey,
			"sso_endpoint":                     app.SSOEndpoint,
			"default_relay_state":              app.DefaultRelayState,
			"name_id_transform_jsonata":        app.NameIDTransformJsonata,
			"saml_attribute_transform_jsonata": app.SamlAttributeTransformJsonata,
		}

		var customAttributes []interface{}
		for _, attr := range app.CustomAttributes {
			customAttributes = append(customAttributes, convertSAMLAttributeStructToSchema(attr))
		}
		if len(customAttributes) != 0 {
			m["custom_attribute"] = customAttributes
		}

		return []interface{}{m}
	}
}

func convertScimConfigStructToSchema(scimConfig *cloudflare.AccessApplicationSCIMConfig) []interface{} {
	if scimConfig == nil {
		return []interface{}{}
	}

	config := map[string]interface{}{
		"enabled":              scimConfig.Enabled,
		"remote_uri":           scimConfig.RemoteURI,
		"idp_uid":              scimConfig.IdPUID,
		"deactivate_on_delete": cloudflare.Bool(scimConfig.DeactivateOnDelete),
		"authentication":       convertScimConfigAuthenticationStructToSchema(scimConfig.Authentication),
		"mappings":             convertScimConfigMappingsStructsToSchema(scimConfig.Mappings),
	}

	return []interface{}{config}
}

func convertScimConfigAuthenticationStructToSchema(scimAuth *cloudflare.AccessApplicationScimAuthenticationJson) []interface{} {
	if scimAuth == nil || scimAuth.Value == nil {
		return []interface{}{}
	}

	auth := map[string]interface{}{}
	switch t := scimAuth.Value.(type) {
	case *cloudflare.AccessApplicationScimAuthenticationHttpBasic:
		auth["scheme"] = t.Scheme
		auth["user"] = t.User
		auth["password"] = t.Password

	case *cloudflare.AccessApplicationScimAuthenticationOauthBearerToken:
		auth["scheme"] = t.Scheme
		auth["token"] = t.Token
	case *cloudflare.AccessApplicationScimAuthenticationOauth2:
		auth["scheme"] = t.Scheme
		auth["client_id"] = t.ClientID
		auth["client_secret"] = t.ClientSecret
		auth["authorization_url"] = t.AuthorizationURL
		auth["token_url"] = t.TokenURL
		auth["scopes"] = t.Scopes
	}

	return []interface{}{auth}
}

func convertScimConfigMappingsStructsToSchema(mappingsData []*cloudflare.AccessApplicationScimMapping) []interface{} {
	mappings := []interface{}{}

	for _, mapping := range mappingsData {
		newMapping := map[string]interface{}{
			"schema":            mapping.Schema,
			"enabled":           mapping.Enabled,
			"filter":            mapping.Filter,
			"transform_jsonata": mapping.TransformJsonata,
		}

		if mapping.Operations != nil {
			newMapping["operations"] = []interface{}{
				map[string]interface{}{
					"create": cloudflare.Bool(mapping.Operations.Create),
					"update": cloudflare.Bool(mapping.Operations.Update),
					"delete": cloudflare.Bool(mapping.Operations.Delete),
				},
			}
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}
