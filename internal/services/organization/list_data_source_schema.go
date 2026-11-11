// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*OrganizationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"User Details Read",
				"User Details Write",
			},
			MarkdownDescription: "User API tokens can list and create Organizations, but reading, updating, or deleting an Organization requires a Global API key and your registered email address. The `cloudflare_organization` Terraform resource therefore requires a Global API key for its full lifecycle.\n\nSee [Organizations authentication limitations](https://developers.cloudflare.com/fundamentals/organizations/limitations/#api-authentication) for credential compatibility by operation. Each endpoint’s **Security** section lists its supported authentication methods, with **Accepted Permissions** where applicable.",
		}.String(),
		Attributes: map[string]schema.Attribute{
			"direction": schema.StringAttribute{
				Description: "Sort direction for the `order_by` field. Defaults to `asc` when `order_by` is\nspecified.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order_by": schema.StringAttribute{
				Description: "Field to order results by. Currently supported values: `name`.\nWhen not specified, results are ordered by hierarchy.\nAvailable values: \"name\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("name"),
				},
			},
			"page_size": schema.Int64Attribute{
				Description: "The amount of items to return. Defaults to 10.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
				},
			},
			"page_token": schema.StringAttribute{
				Description: "An opaque token returned from the last list response that when\nprovided will retrieve the next page.\n\nParameters used to filter the retrieved list must remain in subsequent\nrequests with a page token.",
				Optional:    true,
			},
			"id": schema.ListAttribute{
				Description: "Only return organizations with the specified IDs (ex. id=foo&id=bar). Send multiple elements\nby repeating the query value.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"containing": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that contain this particular\naccount.",
						Optional:    true,
					},
					"organization": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that contain this particular\norganization.",
						Optional:    true,
					},
					"user": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that contain this particular\nuser.\n\nIMPORTANT: Just because an organization \"contains\" a user is not a\nrepresentation of any authorization or privilege to manage any resources\ntherein. An organization \"containing\" a user simply means the user is managed by\nthat organization.",
						Optional:    true,
					},
				},
			},
			"member": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description: "Filter organizations to those where the given user, by email, is an active member of\nthat exact organization. This does not include hierarchy-derived access.",
						Optional:    true,
					},
				},
			},
			"name": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"contains": schema.StringAttribute{
						Description: "(case-insensitive) Filter the list of organizations to where the name contains a particular\nstring.",
						Optional:    true,
					},
					"ends_with": schema.StringAttribute{
						Description: "(case-insensitive) Filter the list of organizations to where the name ends with a particular\nstring.",
						Optional:    true,
					},
					"starts_with": schema.StringAttribute{
						Description: "(case-insensitive) Filter the list of organizations to where the name starts with a\nparticular string.",
						Optional:    true,
					},
				},
			},
			"parent": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that are a sub-organization\nof the specified organization.\n\n\"null\" is a valid value to provide for this parameter. It means \"where\nan organization has no parent (i.e. it is a 'root' organization).\"",
						Computed:    true,
					},
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
				CustomType:  customfield.NewNestedObjectListType[OrganizationsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"create_time": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"meta": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[OrganizationsMetaDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"are_child_accounts_accessible_by_org_members": schema.BoolAttribute{
									Computed: true,
								},
								"hierarchy_tags": schema.ListAttribute{
									Description: "Ordered chain of organization tags from the root organization down to\n(and including) this organization itself. Root organizations return a\nsingle-element array containing their own tag; sub-organizations return\n`[rootTag, ...intermediateTags, parentTag, selfTag]`. Useful for\nconstructing authorization scopes that need to cover every ancestor\nin the hierarchy.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"managed_by": schema.StringAttribute{
									Computed: true,
								},
								"tenant_flags": schema.SingleNestedAttribute{
									Description: "Enable features for Organizations.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[OrganizationsMetaTenantFlagsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"account_creation": schema.StringAttribute{
											Computed: true,
										},
										"account_creation_applies_tenant_defaults": schema.StringAttribute{
											Computed: true,
										},
										"account_deletion": schema.StringAttribute{
											Computed: true,
										},
										"account_migration": schema.StringAttribute{
											Computed: true,
										},
										"account_mobility": schema.StringAttribute{
											Computed: true,
										},
										"enterprise_capability": schema.StringAttribute{
											Computed: true,
										},
										"member_management": schema.StringAttribute{
											Computed: true,
										},
										"sub_org_creation": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"tenant_type": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"parent": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[OrganizationsParentDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"profile": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[OrganizationsProfileDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"business_address": schema.StringAttribute{
									Computed: true,
								},
								"business_email": schema.StringAttribute{
									Computed: true,
								},
								"business_name": schema.StringAttribute{
									Computed: true,
								},
								"business_phone": schema.StringAttribute{
									Computed: true,
								},
								"external_metadata": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *OrganizationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
