// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*APITokensDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"direction": schema.StringAttribute{
				Description: "Direction to order results.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[APITokensResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Token identifier tag.",
							Computed:    true,
						},
						"condition": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"request_ip": schema.SingleNestedAttribute{
									Description: "Client IP restrictions.",
									Computed:    true,
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"in": schema.ListAttribute{
											Description: "List of IPv4/IPv6 CIDR addresses.",
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
										"not_in": schema.ListAttribute{
											Description: "List of IPv4/IPv6 CIDR addresses.",
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
						},
						"expires_on": schema.StringAttribute{
							Description: "The expiration time on or after which the JWT MUST NOT be accepted for processing.",
							Computed:    true,
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"issued_on": schema.StringAttribute{
							Description: "The time on which the token was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_used_on": schema.StringAttribute{
							Description: "Last time the token was used.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"modified_on": schema.StringAttribute{
							Description: "Last time the token was modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "Token name.",
							Computed:    true,
							Optional:    true,
						},
						"not_before": schema.StringAttribute{
							Description: "The time before which the token MUST NOT be accepted for processing.",
							Computed:    true,
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"policies": schema.ListNestedAttribute{
							Description: "List of access policies assigned to the token.",
							Computed:    true,
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Policy identifier.",
										Computed:    true,
									},
									"effect": schema.StringAttribute{
										Description: "Allow or deny operations against the resources.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("allow", "deny"),
										},
									},
									"permission_groups": schema.ListNestedAttribute{
										Description: "A set of permission groups that are specified to the policy.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[APITokensPoliciesPermissionGroupsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "Identifier of the group.",
													Computed:    true,
												},
												"meta": schema.SingleNestedAttribute{
													Description: "Attributes associated to the permission group.",
													Computed:    true,
													Optional:    true,
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Computed: true,
															Optional: true,
														},
														"value": schema.StringAttribute{
															Computed: true,
															Optional: true,
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
									"resources": schema.SingleNestedAttribute{
										Description: "A list of resource names that the policy applies to.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[APITokensPoliciesResourcesDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"resource": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
											"scope": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Description: "Status of the token.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"active",
									"disabled",
									"expired",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *APITokensDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *APITokensDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}