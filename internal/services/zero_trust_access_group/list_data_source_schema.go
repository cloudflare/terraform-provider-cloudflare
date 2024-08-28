// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupsExcludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeEmailDomainDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.StringAttribute{
										Computed:   true,
										CustomType: jsontypes.NormalizedType{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeServiceTokenDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeExternalEvaluationDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsExcludeDevicePostureDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupsIncludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeEmailDomainDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.StringAttribute{
										Computed:   true,
										CustomType: jsontypes.NormalizedType{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeServiceTokenDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeExternalEvaluationDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIncludeDevicePostureDataSourceModel](ctx),
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
						"is_default": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupsIsDefaultDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultEmailDomainDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.StringAttribute{
										Computed:   true,
										CustomType: jsontypes.NormalizedType{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultServiceTokenDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultExternalEvaluationDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsIsDefaultDevicePostureDataSourceModel](ctx),
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
						"name": schema.StringAttribute{
							Description: "The name of the Access group.",
							Computed:    true,
						},
						"require": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupsRequireDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireEmailDomainDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.StringAttribute{
										Computed:   true,
										CustomType: jsontypes.NormalizedType{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connection_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"email": schema.StringAttribute{
												Description: "The email of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireServiceTokenDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireExternalEvaluationDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupsRequireDevicePostureDataSourceModel](ctx),
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
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustAccessGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
