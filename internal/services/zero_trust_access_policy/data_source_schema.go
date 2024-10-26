// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessPolicyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"policy_id": schema.StringAttribute{
				Description: "The UUID of the policy",
				Optional:    true,
			},
			"app_count": schema.Int64Attribute{
				Description: "Number of access applications currently using this policy.",
				Computed:    true,
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
				Description: "The action Access will take if a user matches this policy.",
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
			"id": schema.StringAttribute{
				Description: "The UUID of the policy",
				Computed:    true,
			},
			"isolation_required": schema.BoolAttribute{
				Description: "Require this application to be served in an isolated browser for users matching this policy. 'Client Web Isolation' must be on for the account in order to use this feature.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Access policy.",
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
			"reusable": schema.BoolAttribute{
				Computed: true,
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for the application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"approval_groups": schema.ListNestedAttribute{
				Description: "Administrators who can approve a temporary authentication request.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyApprovalGroupsDataSourceModel](ctx),
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
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyExcludeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeAzureADDataSourceModel](ctx),
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
						"github_organization": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"identity_provider_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Computed:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel](ctx),
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
						"okta": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeExternalEvaluationDataSourceModel](ctx),
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
						"geo": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyExcludeDevicePostureDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyIncludeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeAzureADDataSourceModel](ctx),
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
						"github_organization": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"identity_provider_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Computed:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel](ctx),
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
						"okta": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeExternalEvaluationDataSourceModel](ctx),
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
						"geo": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyIncludeDevicePostureDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
			"require": schema.ListNestedAttribute{
				Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPolicyRequireDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireAzureADDataSourceModel](ctx),
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
						"github_organization": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"identity_provider_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Computed:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGSuiteDataSourceModel](ctx),
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
						"okta": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireExternalEvaluationDataSourceModel](ctx),
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
						"geo": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPolicyRequireDevicePostureDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessPolicyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("policy_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("policy_id")),
	}
}
