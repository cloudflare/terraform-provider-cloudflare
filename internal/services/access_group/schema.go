// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r AccessGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"include": schema.ListNestedAttribute{
				Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Required:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Required:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Required:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Optional:    true,
						},
						"ip": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Required:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Required:    true,
								},
							},
						},
						"certificate": schema.StringAttribute{
							Optional: true,
						},
						"group": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Required:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of an Azure group.",
									Required:    true,
								},
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Azure identity provider.",
									Required:    true,
								},
							},
						},
						"github_organization": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Required:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Required:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Google Workspace identity provider.",
									Required:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Google Workspace group.",
									Required:    true,
								},
							},
						},
						"okta": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Okta identity provider.",
									Required:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Okta group.",
									Required:    true,
								},
							},
						},
						"saml": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the SAML attribute.",
									Required:    true,
								},
								"attribute_value": schema.StringAttribute{
									Description: "The SAML attribute value to look for.",
									Required:    true,
								},
							},
						},
						"service_token": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Required:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Optional:    true,
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"evaluate_url": schema.StringAttribute{
									Description: "The API endpoint containing your business logic.",
									Required:    true,
								},
								"keys_url": schema.StringAttribute{
									Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
									Required:    true,
								},
							},
						},
						"geo": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Required:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
									Required:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Required:    true,
								},
							},
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the Access group.",
				Required:    true,
			},
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Required:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Required:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Required:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Optional:    true,
						},
						"ip": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Required:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Required:    true,
								},
							},
						},
						"certificate": schema.StringAttribute{
							Optional: true,
						},
						"group": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Required:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of an Azure group.",
									Required:    true,
								},
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Azure identity provider.",
									Required:    true,
								},
							},
						},
						"github_organization": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Required:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Required:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Google Workspace identity provider.",
									Required:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Google Workspace group.",
									Required:    true,
								},
							},
						},
						"okta": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Okta identity provider.",
									Required:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Okta group.",
									Required:    true,
								},
							},
						},
						"saml": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the SAML attribute.",
									Required:    true,
								},
								"attribute_value": schema.StringAttribute{
									Description: "The SAML attribute value to look for.",
									Required:    true,
								},
							},
						},
						"service_token": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Required:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Optional:    true,
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"evaluate_url": schema.StringAttribute{
									Description: "The API endpoint containing your business logic.",
									Required:    true,
								},
								"keys_url": schema.StringAttribute{
									Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
									Required:    true,
								},
							},
						},
						"geo": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Required:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
									Required:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Required:    true,
								},
							},
						},
					},
				},
			},
			"is_default": schema.BoolAttribute{
				Description: "Whether this is the default group",
				Optional:    true,
			},
			"require": schema.ListNestedAttribute{
				Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The email of the user.",
									Required:    true,
								},
							},
						},
						"email_list": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created email list.",
									Required:    true,
								},
							},
						},
						"email_domain": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"domain": schema.StringAttribute{
									Description: "The email domain to match.",
									Required:    true,
								},
							},
						},
						"everyone": schema.StringAttribute{
							Description: "An empty object which matches on all users.",
							Optional:    true,
						},
						"ip": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "An IPv4 or IPv6 CIDR block.",
									Required:    true,
								},
							},
						},
						"ip_list": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created IP list.",
									Required:    true,
								},
							},
						},
						"certificate": schema.StringAttribute{
							Optional: true,
						},
						"group": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Required:    true,
								},
							},
						},
						"azure_ad": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of an Azure group.",
									Required:    true,
								},
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Azure identity provider.",
									Required:    true,
								},
							},
						},
						"github_organization": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Github identity provider.",
									Required:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the organization.",
									Required:    true,
								},
							},
						},
						"gsuite": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Google Workspace identity provider.",
									Required:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Google Workspace group.",
									Required:    true,
								},
							},
						},
						"okta": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"connection_id": schema.StringAttribute{
									Description: "The ID of your Okta identity provider.",
									Required:    true,
								},
								"email": schema.StringAttribute{
									Description: "The email of the Okta group.",
									Required:    true,
								},
							},
						},
						"saml": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description: "The name of the SAML attribute.",
									Required:    true,
								},
								"attribute_value": schema.StringAttribute{
									Description: "The SAML attribute value to look for.",
									Required:    true,
								},
							},
						},
						"service_token": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"token_id": schema.StringAttribute{
									Description: "The ID of a Service Token.",
									Required:    true,
								},
							},
						},
						"any_valid_service_token": schema.StringAttribute{
							Description: "An empty object which matches on all service tokens.",
							Optional:    true,
						},
						"external_evaluation": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"evaluate_url": schema.StringAttribute{
									Description: "The API endpoint containing your business logic.",
									Required:    true,
								},
								"keys_url": schema.StringAttribute{
									Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
									Required:    true,
								},
							},
						},
						"geo": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"country_code": schema.StringAttribute{
									Description: "The country code that should be matched.",
									Required:    true,
								},
							},
						},
						"auth_method": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"auth_method": schema.StringAttribute{
									Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.",
									Required:    true,
								},
							},
						},
						"device_posture": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"integration_uid": schema.StringAttribute{
									Description: "The ID of a device posture integration.",
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}
