// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the policy",
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
			"application_id": schema.StringAttribute{
				Description: "The ID of the application the policy is associated with.",
				Required:    true,
			},
			"decision": schema.StringAttribute{
				Description: "The action Access will take if a user matches this policy. Available values: `allow`, `deny`, `non_identity`, `bypass`.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"allow",
						"deny",
						"non_identity",
						"bypass",
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the Access policy.",
				Required:    true,
			},
			"precedence": schema.Int64Attribute{
				Description: "The precedence to determine the order of policy execution. A lower number indicates a higher precedence.",
				Required:    true,
			},
			"approval_required": schema.BoolAttribute{
				Description: "Requires the user to request access from an administrator at the start of each session.",
				Optional:    true,
			},
			"isolation_required": schema.BoolAttribute{
				Description: "Require this application to be served in an isolated browser for users matching this policy.",
				Optional:    true,
			},
			"purpose_justification_prompt": schema.StringAttribute{
				Description: "A custom message that will appear on the purpose justification screen.",
				Optional:    true,
			},
			"purpose_justification_required": schema.BoolAttribute{
				Description: "Require users to enter a justification when they log in to the application.",
				Optional:    true,
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for the application will be valid. Must be in the format `300ms` or `2h45m`.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("24h"),
			},
			"exclude": schema.ListNestedAttribute{
				Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Exclude rules.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Required:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Optional:    true,
							Attributes:  map[string]schema.Attribute{},
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
					},
				},
			},
			"include": schema.ListNestedAttribute{
				Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Required:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Optional:    true,
							Attributes:  map[string]schema.Attribute{},
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
					},
				},
			},
			"require": schema.ListNestedAttribute{
				Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of a previously created Access group.",
									Required:    true,
								},
							},
						},
						"everyone": schema.SingleNestedAttribute{
							Description: "An empty object which matches on all users.",
							Optional:    true,
							Attributes:  map[string]schema.Attribute{},
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
					},
				},
			},
		},
	}
}
