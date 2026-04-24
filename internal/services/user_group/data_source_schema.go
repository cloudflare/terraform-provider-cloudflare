// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UserGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Settings Read",
				"Account Settings Write",
				"SCIM Provisioning",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "User Group identifier tag.",
				Computed:    true,
			},
			"user_group_id": schema.StringAttribute{
				Description: "User Group identifier tag.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "Timestamp for the creation of the user group",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "Last time the user group was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the user group.",
				Computed:    true,
			},
			"policies": schema.ListNestedAttribute{
				Description: "Policies attached to the User group",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[UserGroupPoliciesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Policy identifier.",
							Computed:    true,
						},
						"access": schema.StringAttribute{
							Description: "Allow or deny operations against the resources.\nAvailable values: \"allow\", \"deny\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("allow", "deny"),
							},
						},
						"permission_groups": schema.ListNestedAttribute{
							Description: "A set of permission groups that are specified to the policy.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[UserGroupPoliciesPermissionGroupsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Identifier of the permission group.",
										Computed:    true,
									},
									"meta": schema.SingleNestedAttribute{
										Description: "Attributes associated to the permission group.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[UserGroupPoliciesPermissionGroupsMetaDataSourceModel](ctx),
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
										Description: "Name of the permission group.",
										Computed:    true,
									},
								},
							},
						},
						"resource_groups": schema.ListNestedAttribute{
							Description: "A list of resource groups that the policy applies to.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[UserGroupPoliciesResourceGroupsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Identifier of the resource group.",
										Computed:    true,
									},
									"scope": schema.ListNestedAttribute{
										Description: "The scope associated to the resource group",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[UserGroupPoliciesResourceGroupsScopeDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
													Computed:    true,
												},
												"objects": schema.ListNestedAttribute{
													Description: "A list of scope objects for additional context.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectListType[UserGroupPoliciesResourceGroupsScopeObjectsDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[UserGroupPoliciesResourceGroupsMetaDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "ID of the user group to be fetched.",
						Optional:    true,
					},
					"direction": schema.StringAttribute{
						Description: "The sort order of returned user groups by name (ascending or descending).\nAvailable values: \"asc\", \"desc\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"fuzzy_name": schema.StringAttribute{
						Description: "A string used for searching for user groups containing that substring.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the user group to be fetched.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *UserGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *UserGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("user_group_id"), path.MatchRoot("filter")),
	}
}
