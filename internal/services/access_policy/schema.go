// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

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

func (r AccessPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the policy",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"app_id": schema.StringAttribute{
				Description: "UUID",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"decision": schema.StringAttribute{
				Description: "The action Access will take if a user matches this policy.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "deny", "non_identity", "bypass"),
				},
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
				Description: "The name of the Access policy.",
				Required:    true,
			},
			"approval_groups": schema.ListNestedAttribute{
				Description: "Administrators who can approve a temporary authentication request.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"approvals_needed": schema.Float64Attribute{
							Description: "The number of approvals needed to obtain access.",
							Required:    true,
						},
						"email_addresses": schema.ListAttribute{
							Description: "A list of emails that can approve the access request.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"email_list_uuid": schema.StringAttribute{
							Description: "The UUID of an re-usable email list.",
							Optional:    true,
						},
					},
				},
			},
			"approval_required": schema.BoolAttribute{
				Description: "Requires the user to request access from an administrator at the start of each session.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
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
			"isolation_required": schema.BoolAttribute{
				Description: "Require this application to be served in an isolated browser for users matching this policy. 'Client Web Isolation' must be on for the account in order to use this feature.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"precedence": schema.Int64Attribute{
				Description: "The order of execution for this policy. Must be unique for each policy within an app.",
				Optional:    true,
			},
			"purpose_justification_prompt": schema.StringAttribute{
				Description: "A custom message that will appear on the purpose justification screen.",
				Optional:    true,
			},
			"purpose_justification_required": schema.BoolAttribute{
				Description: "Require users to enter a justification when they log in to the application.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"require": schema.ListNestedAttribute{
				Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
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
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for the application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("24h"),
			},
		},
	}
}
