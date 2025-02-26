// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountMembersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order results.\navailable values: \"asc\", \"desc\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Field to order results by.\navailable values: \"user.first_name\", \"user.last_name\", \"user.email\", \"status\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"user.first_name",
						"user.last_name",
						"user.email",
						"status",
					),
				},
			},
			"status": schema.StringAttribute{
				Description: "A member's status in the account.\navailable values: \"accepted\", \"pending\", \"rejected\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"accepted",
						"pending",
						"rejected",
					),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AccountMembersResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Membership identifier tag.",
							Computed:    true,
						},
						"policies": schema.ListNestedAttribute{
							Description: "Access policy for the membership",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[AccountMembersPoliciesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Policy identifier.",
										Computed:    true,
									},
									"access": schema.StringAttribute{
										Description: "Allow or deny operations against the resources.\navailable values: \"allow\", \"deny\"",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("allow", "deny"),
										},
									},
									"permission_groups": schema.ListNestedAttribute{
										Description: "A set of permission groups that are specified to the policy.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[AccountMembersPoliciesPermissionGroupsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "Identifier of the group.",
													Computed:    true,
												},
												"meta": schema.SingleNestedAttribute{
													Description: "Attributes associated to the permission group.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectType[AccountMembersPoliciesPermissionGroupsMetaDataSourceModel](ctx),
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Computed: true,
														},
														"value": schema.StringAttribute{
															Computed: true,
														},
													},
												},
												"name": schema.StringAttribute{
													Description: "Name of the group.",
													Computed:    true,
												},
											},
										},
									},
									"resource_groups": schema.ListNestedAttribute{
										Description: "A list of resource groups that the policy applies to.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[AccountMembersPoliciesResourceGroupsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "Identifier of the group.",
													Computed:    true,
												},
												"scope": schema.ListNestedAttribute{
													Description: "The scope associated to the resource group",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectListType[AccountMembersPoliciesResourceGroupsScopeDataSourceModel](ctx),
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
																Computed:    true,
															},
															"objects": schema.ListNestedAttribute{
																Description: "A list of scope objects for additional context.",
																Computed:    true,
																CustomType:  customfield.NewNestedObjectListType[AccountMembersPoliciesResourceGroupsScopeObjectsDataSourceModel](ctx),
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description: "This is a combination of pre-defined resource name and identifier (like Zone ID etc.)",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"meta": schema.SingleNestedAttribute{
													Description: "Attributes associated to the resource group.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectType[AccountMembersPoliciesResourceGroupsMetaDataSourceModel](ctx),
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Computed: true,
														},
														"value": schema.StringAttribute{
															Computed: true,
														},
													},
												},
												"name": schema.StringAttribute{
													Description: "Name of the resource group.",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"roles": schema.ListNestedAttribute{
							Description: "Roles assigned to this Member.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[AccountMembersRolesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Role identifier tag.",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "Description of role's permissions.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Role name.",
										Computed:    true,
									},
									"permissions": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"analytics": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsAnalyticsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"billing": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsBillingDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"cache_purge": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsCachePurgeDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"dns": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsDNSDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"dns_records": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsDNSRecordsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"lb": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsLBDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"logs": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsLogsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"organization": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsOrganizationDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"ssl": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsSSLDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"waf": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsWAFDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"zone_settings": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsZoneSettingsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"zones": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[AccountMembersRolesPermissionsZonesDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"read": schema.BoolAttribute{
														Computed: true,
													},
													"write": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Description: "A member's status in the account.\navailable values: \"accepted\", \"pending\"",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
							},
						},
						"user": schema.SingleNestedAttribute{
							Description: "Details of the user associated to the membership.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[AccountMembersUserDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description: "The contact email address of the user.",
									Computed:    true,
								},
								"id": schema.StringAttribute{
									Description: "Identifier",
									Computed:    true,
								},
								"first_name": schema.StringAttribute{
									Description: "User's first name",
									Computed:    true,
								},
								"last_name": schema.StringAttribute{
									Description: "User's last name",
									Computed:    true,
								},
								"two_factor_authentication_enabled": schema.BoolAttribute{
									Description: "Indicates whether two-factor authentication is enabled for the user account. Does not apply to API authentication.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *AccountMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountMembersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
