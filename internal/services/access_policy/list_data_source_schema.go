// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &AccessPoliciesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessPoliciesDataSource{}

func (r AccessPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the policy",
							Computed:    true,
						},
						"approval_groups": schema.ListNestedAttribute{
							Description: "Administrators who can approve a temporary authentication request.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"approvals_needed": schema.Float64Attribute{
										Description: "The number of approvals needed to obtain access.",
										Computed:    true,
									},
									"email_addresses": schema.ListAttribute{
										Description: "A list of emails that can approve the access request.",
										Computed:    true,
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
							Computed: true,
						},
						"decision": schema.StringAttribute{
							Description: "The action Access will take if a user matches this policy.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("allow", "deny", "non_identity", "bypass"),
							},
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
									},
								},
							},
						},
						"include": schema.ListNestedAttribute{
							Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
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
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
									},
								},
							},
						},
						"session_duration": schema.StringAttribute{
							Description: "The amount of time that tokens issued for the application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *AccessPoliciesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessPoliciesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
