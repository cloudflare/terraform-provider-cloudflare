// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessGroupResource)(nil)

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
			"is_default": schema.BoolAttribute{
				Description: "Whether this is the default group",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Access group.",
				Computed:    true,
				Optional:    true,
			},
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupExcludeModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEmailModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEmailListModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeEmailDomainModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeIPModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeIPListModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGroupModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeAzureADModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGitHubOrganizationModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGSuiteModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeOktaModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeSAMLModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeServiceTokenModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeExternalEvaluationModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeGeoModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeAuthMethodModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupExcludeDevicePostureModel](ctx),
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupIncludeModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEmailModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEmailListModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeEmailDomainModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeIPModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeIPListModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGroupModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeAzureADModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGitHubOrganizationModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGSuiteModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeOktaModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeSAMLModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeServiceTokenModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeExternalEvaluationModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeGeoModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeAuthMethodModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupIncludeDevicePostureModel](ctx),
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
				Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessGroupRequireModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEmailModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEmailListModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireEmailDomainModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireIPModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireIPListModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGroupModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireAzureADModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGitHubOrganizationModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGSuiteModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireOktaModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireSAMLModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireServiceTokenModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireExternalEvaluationModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireGeoModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireAuthMethodModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[ZeroTrustAccessGroupRequireDevicePostureModel](ctx),
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

func (r *ZeroTrustAccessGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
