// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"group_id": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Access group.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupExcludeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEveryoneDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeAzureADDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGSuiteDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeAnyValidServiceTokenDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupIncludeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEveryoneDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeAzureADDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGSuiteDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeAnyValidServiceTokenDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupIsDefaultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultEveryoneDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultAnyValidServiceTokenDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel](ctx),
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
				Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupRequireDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEmailDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Computed:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEmailListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Computed:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEmailDomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Computed:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEveryoneDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"ip": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Computed:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireIPListDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Computed:    true,
								},
							},
						},
						"certificate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireCertificateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{},
						},
						"group": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGroupDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Computed:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireAzureADDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGSuiteDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireOktaDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireSAMLDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireServiceTokenDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Computed:    true,
								},
							},
						},
						"any_valid_service_token": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all service tokens.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireAnyValidServiceTokenDataSourceModel](ctx),
							Attributes:  map[string]schema.Attribute{},
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGeoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Computed:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireAuthMethodDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
									Computed:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireDevicePostureDataSourceModel](ctx),
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
				Validators: []validator.Object{
					objectvalidator.ExactlyOneOf(path.MatchRelative().AtName("account_id"), path.MatchRelative().AtName("zone_id")),
				},
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
						Optional:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("group_id")),
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.Conflicting(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.Conflicting(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}
