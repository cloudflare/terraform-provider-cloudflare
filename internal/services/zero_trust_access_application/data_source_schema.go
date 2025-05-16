// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessApplicationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "Identifier.",
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
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate to this application using their WARP session.  When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.",
				Computed:    true,
			},
			"allow_iframe": schema.BoolAttribute{
				Description: "Enables loading application content in an iFrame.",
				Computed:    true,
			},
			"app_launcher_logo_url": schema.StringAttribute{
				Description: "The image URL of the logo shown in the App Launcher header.",
				Computed:    true,
			},
			"app_launcher_visible": schema.BoolAttribute{
				Description: "Displays the application in the App Launcher.",
				Computed:    true,
			},
			"aud": schema.StringAttribute{
				Description: "Audience tag.",
				Computed:    true,
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login. You must specify only one identity provider in allowed_idps.",
				Computed:    true,
			},
			"bg_color": schema.StringAttribute{
				Description: "The background color of the App Launcher page.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"custom_deny_message": schema.StringAttribute{
				Description: "The custom error message shown to a user when they are denied access to the application.",
				Computed:    true,
			},
			"custom_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing identity-based rules.",
				Computed:    true,
			},
			"custom_non_identity_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing non-identity rules.",
				Computed:    true,
			},
			"domain": schema.StringAttribute{
				Description: "The primary hostname and path secured by Access. This domain will be displayed if the app is visible in the App Launcher.",
				Computed:    true,
			},
			"enable_binding_cookie": schema.BoolAttribute{
				Description: "Enables the binding cookie, which increases security against compromised authorization tokens and CSRF attacks.",
				Computed:    true,
			},
			"header_bg_color": schema.StringAttribute{
				Description: "The background color of the App Launcher header.",
				Computed:    true,
			},
			"http_only_cookie_attribute": schema.BoolAttribute{
				Description: "Enables the HttpOnly cookie attribute, which increases security against XSS attacks.",
				Computed:    true,
			},
			"logo_url": schema.StringAttribute{
				Description: "The image URL for the logo shown in the App Launcher dashboard.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the application.",
				Computed:    true,
			},
			"options_preflight_bypass": schema.BoolAttribute{
				Description: "Allows options preflight requests to bypass Access authentication and go directly to the origin. Cannot turn on if cors_headers is set.",
				Computed:    true,
			},
			"path_cookie_attribute": schema.BoolAttribute{
				Description: "Enables cookie paths to scope an application's JWT to the application path. If disabled, the JWT will scope to the hostname by default",
				Computed:    true,
			},
			"read_service_tokens_from_header": schema.StringAttribute{
				Description: "Allows matching Access Service Tokens passed HTTP in a single header with this name.\nThis works as an alternative to the (CF-Access-Client-Id, CF-Access-Client-Secret) pair of headers.\nThe header value will be interpreted as a json object similar to: \n  {\n    \"cf-access-client-id\": \"88bf3b6d86161464f6509f7219099e57.access.example.com\",\n    \"cf-access-client-secret\": \"bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5\"\n  }",
				Computed:    true,
			},
			"same_site_cookie_attribute": schema.StringAttribute{
				Description: "Sets the SameSite cookie setting, which provides increased security against CSRF attacks.",
				Computed:    true,
			},
			"service_auth_401_redirect": schema.BoolAttribute{
				Description: "Returns a 401 status code when the request is blocked by a Service Auth policy.",
				Computed:    true,
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for this application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. Note: unsupported for infrastructure type applications.",
				Computed:    true,
			},
			"skip_app_launcher_login_page": schema.BoolAttribute{
				Description: "Determines when to skip the App Launcher landing page.",
				Computed:    true,
			},
			"skip_interstitial": schema.BoolAttribute{
				Description: "Enables automatic authentication through cloudflared.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The application type.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"allowed_idps": schema.ListAttribute{
				Description: "The identity providers your users can select when connecting to this application. Defaults to all IdPs configured in your account.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"custom_pages": schema.ListAttribute{
				Description: "The custom pages that will be displayed when applicable for this application",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"self_hosted_domains": schema.ListAttribute{
				Description:        "List of public domains that Access will secure. This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         customfield.NewListType[types.String](ctx),
				ElementType:        types.StringType,
			},
			"tags": schema.ListAttribute{
				Description: "The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"cors_headers": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationCORSHeadersDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"allow_all_headers": schema.BoolAttribute{
						Description: "Allows all HTTP request headers.",
						Computed:    true,
					},
					"allow_all_methods": schema.BoolAttribute{
						Description: "Allows all HTTP request methods.",
						Computed:    true,
					},
					"allow_all_origins": schema.BoolAttribute{
						Description: "Allows all origins.",
						Computed:    true,
					},
					"allow_credentials": schema.BoolAttribute{
						Description: "When set to `true`, includes credentials (cookies, authorization headers, or TLS client certificates) with requests.",
						Computed:    true,
					},
					"allowed_headers": schema.ListAttribute{
						Description: "Allowed HTTP request headers.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"allowed_methods": schema.ListAttribute{
						Description: "Allowed HTTP request methods.",
						Computed:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"GET",
									"POST",
									"HEAD",
									"PUT",
									"DELETE",
									"CONNECT",
									"OPTIONS",
									"TRACE",
									"PATCH",
								),
							),
						},
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"allowed_origins": schema.ListAttribute{
						Description: "Allowed origins.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"max_age": schema.Float64Attribute{
						Description: "The maximum number of seconds the results of a preflight request can be cached.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(-1, 86400),
						},
					},
				},
			},
			"destinations": schema.ListNestedAttribute{
				Description: "List of destinations secured by Access. This supersedes `self_hosted_domains` to allow for more flexibility in defining different types of domains. If `destinations` are provided, then `self_hosted_domains` will be ignored.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationDestinationsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: `Available values: "public", "private".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("public", "private"),
							},
						},
						"uri": schema.StringAttribute{
							Description: "The URI of the destination. Public destinations' URIs can include a domain and path with [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).",
							Computed:    true,
						},
						"cidr": schema.StringAttribute{
							Description: "The CIDR range of the destination. Single IPs will be computed as /32.",
							Computed:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname of the destination. Matches a valid SNI served by an HTTPS origin.",
							Computed:    true,
						},
						"l4_protocol": schema.StringAttribute{
							Description: "The L4 protocol of the destination. When omitted, both UDP and TCP traffic will match.\nAvailable values: \"tcp\", \"udp\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("tcp", "udp"),
							},
						},
						"port_range": schema.StringAttribute{
							Description: "The port range of the destination. Can be a single port or a range of ports. When omitted, all ports will match.",
							Computed:    true,
						},
						"vnet_id": schema.StringAttribute{
							Description: "The VNET ID to match the destination. When omitted, all VNETs will match.",
							Computed:    true,
						},
					},
				},
			},
			"footer_links": schema.ListNestedAttribute{
				Description: "The links in the App Launcher footer.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationFooterLinksDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "The hypertext in the footer link.",
							Computed:    true,
						},
						"url": schema.StringAttribute{
							Description: "the hyperlink in the footer link.",
							Computed:    true,
						},
					},
				},
			},
			"landing_page_design": schema.SingleNestedAttribute{
				Description: "The design of the App Launcher landing page shown to users when they log in.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationLandingPageDesignDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"button_color": schema.StringAttribute{
						Description: "The background color of the log in button on the landing page.",
						Computed:    true,
					},
					"button_text_color": schema.StringAttribute{
						Description: "The color of the text in the log in button on the landing page.",
						Computed:    true,
					},
					"image_url": schema.StringAttribute{
						Description: "The URL of the image shown on the landing page.",
						Computed:    true,
					},
					"message": schema.StringAttribute{
						Description: "The message shown on the landing page.",
						Computed:    true,
					},
					"title": schema.StringAttribute{
						Description: "The title shown on the landing page.",
						Computed:    true,
					},
				},
			},
			"policies": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the policy",
							Computed:    true,
						},
						"approval_groups": schema.ListNestedAttribute{
							Description: "Administrators who can approve a temporary authentication request.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesApprovalGroupsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"approvals_needed": schema.Float64Attribute{
										Description: "The number of approvals needed to obtain access.",
										Computed:    true,
										Validators: []validator.Float64{
											float64validator.AtLeast(0),
										},
									},
									"email_addresses": schema.ListAttribute{
										Description: "A list of emails that can approve the access request.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"email_list_uuid": schema.StringAttribute{
										Description: "The UUID of an re-usable email list.",
										Computed:    true,
									},
								},
							},
						},
						"approval_required": schema.BoolAttribute{
							Description: "Requires the user to request access from an administrator at the start of each session.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"decision": schema.StringAttribute{
							Description: "The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.\nAvailable values: \"allow\", \"deny\", \"non_identity\", \"bypass\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"allow",
									"deny",
									"non_identity",
									"bypass",
								),
							},
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesExcludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAuthContextDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Computed:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeCertificateDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeCommonNameDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Computed:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeDevicePostureDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEmailDomainDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEveryoneDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Computed:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Computed:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Computed:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeLoginMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeServiceTokenDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"include": schema.ListNestedAttribute{
							Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesIncludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAuthContextDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Computed:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeCertificateDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeCommonNameDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Computed:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeDevicePostureDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEmailDomainDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEveryoneDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Computed:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Computed:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Computed:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeLoginMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeServiceTokenDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"isolation_required": schema.BoolAttribute{
							Description: "Require this application to be served in an isolated browser for users matching this policy. 'Client Web Isolation' must be on for the account in order to use this feature.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Access policy.",
							Computed:    true,
						},
						"precedence": schema.Int64Attribute{
							Description: "The order of execution for this policy. Must be unique for each policy within an app.",
							Computed:    true,
						},
						"purpose_justification_prompt": schema.StringAttribute{
							Description: "A custom message that will appear on the purpose justification screen.",
							Computed:    true,
						},
						"purpose_justification_required": schema.BoolAttribute{
							Description: "Require users to enter a justification when they log in to the application.",
							Computed:    true,
						},
						"require": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesRequireDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAuthContextDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Computed:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireCertificateDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireCommonNameDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Computed:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireDevicePostureDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEmailDomainDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEveryoneDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Computed:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Computed:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Computed:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireLoginMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireServiceTokenDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"session_duration": schema.StringAttribute{
							Description: "The amount of time that tokens issued for the application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"connection_rules": schema.SingleNestedAttribute{
							Description: "The rules that define how users may connect to the targets secured by your application.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesConnectionRulesDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ssh": schema.SingleNestedAttribute{
									Description: "The SSH-specific rules that define how users may connect to the targets secured by your application.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesConnectionRulesSSHDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"usernames": schema.ListAttribute{
											Description: "Contains the Unix usernames that may be used when connecting over SSH.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"allow_email_alias": schema.BoolAttribute{
											Description: "Enables using Identity Provider email alias as SSH username.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"saas_app": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Description: "Optional identifier indicating the authentication protocol used for the saas app. Required for OIDC. Default if unset is \"saml\"\nAvailable values: \"saml\", \"oidc\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("saml", "oidc"),
						},
					},
					"consumer_service_url": schema.StringAttribute{
						Description: "The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"custom_attributes": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSaaSAppCustomAttributesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"friendly_name": schema.StringAttribute{
									Description: "The SAML FriendlyName of the attribute.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the attribute.",
									Computed:    true,
								},
								"name_format": schema.StringAttribute{
									Description: "A globally unique name for an identity or service provider.\nAvailable values: \"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified\", \"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\", \"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified",
											"urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
											"urn:oasis:names:tc:SAML:2.0:attrname-format:uri",
										),
									},
								},
								"required": schema.BoolAttribute{
									Description: "If the attribute is required when building a SAML assertion.",
									Computed:    true,
								},
								"source": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "The name of the IdP attribute.",
											Computed:    true,
										},
										"name_by_idp": schema.ListNestedAttribute{
											Description: "A mapping from IdP ID to attribute name.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceNameByIdPDataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"idp_id": schema.StringAttribute{
														Description: "The UID of the IdP.",
														Computed:    true,
													},
													"source_name": schema.StringAttribute{
														Description: "The name of the IdP provided attribute.",
														Computed:    true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"default_relay_state": schema.StringAttribute{
						Description: "The URL that the user will be redirected to after a successful login for IDP initiated logins.",
						Computed:    true,
					},
					"idp_entity_id": schema.StringAttribute{
						Description: "The unique identifier for your SaaS application.",
						Computed:    true,
					},
					"name_id_format": schema.StringAttribute{
						Description: "The format of the name identifier sent to the SaaS application.\nAvailable values: \"id\", \"email\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "email"),
						},
					},
					"name_id_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.",
						Computed:    true,
					},
					"public_key": schema.StringAttribute{
						Description: "The Access public certificate that will be used to verify your identity.",
						Computed:    true,
					},
					"saml_attribute_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata] (https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.",
						Computed:    true,
					},
					"sp_entity_id": schema.StringAttribute{
						Description: "A globally unique name for an identity or service provider.",
						Computed:    true,
					},
					"sso_endpoint": schema.StringAttribute{
						Description: "The endpoint where your SaaS application will send login requests.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"access_token_lifetime": schema.StringAttribute{
						Description: "The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must be greater than or equal to 1m and less than or equal to 24h.",
						Computed:    true,
					},
					"allow_pkce_without_client_secret": schema.BoolAttribute{
						Description: "If client secret should be required on the token endpoint when authorization_code_with_pkce grant is used.",
						Computed:    true,
					},
					"app_launcher_url": schema.StringAttribute{
						Description: "The URL where this applications tile redirects users",
						Computed:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "The application client id",
						Computed:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "The application client secret, only returned on POST request.",
						Computed:    true,
						Sensitive:   true,
					},
					"custom_claims": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSaaSAppCustomClaimsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "The name of the claim.",
									Computed:    true,
								},
								"required": schema.BoolAttribute{
									Description: "If the claim is required when building an OIDC token.",
									Computed:    true,
								},
								"scope": schema.StringAttribute{
									Description: "The scope of the claim.\nAvailable values: \"groups\", \"profile\", \"email\", \"openid\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"groups",
											"profile",
											"email",
											"openid",
										),
									},
								},
								"source": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "The name of the IdP claim.",
											Computed:    true,
										},
										"name_by_idp": schema.MapAttribute{
											Description: "A mapping from IdP ID to claim name.",
											Computed:    true,
											CustomType:  customfield.NewMapType[types.String](ctx),
											ElementType: types.StringType,
										},
									},
								},
							},
						},
					},
					"grant_types": schema.ListAttribute{
						Description: "The OIDC flows supported by this application",
						Computed:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"authorization_code",
									"authorization_code_with_pkce",
									"refresh_tokens",
									"hybrid",
									"implicit",
								),
							),
						},
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"group_filter_regex": schema.StringAttribute{
						Description: "A regex to filter Cloudflare groups returned in ID token and userinfo endpoint",
						Computed:    true,
					},
					"hybrid_and_implicit_options": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"return_access_token_from_authorization_endpoint": schema.BoolAttribute{
								Description: "If an Access Token should be returned from the OIDC Authorization endpoint",
								Computed:    true,
							},
							"return_id_token_from_authorization_endpoint": schema.BoolAttribute{
								Description: "If an ID Token should be returned from the OIDC Authorization endpoint",
								Computed:    true,
							},
						},
					},
					"redirect_uris": schema.ListAttribute{
						Description: "The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"refresh_token_options": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"lifetime": schema.StringAttribute{
								Description: "How long a refresh token will be valid for after creation. Valid units are m,h,d. Must be longer than 1m.",
								Computed:    true,
							},
						},
					},
					"scopes": schema.ListAttribute{
						Description: `Define the user information shared with access, "offline_access" scope will be automatically enabled if refresh tokens are enabled`,
						Computed:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"openid",
									"groups",
									"email",
									"profile",
								),
							),
						},
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
			"scim_config": schema.SingleNestedAttribute{
				Description: "Configuration for provisioning to this application via SCIM. This is currently in closed beta.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationSCIMConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"idp_uid": schema.StringAttribute{
						Description: "The UID of the IdP to use as the source for SCIM resources to provision to this application.",
						Computed:    true,
					},
					"remote_uri": schema.StringAttribute{
						Description: "The base URI for the application's SCIM-compatible API.",
						Computed:    true,
					},
					"authentication": schema.SingleNestedAttribute{
						Description: "Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationSCIMConfigAuthenticationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"password": schema.StringAttribute{
								Description: "Password used to authenticate with the remote SCIM service.",
								Computed:    true,
							},
							"scheme": schema.StringAttribute{
								Description: "The authentication scheme to use when making SCIM requests to this application.\nAvailable values: \"httpbasic\", \"oauthbearertoken\", \"oauth2\", \"access_service_token\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"httpbasic",
										"oauthbearertoken",
										"oauth2",
										"access_service_token",
									),
								},
							},
							"user": schema.StringAttribute{
								Description: "User name used to authenticate with the remote SCIM service.",
								Computed:    true,
							},
							"token": schema.StringAttribute{
								Description: "Token used to authenticate with the remote SCIM service.",
								Computed:    true,
								Sensitive:   true,
							},
							"authorization_url": schema.StringAttribute{
								Description: "URL used to generate the auth code used during token generation.",
								Computed:    true,
							},
							"client_id": schema.StringAttribute{
								Description: "Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.",
								Computed:    true,
							},
							"client_secret": schema.StringAttribute{
								Description: "Secret used to authenticate when generating a token for authenticating with the remove SCIM service.",
								Computed:    true,
								Sensitive:   true,
							},
							"token_url": schema.StringAttribute{
								Description: "URL used to generate the token used to authenticate with the remote SCIM service.",
								Computed:    true,
							},
							"scopes": schema.ListAttribute{
								Description: "The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"deactivate_on_delete": schema.BoolAttribute{
						Description: "If false, propagates DELETE requests to the target application for SCIM resources. If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.",
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether SCIM provisioning is turned on for this application.",
						Computed:    true,
					},
					"mappings": schema.ListNestedAttribute{
						Description: "A list of mappings to apply to SCIM resources before provisioning them in this application. These can transform or filter the resources to be provisioned.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSCIMConfigMappingsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"schema": schema.StringAttribute{
									Description: "Which SCIM resource type this mapping applies to.",
									Computed:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether or not this mapping is enabled.",
									Computed:    true,
								},
								"filter": schema.StringAttribute{
									Description: "A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.",
									Computed:    true,
								},
								"operations": schema.SingleNestedAttribute{
									Description: "Whether or not this mapping applies to creates, updates, or deletes.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationSCIMConfigMappingsOperationsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"create": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to create (POST) operations.",
											Computed:    true,
										},
										"delete": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to DELETE operations.",
											Computed:    true,
										},
										"update": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to update (PATCH/PUT) operations.",
											Computed:    true,
										},
									},
								},
								"strictness": schema.StringAttribute{
									Description: "The level of adherence to outbound resource schemas when provisioning to this mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown values to the target.\nAvailable values: \"strict\", \"passthrough\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("strict", "passthrough"),
									},
								},
								"transform_jsonata": schema.StringAttribute{
									Description: "A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
			"target_criteria": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationTargetCriteriaDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Int64Attribute{
							Description: "The port that the targets use for the chosen communication protocol. A port cannot be assigned to multiple protocols.",
							Computed:    true,
						},
						"protocol": schema.StringAttribute{
							Description: "The communication protocol your application secures.\nAvailable values: \"ssh\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("ssh"),
							},
						},
						"target_attributes": schema.MapAttribute{
							Description: "Contains a map of target attribute keys to target attribute values.",
							Computed:    true,
							CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
							ElementType: types.ListType{
								ElemType: types.StringType,
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"aud": schema.StringAttribute{
						Description: "The aud of the app.",
						Optional:    true,
					},
					"domain": schema.StringAttribute{
						Description: "The domain of the app.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the app.",
						Optional:    true,
					},
					"search": schema.StringAttribute{
						Description: "Search for apps by other listed query parameters.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessApplicationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("app_id"), path.MatchRoot("filter")),
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
