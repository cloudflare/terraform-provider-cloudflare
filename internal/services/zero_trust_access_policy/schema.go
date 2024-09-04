// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the policy",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"app_id": schema.StringAttribute{
				Description:   "UUID",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"precedence": schema.Int64Attribute{
				Description: "The order of execution for this policy. Must be unique for each policy within an app.",
				Optional:    true,
			},
			"approval_required": schema.BoolAttribute{
				Description: "Requires the user to request access from an administrator at the start of each session.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"decision": schema.StringAttribute{
				Description: "The action Access will take if a user matches this policy.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"allow",
						"deny",
						"non_identity",
						"bypass",
					),
				},
			},
			"isolation_required": schema.BoolAttribute{
				Description: "Require this application to be served in an isolated browser for users matching this policy. 'Client Web Isolation' must be on for the account in order to use this feature.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				Description: "The name of the Access policy.",
				Computed:    true,
				Optional:    true,
			},
			"purpose_justification_prompt": schema.StringAttribute{
				Description: "A custom message that will appear on the purpose justification screen.",
				Computed:    true,
				Optional:    true,
			},
			"purpose_justification_required": schema.BoolAttribute{
				Description: "Require users to enter a justification when they log in to the application.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for the application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("24h"),
			},
			"approval_groups": schema.ListNestedAttribute{
				Description: "Administrators who can approve a temporary authentication request.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyApprovalGroupsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"approvals_needed": schema.Float64Attribute{
							Description: "The number of approvals needed to obtain access.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
						"email_addresses": schema.ListAttribute{
							Description: "A list of emails that can approve the access request.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"email_list_uuid": schema.StringAttribute{
							Description: "The UUID of an re-usable email list.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyExcludeModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeEmailModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeEmailListModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeEmailDomainModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeIPModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeIPListModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"certificate": schema.StringAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: jsontypes.NormalizedType{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGroupModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeAzureADModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of an Azure group.",
									Computed:    true,
									Optional:    true,
								},
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Azure identity provider.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"github_organization": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGitHubOrganizationModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGSuiteModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Google Workspace identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Google Workspace group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"okta": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeOktaModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Okta identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Okta group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"saml": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeSAMLModel](ctx),
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the SAML attribute.",
									Computed:    true,
									Optional:    true,
								},
								"attribute_value": schema.StringAttribute{
									Description: "The SAML attribute value to look for.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"service_token": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeServiceTokenModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeExternalEvaluationModel](ctx),
							Attributes: map[string]schema.Attribute{
								"evaluate_url": schema.StringAttribute{
									Description: "The API endpoint containing your business logic.",
									Computed:    true,
									Optional:    true,
								},
								"keys_url": schema.StringAttribute{
									Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"geo": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGeoModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeAuthMethodModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeDevicePostureModel](ctx),
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"include": schema.ListNestedAttribute{
				Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyIncludeModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeEmailModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeEmailListModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeEmailDomainModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeIPModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeIPListModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"certificate": schema.StringAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: jsontypes.NormalizedType{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGroupModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeAzureADModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of an Azure group.",
									Computed:    true,
									Optional:    true,
								},
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Azure identity provider.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"github_organization": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGitHubOrganizationModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGSuiteModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Google Workspace identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Google Workspace group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"okta": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeOktaModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Okta identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Okta group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"saml": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeSAMLModel](ctx),
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the SAML attribute.",
									Computed:    true,
									Optional:    true,
								},
								"attribute_value": schema.StringAttribute{
									Description: "The SAML attribute value to look for.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"service_token": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeServiceTokenModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeExternalEvaluationModel](ctx),
							Attributes: map[string]schema.Attribute{
								"evaluate_url": schema.StringAttribute{
									Description: "The API endpoint containing your business logic.",
									Computed:    true,
									Optional:    true,
								},
								"keys_url": schema.StringAttribute{
									Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"geo": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGeoModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeAuthMethodModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeDevicePostureModel](ctx),
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"require": schema.ListNestedAttribute{
				Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyRequireModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireEmailModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireEmailListModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireEmailDomainModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireIPModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireIPListModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"certificate": schema.StringAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: jsontypes.NormalizedType{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGroupModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireAzureADModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of an Azure group.",
									Computed:    true,
									Optional:    true,
								},
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Azure identity provider.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"github_organization": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGitHubOrganizationModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGSuiteModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Google Workspace identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Google Workspace group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"okta": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireOktaModel](ctx),
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Okta identity provider.",
									Computed:    true,
									Optional:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Okta group.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"saml": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireSAMLModel](ctx),
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the SAML attribute.",
									Computed:    true,
									Optional:    true,
								},
								"attribute_value": schema.StringAttribute{
									Description: "The SAML attribute value to look for.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"service_token": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireServiceTokenModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireExternalEvaluationModel](ctx),
							Attributes: map[string]schema.Attribute{
								"evaluate_url": schema.StringAttribute{
									Description: "The API endpoint containing your business logic.",
									Computed:    true,
									Optional:    true,
								},
								"keys_url": schema.StringAttribute{
									Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"geo": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGeoModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireAuthMethodModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireDevicePostureModel](ctx),
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
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

func (r *ZeroTrustAccessPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
